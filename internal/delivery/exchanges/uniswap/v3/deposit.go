package uniswapv3

import (
	"exchange-provider/internal/entity"
	"fmt"
)

func (u *dex) TrackDeposit(d *entity.Deposit, done chan<- struct{},
	proccessed <-chan bool) {
	if d.ChainId != u.cfg.TokenStandard {
		d.Status = entity.DepositFailed
		d.FailedDesc = fmt.Sprintf("chain %s not supported", d.ChainId)
		return
	}

	t, err := u.tokens.get(d.CoinId)
	if err != nil {
		d.Status = entity.DepositFailed
		d.FailedDesc = err.Error()
		done <- struct{}{}
		<-proccessed
		return
	}

	u.trackDeposit(&dtFeed{
		d:     d,
		token: &t,
		done:  done,
		pCh:   proccessed,
	})
}
