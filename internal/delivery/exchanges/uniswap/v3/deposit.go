package uniswapv3

import (
	"fmt"
	"exchange-provider/internal/entity"
)

func (u *UniSwapV3) TrackDeposit(d *entity.Deposit, done chan<- struct{},
	proccessed <-chan bool) {
	if d.ChainId != chainId {
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

	u.dt.push(&dtFeed{
		d:     d,
		token: &t,
		done:  done,
		pCh:   proccessed,
	})
}
