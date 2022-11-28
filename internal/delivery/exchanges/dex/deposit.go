package dex

import (
	"exchange-provider/internal/entity"
	"fmt"
)

func (d *dex) TrackDeposit(o *entity.Order, done chan<- struct{},
	proccessed <-chan bool) {

	de := o.Deposit
	if de.ChainId != d.cfg.chainId {
		de.Status = entity.DepositFailed
		de.FailedDesc = fmt.Sprintf("chain %s not supported", de.ChainId)
		return
	}

	t, err := d.tokens.get(de.CoinId)
	if err != nil {
		de.Status = entity.DepositFailed
		de.FailedDesc = err.Error()
		done <- struct{}{}
		<-proccessed
		return
	}

	d.trackDeposit(&dtFeed{
		d:     de,
		token: &t,
		done:  done,
		pCh:   proccessed,
	})
}
