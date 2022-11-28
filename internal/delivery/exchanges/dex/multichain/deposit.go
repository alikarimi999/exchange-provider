package multichain

import (
	"exchange-provider/internal/entity"
)

func (m *Multichain) TrackDeposit(o *entity.Order, done chan<- struct{},
	proccessed <-chan bool) {

	in := c2T(o.Routes[0].In)
	out := c2T(o.Routes[0].Out)

	p, err := m.pairs.get(in, out)
	if err != nil {
		o.Deposit.Status = entity.DepositFailed
		o.Deposit.FailedDesc = err.Error()
		done <- struct{}{}
		<-proccessed
		return
	}
	var t *token
	if p.t1.Symbol == in.Symbol {
		t = p.t1
	} else {
		t = p.t2
	}

	m.trackDeposit(&dtFeed{
		d:    o.Deposit,
		t:    t,
		done: done,
		pCh:  proccessed,
	})
}
