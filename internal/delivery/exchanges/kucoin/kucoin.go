package kucoin

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/logger"
	"fmt"
	"strings"
	"sync"
	"time"

	"exchange-provider/pkg/errors"

	"github.com/Kucoin/kucoin-go-sdk"
	"github.com/go-redis/redis/v9"
	"github.com/spf13/viper"
)

type Configs struct {
	ApiKey        string `json:"api_key"`
	ApiSecret     string `json:"api_secret"`
	ApiPassphrase string `json:"api_passphrase"`
	ApiVersion    string
	ApiUrl        string
	Chains        map[chainId]struct {
		standard
		time.Duration
	}
	// WsTopics      []string
}

// kucoinExchange is a concrete implementation of entity.Exchange interface.
type kucoinExchange struct {
	cfg *Configs
	mux *sync.Mutex

	api *kucoin.ApiService
	// ws   *webSocket

	wt  *withdrawalTracker
	da  *depositAggregator
	dt  *depositTracker
	wa  *withdrawalAggregator
	pls *pairList

	v *viper.Viper
	l logger.Logger

	exchangePairs  *exPairs
	supportedCoins *supportedCoins

	stopCh   chan struct{}
	stopedAt time.Time
}

func NewKucoinExchange(cfgi interface{}, rc *redis.Client, v *viper.Viper,
	l logger.Logger, readConfig bool) (entity.Exchange, error) {
	const op = errors.Op("Kucoin-Exchange.NewKucoinExchange")

	cfg, err := validateConfigs(cfgi)
	if err != nil {
		return nil, errors.Wrap(string(op), err)
	}

	k := &kucoinExchange{
		cfg: cfg,
		mux: &sync.Mutex{},

		api: kucoin.NewApiService(
			kucoin.ApiBaseURIOption(cfg.ApiUrl),
			kucoin.ApiKeyOption(cfg.ApiKey),
			kucoin.ApiSecretOption(cfg.ApiSecret),
			kucoin.ApiPassPhraseOption(cfg.ApiPassphrase),
			kucoin.ApiKeyVersionOption(cfg.ApiVersion),
		),
		exchangePairs:  newExPairs(),
		supportedCoins: newSupportedCoins(),
		v:              v,
		l:              l,

		stopCh: make(chan struct{}),
	}

	if err := k.ping(); err != nil {
		return nil, err
	}
	k.l.Debug(string(op), "ping was successful")
	c := newCache(k, rc, k.l)

	k.wt = newWithdrawalTracker(k, c)
	k.da = newDepositAggregator(k, c)
	k.dt = newDepositTracker(k, c)
	k.wa = newWithdrawalAggregator(k, c)
	k.pls = newPairList(k, k.api, l)

	if readConfig {
		k.l.Debug(string(op), fmt.Sprintf("retriving pairs from config file %s", k.v.ConfigFileUsed()))

		i := k.v.Get(fmt.Sprintf("%s.pairs", k.Id()))
		if i != nil {
			if err := k.pls.download(); err != nil {
				return nil, err
			}

			psi := i.(map[string]interface{})

			ps := make(map[string]*pair)
			for _, v := range psi {
				p := v.(map[string]interface{})
				pc := &pair{}
				if p["bc"] != nil && p["qc"] != nil {
					pc.BC = &kuToken{
						TokenId:  p["bc"].(map[string]interface{})["tokenid"].(string),
						ChainId:  chainId(p["bc"].(map[string]interface{})["chainid"].(string)),
						Standard: standard((p["bc"].(map[string]interface{})["standard"].(string))),

						BlockTime:           time.Duration(p["bc"].(map[string]interface{})["block_time"].(float64)),
						ConfirmBlocks:       int64(p["bc"].(map[string]interface{})["confirm_blocks"].(float64)),
						WithdrawalPrecision: int(p["bc"].(map[string]interface{})["withdrawal_precision"].(float64)),
						needChain:           true,
					}
					if _, ok := k.cfg.Chains[chainId(pc.BC.ChainId)]; !ok {
						k.cfg.Chains[chainId(pc.BC.ChainId)] = struct {
							standard
							time.Duration
						}{
							pc.BC.Standard,
							pc.BC.BlockTime,
						}
					}

					pc.QC = &kuToken{
						TokenId:  p["qc"].(map[string]interface{})["tokenid"].(string),
						ChainId:  chainId(p["qc"].(map[string]interface{})["chainid"].(string)),
						Standard: standard((p["qc"].(map[string]interface{})["standard"].(string))),

						BlockTime:           time.Duration(p["qc"].(map[string]interface{})["block_time"].(float64)),
						ConfirmBlocks:       int64(p["qc"].(map[string]interface{})["confirm_blocks"].(float64)),
						WithdrawalPrecision: int(p["qc"].(map[string]interface{})["withdrawal_precision"].(float64)),
						needChain:           true,
					}
					if _, ok := k.cfg.Chains[chainId(pc.QC.ChainId)]; !ok {
						k.cfg.Chains[chainId(pc.QC.ChainId)] = struct {
							standard
							time.Duration
						}{
							pc.QC.Standard,
							pc.QC.BlockTime,
						}
					}

					ps[pc.Id()] = pc
				}
			}

			newPs := []*pair{}
			newCs := map[string]*kuToken{}
			for _, p := range ps {
				ok, _ := k.pls.support(p)
				if !ok {
					k.l.Debug(string(op), fmt.Sprintf("pair %s is not supported by kucoin anymore", p.String()))
					delete(k.v.Get(fmt.Sprintf("%s.pairs", k.Id())).(map[string]interface{}), strings.ToLower(p.Id()))
					if err := k.v.WriteConfig(); err != nil {
						k.l.Error(string(op), err.Error())
					}
					continue
				}

				if err := k.setInfos(p); err != nil {
					k.l.Error(string(op), fmt.Sprintf("failed to set infos for pair %s du to error (%s)",
						p.String(), err.Error()))
					continue
				}

				newPs = append(newPs, p)
				newCs[p.BC.TokenId+string(p.BC.ChainId)] = p.BC.snapshot()
				newCs[p.QC.TokenId+string(p.QC.ChainId)] = p.QC.snapshot()

			}

			k.exchangePairs.add(newPs...)
			k.supportedCoins.add(newCs)

			k.l.Info(string(op), fmt.Sprintf("%d pairs loaded", len(newPs)))
			k.l.Info(string(op), fmt.Sprintf("%d pairs couldn't be loaded", len(ps)-len(newPs)))
			k.l.Info(string(op), fmt.Sprintf("exchange %s started again successfully", k.Id()))

		}
	}
	return k, nil
}

func (k *kucoinExchange) Run(wg *sync.WaitGroup) {
	defer wg.Done()
	w := &sync.WaitGroup{}

	w.Add(1)
	go k.wt.run(w, k.stopCh)
	w.Add(1)
	go k.da.run(w, k.stopCh)
	wg.Add(1)
	go k.dt.run(w, k.stopCh)
	w.Add(1)
	go k.wa.run(w, k.stopCh)

	k.l.Debug(fmt.Sprintf("%s.Run", k.Id()), "started")
	w.Wait()
}

func (k *kucoinExchange) Stop() {
	op := fmt.Sprintf("%s.Stop", k.Id())
	close(k.stopCh)
	k.stopedAt = time.Now()
	k.l.Debug(string(op), "stopped")
}

func (k *kucoinExchange) Type() entity.ExType {
	return entity.CEX
}
