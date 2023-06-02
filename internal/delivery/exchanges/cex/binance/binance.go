package binance

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/logger"
	"fmt"
	"sync"
	"time"

	"github.com/adshao/go-binance/v2"
)

const max_conccurrent_jobs = 20

type exchange struct {
	cfg *Configs
	mux *sync.Mutex

	c     *binance.Client
	pairs entity.PairsRepo
	st    entity.SpreadTable
	repo  entity.OrderRepo

	wa *withdrawalAggregator
	cl *coinsList
	l  logger.Logger

	stopedAt time.Time
	stopCh   chan struct{}
}

func NewExchange(cfg *Configs, repo entity.OrderRepo, pairs entity.PairsRepo,
	st entity.SpreadTable, l logger.Logger, fromDB bool) (entity.Cex, error) {
	ex := &exchange{
		cfg:    cfg,
		mux:    &sync.Mutex{},
		st:     st,
		repo:   repo,
		pairs:  pairs,
		c:      binance.NewClient(cfg.Api.ApiKey, cfg.Api.ApiSecret),
		l:      l,
		stopCh: make(chan struct{}),
	}
	agent := ex.agent("NewExchange")

	ex.wa = newWithdrawalAggregator(ex)
	cl, err := newCoinsLIst(ex.c, l, ex.stopCh)
	if err != nil {
		return nil, err
	}
	ex.cl = cl

	if fromDB {
		ps := pairs.GetAll(ex.cfg.Id)
		wg := &sync.WaitGroup{}
		waitChan := make(chan struct{}, max_conccurrent_jobs)

		for _, p := range ps {
			waitChan <- struct{}{}
			wg.Add(1)
			go func(p *entity.Pair) {
				defer func() {
					<-waitChan
					wg.Done()
				}()
				bt := p.T1.ET.(*Token)
				bc, ok := ex.cl.getCoin(bt.Coin, bt.Network)
				if !ok || (!bc.DepositEnable || !bc.WithdrawEnable) {
					ex.l.Debug(agent, fmt.Sprintf("somthing is wrong about pair %s: coin %s-%s",
						p.String(), bt.Coin, bt.Network))
					pairs.Remove(ex.cfg.Id, p.T1.String(), p.T2.String(), false)
					return
				}

				qt := p.T2.ET.(*Token)
				qc, ok := ex.cl.getCoin(qt.Coin, qt.Network)
				if !ok || (!qc.DepositEnable || !qc.WithdrawEnable) {
					ex.l.Debug(agent, fmt.Sprintf("somthing is wrong about pair %s: coin %s-%s",
						p.String(), qt.Coin, qt.Network))
					pairs.Remove(ex.cfg.Id, p.T1.String(), p.T2.String(), false)
					return
				}
				p.T1.ET.(*Token).setInfos(bc)
				p.T2.ET.(*Token).setInfos(qc)

				if err := ex.infos(p); err != nil {
					ex.l.Debug(agent, fmt.Sprintf("%s: %s", p.String(), err.Error()))
					pairs.Remove(ex.cfg.Id, p.T1.String(), p.T2.String(), false)
					return
				}
				if err := ex.pairs.Update(ex.Id(), p); err != nil {
					ex.pairs.Remove(ex.cfg.Id, p.T1.String(), p.T2.String(), false)
					ex.l.Debug(agent, err.Error())
					return
				}

			}(p)
		}
		wg.Wait()
	}

	go ex.wa.run(ex.stopCh)
	return ex, nil
}

func (k *exchange) Id() uint {
	return k.cfg.Id
}

func (ex *exchange) Name() string {
	return "binance"
}

func (ex *exchange) EnableDisable(enable bool) {
	ex.mux.Lock()
	defer ex.mux.Unlock()
	ex.cfg.Enable = enable
}
func (ex *exchange) IsEnable() bool {
	ex.mux.Lock()
	defer ex.mux.Unlock()
	return ex.cfg.Enable
}

func (ex *exchange) NID() string {
	return fmt.Sprintf("%s-%d", ex.Name(), ex.Id())
}

func (ex *exchange) Type() entity.ExType { return entity.CEX }

func (ex *exchange) Configs() interface{} { return ex.cfg }
func (ex *exchange) Remove() {
	op := fmt.Sprintf("%s.Stop", ex.NID())
	close(ex.stopCh)
	ex.stopedAt = time.Now()
	ex.l.Debug(string(op), "stopped")
}
