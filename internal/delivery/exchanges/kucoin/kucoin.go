package kucoin

import (
	"fmt"
	"order_service/pkg/logger"
	"strings"
	"sync"
	"time"

	"order_service/pkg/errors"

	"github.com/Kucoin/kucoin-go-sdk"
	"github.com/go-redis/redis/v9"
	"github.com/spf13/viper"
)

type kucoinAdress struct {
	Address  string
	Chain    string
	Currency string
}

type Configs struct {
	ApiKey        string `json:"api_key"`
	ApiSecret     string `json:"api_secret"`
	ApiPassphrase string `json:"api_passphrase"`
	ApiVersion    string
	ApiUrl        string
	// WsTopics      []string
}

// kucoinExchange is a concrete implementation of entity.Exchange interface.
type kucoinExchange struct {
	cfg       *Configs
	mux       *sync.Mutex
	accountId string

	api *kucoin.ApiService
	// ws   *webSocket
	ot  *orderTracker
	wt  *withdrawalTracker
	wa  *withdrawalAggregator
	pls *pairList

	v *viper.Viper
	l logger.Logger

	exchangePairs   *exPairs
	withdrawalCoins *withdrawalCoins

	stopCh   chan struct{}
	stopedAt time.Time
}

func NewKucoinExchange(cfgi interface{}, rc *redis.Client, v *viper.Viper, l logger.Logger, readConfig bool) (*kucoinExchange, error) {
	const op = errors.Op("Kucoin-Exchange.NewKucoinExchange")

	cfg, err := validateConfigs(cfgi)
	if err != nil {
		return nil, errors.Wrap(string(op), err)
	}

	k := &kucoinExchange{
		cfg:       cfg,
		mux:       &sync.Mutex{},
		accountId: hash(cfg.ApiKey, cfg.ApiSecret, cfg.ApiPassphrase),

		api: kucoin.NewApiService(
			kucoin.ApiBaseURIOption(cfg.ApiUrl),
			kucoin.ApiKeyOption(cfg.ApiKey),
			kucoin.ApiSecretOption(cfg.ApiSecret),
			kucoin.ApiPassPhraseOption(cfg.ApiPassphrase),
			kucoin.ApiKeyVersionOption(cfg.ApiVersion),
		),
		exchangePairs:   newExPairs(),
		withdrawalCoins: newWithdrawalCoins(),
		v:               v,
		l:               l,

		stopCh: make(chan struct{}),
	}

	if err := k.ping(); err != nil {
		return nil, err
	}
	k.l.Debug(string(op), fmt.Sprintf("ping was successful"))

	k.ot = newOrderTracker(k, k.api, l)
	k.wt = newWithdrawalTracker(k, rc, l)
	k.wa = newWithdrawalAggregator(k, k.api, l, rc)
	k.pls = newPairList(k, k.api, l)

	if readConfig {
		k.l.Debug(string(op), fmt.Sprintf("retriving pairs from config file %s", k.v.ConfigFileUsed()))

		i := k.v.Get(fmt.Sprintf("%s.pairs", k.NID()))
		if i != nil {
			if err := k.pls.download(); err != nil {
				return nil, err
			}

			psi := i.(map[string]interface{})

			ps := make(map[string]*pair)
			for _, v := range psi {
				p := v.(map[string]interface{})
				pc := &pair{
					Id:     p["id"].(string),
					Symbol: p["symbol"].(string),
				}
				if p["bc"] != nil && p["qc"] != nil {
					pc.Bc = &kuCoin{
						CoinId:              p["bc"].(map[string]interface{})["coin_id"].(string),
						ChainId:             p["bc"].(map[string]interface{})["chain_id"].(string),
						WithdrawalPrecision: int(p["bc"].(map[string]interface{})["withdrawal_precision"].(float64)),
						needChain:           true,
					}
					pc.Qc = &kuCoin{
						CoinId:              p["qc"].(map[string]interface{})["coin_id"].(string),
						ChainId:             p["qc"].(map[string]interface{})["chain_id"].(string),
						WithdrawalPrecision: int(p["qc"].(map[string]interface{})["withdrawal_precision"].(float64)),
						needChain:           true,
					}
					ps[pc.Id] = pc
				}
			}

			newPs := []*pair{}
			newCs := map[string]*withdrawalCoin{}
			for _, p := range ps {
				pe := p.toEntity()
				ok, _ := k.pls.support(pe)
				if !ok {
					k.l.Debug(string(op), fmt.Sprintf("pair %s is not supported by kucoin anymore", pe.String()))
					delete(k.v.Get(fmt.Sprintf("%s.pairs", k.NID())).(map[string]interface{}), strings.ToLower(p.Id))
					if err := k.v.WriteConfig(); err != nil {
						k.l.Error(string(op), err.Error())
					}
					continue
				}

				if err := k.setInfos(pe); err != nil {
					k.l.Error(string(op), fmt.Sprintf("failed to set infos for pair %s du to error (%s)", pe.String(), err.Error()))
					continue
				}

				newPs = append(newPs, fromEntity(pe))
				newCs[pe.BC.CoinId+pe.BC.ChainId] = &withdrawalCoin{
					precision: p.Bc.WithdrawalPrecision,
					needChain: pe.BC.SetChain,
				}

				newCs[pe.QC.CoinId+pe.QC.ChainId] = &withdrawalCoin{
					precision: p.Qc.WithdrawalPrecision,
					needChain: pe.QC.SetChain,
				}

			}

			k.exchangePairs.add(newPs)
			k.withdrawalCoins.add(newCs)

			k.l.Info(string(op), fmt.Sprintf("%d pairs loaded", len(newPs)))
			k.l.Info(string(op), fmt.Sprintf("%d pairs couldn't be loaded", len(ps)-len(newPs)))
			k.l.Info(string(op), fmt.Sprintf("exchange %s started again successfully", k.NID()))

		}
	}
	return k, nil
}

func (k *kucoinExchange) Run(wg *sync.WaitGroup) {
	defer wg.Done()
	w := &sync.WaitGroup{}
	w.Add(1)
	go k.ot.run(w, k.stopCh)
	w.Add(1)
	go k.wt.run(w, k.stopCh)
	w.Add(1)
	go k.wa.run(w, k.stopCh)

	k.l.Debug("Kucoin-Exchange.Run", "started")
	w.Wait()
}

func (k *kucoinExchange) Stop() {
	close(k.stopCh)
	k.stopedAt = time.Now()
	k.l.Debug("Kucoin-Exchange", "stopped")
}
