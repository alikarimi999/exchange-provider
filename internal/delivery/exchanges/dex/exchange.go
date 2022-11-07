package dex

import (
	ts "exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/internal/entity"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

func (d *dex) Exchange(o *entity.UserOrder, size, funds string) (string, error) {
	agent := d.agent("Exchange")

	bt := o.BC
	qt := o.QC
	side := o.Side

	pair, err := d.pairs.get(bt.CoinId, qt.CoinId)
	if err != nil {
		return "", err
	}

	sAddr := common.HexToAddress(o.Deposit.Addr)

	var tIn ts.Token
	var tOut ts.Token
	var amount string
	if side == entity.SideBuy {
		tIn = pair.QT
		tOut = pair.BT
		amount = funds
	} else {
		tIn = pair.BT
		tOut = pair.QT
		amount = size
	}

	tx, nonce, err := d.Swap(o, tIn, tOut, amount, sAddr, sAddr)
	if err != nil {
		if nonce != nil {
			d.wallet.ReleaseNonce(sAddr, nonce.Uint64())
		}
		return "", err
	}
	d.wallet.BurnNonce(sAddr, nonce.Uint64())

	o.ExchangeOrder.Funds = funds
	o.ExchangeOrder.Size = size
	o.ExchangeOrder.Symbol = pair.String()

	d.l.Debug(agent, fmt.Sprintf("order: `%d`, tx: `%s`", o.Id, tx.Hash().String()))
	return tx.Hash().String(), nil

}

func (d *dex) TrackExchangeOrder(o *entity.UserOrder, done chan<- struct{}, proccessed <-chan bool) {
	agent := d.agent("TrackExchangeOrder")
	pair, err := d.pairs.get(o.BC.CoinId, o.QC.CoinId)
	if err != nil {
		return
	}

	doneCh := make(chan struct{})
	tf := &ttFeed{
		txHash:   common.HexToHash(o.ExchangeOrder.ExId),
		receiver: &d.cfg.Router,
		needTx:   true,
		doneCh:   doneCh,
	}

	go d.tt.track(tf)

	<-doneCh

	switch tf.status {
	case txSuccess:
		amont, fee, err := d.ParseSwapLogs(o, tf.tx, pair, tf.Receipt)
		if err != nil {
			o.ExchangeOrder.Status = entity.ExOrderFailed
			o.ExchangeOrder.FailedDesc = err.Error()
		}
		if o.Side == entity.SideBuy {
			o.ExchangeOrder.Size = amont
		} else {
			o.ExchangeOrder.Funds = amont
		}
		o.ExchangeOrder.Status = entity.ExOrderSucceed
		o.ExchangeOrder.Fee = fee
		o.ExchangeOrder.FeeCurrency = d.cfg.NativeToken

		d.l.Debug(agent, fmt.Sprintf("order: `%d`, tx: `%s`, confirm: `%d/%d`",
			o.Id, tf.txHash, tf.confirmed, tf.confirms))

	case txFailed:
		o.ExchangeOrder.Status = entity.ExOrderFailed
		o.ExchangeOrder.FailedDesc = tf.faildesc
	}

	done <- struct{}{}
	<-proccessed

}
