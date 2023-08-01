package binance

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/logger"
	"fmt"
	"time"

	"github.com/adshao/go-binance/v2"
)

type exchange struct {
	cfg *Configs

	c     *binance.Client
	pairs entity.PairsRepo
	st    entity.SpreadTable
	repo  entity.OrderRepo

	si *serverInfos

	wa *withdrawalAggregator
	l  logger.Logger

	stopedAt time.Time
	stopCh   chan struct{}
}

func NewExchange(cfg *Configs, repo entity.OrderRepo, pairs entity.PairsRepo,
	st entity.SpreadTable, l logger.Logger, fromDB bool) (entity.Cex, error) {
	ex := &exchange{
		cfg:    cfg,
		st:     st,
		repo:   repo,
		pairs:  pairs,
		c:      binance.NewClient(cfg.Api.ApiKey, cfg.Api.ApiSecret),
		l:      l,
		stopCh: make(chan struct{}),
	}

	si, err := newServerInfos(ex)
	if err != nil {
		return nil, err
	}
	ex.si = si
	ex.wa = newWithdrawalAggregator(ex)

	if fromDB {
		ps := pairs.GetAll(ex.cfg.Id)
		for _, p := range ps {
			bt := p.T1.ET.(*Token)
			bc, err := si.getCoin(bt.Coin, bt.Network)
			if err != nil || (!bc.DepositEnable || !bc.WithdrawEnable) {
				pairs.Remove(ex.cfg.Id, p.T1.String(), p.T2.String(), false)
				continue
			}

			qt := p.T2.ET.(*Token)
			qc, err := si.getCoin(qt.Coin, qt.Network)
			if err != nil || (!qc.DepositEnable || !qc.WithdrawEnable) {
				pairs.Remove(ex.cfg.Id, p.T1.String(), p.T2.String(), false)
				continue
			}
			p.T1.ET.(*Token).setInfos(bc)
			p.T2.ET.(*Token).setInfos(qc)

			if err := ex.infos(p); err != nil {
				pairs.Remove(ex.cfg.Id, p.T1.String(), p.T2.String(), false)
				continue
			}
			if err := ex.pairs.Update(ex.Id(), p); err != nil {
				ex.pairs.Remove(ex.cfg.Id, p.T1.String(), p.T2.String(), false)
				continue
			}
		}
	}

	go ex.si.run(ex, ex.stopCh)
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
	ex.cfg.Enable = enable
}
func (ex *exchange) IsEnable() bool {
	return ex.cfg.Enable
}

func (ex *exchange) NID() string {
	return fmt.Sprintf("%s-%d", ex.Name(), ex.Id())
}

func (ex *exchange) UpdateStatus(eo entity.Order) error { return nil }

func (ex *exchange) Type() entity.ExType { return entity.CEX }

func (ex *exchange) Configs() interface{} { return ex.cfg }
func (ex *exchange) Remove() {
	op := fmt.Sprintf("%s.Stop", ex.NID())
	close(ex.stopCh)
	ex.stopedAt = time.Now()
	ex.l.Debug(string(op), "stopped")
}
