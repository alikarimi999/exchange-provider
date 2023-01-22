package dex

import (
	ts "exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/internal/entity"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

func (d *dex) Swap(o *entity.CexOrder, index int) (string, error) {
	agent := d.agent("Swap")

	in := o.Routes[index].In
	out := o.Routes[index].Out

	pair, err := d.pairs.get(in.TokenId, out.TokenId)
	if err != nil {
		return "", err
	}

	sAddr := common.HexToAddress(o.Deposit.Addr)

	var tIn ts.Token
	var tOut ts.Token

	if in.TokenId == pair.T1.Symbol {
		tIn = pair.T1
		tOut = pair.T2
	} else {
		tOut = pair.T1
		tIn = pair.T2
	}

	tx, nonce, err := d.Dex.Swap(o, tIn, tOut, o.Swaps[index].InAmount, sAddr, sAddr)
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

func (d *dex) TrackSwap(o *entity.CexOrder, index int,
	done chan<- struct{}, proccessed <-chan bool) {

	pair, err := d.pairs.get(o.Routes[index].In.TokenId, o.Routes[index].Out.TokenId)
	if err != nil {
		o.Swaps[index].Status = entity.SwapFailed
		o.Swaps[index].FailedDesc = err.Error()
		done <- struct{}{}
		<-proccessed

		return
	}

	d.Dex.TrackSwap(o, pair, index)

	done <- struct{}{}
	<-proccessed

}
