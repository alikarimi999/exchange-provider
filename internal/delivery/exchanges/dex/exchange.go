package dex

import (
	ts "exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/internal/entity"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

func (d *dex) Exchange(o *entity.Order, index int) (string, error) {
	agent := d.agent("Exchange")

	in := o.Routes[index].Input
	out := o.Routes[index].Output

	pair, err := d.pairs.get(in.CoinId, out.CoinId)
	if err != nil {
		return "", err
	}

	sAddr := common.HexToAddress(o.Deposit.Addr)

	var tIn ts.Token
	var tOut ts.Token

	if in.CoinId == pair.T1.Symbol {
		tIn = pair.T1
		tOut = pair.T2
	} else {
		tOut = pair.T2
		tIn = pair.T1
	}

	tx, nonce, err := d.Swap(o, tIn, tOut, o.Swaps[index].InAmount, sAddr, sAddr)
	if err != nil {
		if nonce != nil {
			d.wallet.ReleaseNonce(sAddr, nonce.Uint64())
		}
		return "", err
	}
	if nonce != nil {
		d.wallet.BurnNonce(sAddr, nonce.Uint64())
	}

	d.l.Debug(agent, fmt.Sprintf("order: `%d`, tx: `%s`", o.Id, tx.Hash().String()))
	return tx.Hash().String(), nil

}

func (d *dex) TrackExchangeOrder(o *entity.Order, index int,
	done chan<- struct{}, proccessed <-chan bool) {

	pair, err := d.pairs.get(o.Routes[index].Input.CoinId, o.Routes[index].Output.CoinId)
	if err != nil {
		o.Swaps[index].Status = entity.ExOrderFailed
		o.Swaps[index].FailedDesc = err.Error()
		done <- struct{}{}
		<-proccessed

		return
	}

	d.TrackSwap(o, pair, index)

	done <- struct{}{}
	<-proccessed

}
