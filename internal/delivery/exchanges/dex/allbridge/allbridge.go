package allbridge

import (
	"exchange-provider/internal/delivery/exchanges/dex/allbridge/calculate"
	"exchange-provider/internal/delivery/exchanges/dex/allbridge/contracts/erc20"
	"exchange-provider/internal/delivery/exchanges/dex/allbridge/networks/evm"
	"exchange-provider/internal/delivery/exchanges/dex/allbridge/types"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/logger"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type exchange struct {
	cfg *Config
	exs entity.ExchangeStore

	c     *cache
	pairs entity.PairsRepo
	repo  entity.OrderRepo
	ns    map[string]types.Network

	erc20 *abi.ABI

	tl     tokenList
	l      logger.Logger
	stopCh chan struct{}
}

func NewExchange(cfg *Config, exs entity.ExchangeStore, repo entity.OrderRepo,
	p entity.PairsRepo, l logger.Logger, fromDB bool) (entity.CrossDEX, error) {

	if err := cfg.validate(); err != nil {
		return nil, err
	}
	erc20, err := erc20.ContractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	ex := &exchange{
		cfg:    cfg,
		exs:    exs,
		pairs:  p,
		repo:   repo,
		ns:     make(map[string]types.Network),
		erc20:  erc20,
		l:      l,
		stopCh: make(chan struct{}),
	}

	calculate.Allbridge_Precision = 3

	tl, err := getTokenInfo(ex.cfg.Networks)
	if err != nil {
		return nil, err
	}
	ex.tl = tl

	for _, n := range cfg.Networks {
		switch n.Type {
		case types.EvmNetwork:
			net, err := evm.NewEvmNetwork(ex.NID(), n.Network, n.AllbridgeContract,
				tl[n.Network].BridgeAddress, n.MainContract, n.client, n.prvKey)
			if err != nil {
				return nil, err
			}
			ex.ns[n.Network] = net

		default:
			return nil, fmt.Errorf("type %s not supported", n.Type)
		}
	}

	c, err := newCache(ex, fromDB)
	if err != nil {
		return nil, err
	}
	ex.c = c

	ps, err := ex.createPairs(ex.cfg.ExchangeFee, ex.cfg.FeeRate, ex.tl, true)
	if err != nil {
		return nil, err
	}
	if len(ps) > 0 {
		err := ex.pairs.Add(ex, ps...)
		if err != nil {
			return nil, err
		}
	}
	go ex.run(ex.stopCh)
	return ex, nil
}

func (ex *exchange) network(t entity.TokenId) (types.Network, error) {
	n, ok := ex.ns[t.Standard+t.Network]
	if !ok {
		return nil, fmt.Errorf("network for %s not found", t.String())
	}
	return n, nil
}

func (ex *exchange) Id() uint                  { return ex.cfg.Id }
func (ex *exchange) Name() string              { return "allbridge" }
func (ex *exchange) NID() string               { return fmt.Sprintf("%s-%d", ex.Name(), ex.cfg.Id) }
func (ex *exchange) EnableDisable(enable bool) { ex.cfg.Enable = enable }
func (ex *exchange) IsEnable() bool            { return ex.cfg.Enable }
func (ex *exchange) Type() entity.ExType       { return entity.CrossDex }
func (ex *exchange) Configs() interface{}      { return ex.cfg }
func (ex *exchange) Remove()                   { close(ex.stopCh) }

func (ex *exchange) UpdatePairs() ([]string, error) {
	ps, err := ex.createPairs(ex.cfg.ExchangeFee, ex.cfg.FeeRate, ex.tl, false)
	if err != nil {
		return nil, err
	}
	pss := []string{}
	if len(ps) > 0 {
		err := ex.pairs.Add(ex, ps...)
		if err != nil {
			return nil, err
		}
		for _, p := range ps {
			pss = append(pss, p.String())
		}
	}
	return pss, nil
}
func (ex *exchange) run(stopCh <-chan struct{}) {
	agent := ex.agent("run")
	go ex.c.run(stopCh)
	t := time.NewTicker(time.Hour * 1)
	for {
		select {
		case <-t.C:
			ps, err := ex.createPairs(ex.cfg.ExchangeFee, ex.cfg.FeeRate, ex.tl, false)
			if err != nil {
				ex.l.Debug(agent, err.Error())
			}
			if len(ps) > 0 {
				err := ex.pairs.Add(ex, ps...)
				if err != nil {
					ex.l.Debug(agent, err.Error())
				}
			}
		case <-stopCh:
			return
		}
	}
}
