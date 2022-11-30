package multichain

import (
	"exchange-provider/internal/entity"
)

func (m *Multichain) GetAddress(c *entity.Coin) (*entity.Address, error) {
	a, err := m.cs[chainId(c.ChainId)].w.RandAddress()
	if err != nil {
		return nil, err
	}
	return &entity.Address{Addr: a.String()}, nil
}

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
	var t *Token
	if p.T1.ChainId == in.ChainId {
		t = p.T1
	} else {
		t = p.T2
	}

	m.trackDeposit(&dtFeed{
		d:    o.Deposit,
		t:    t,
		done: done,
		pCh:  proccessed,
	})
}
