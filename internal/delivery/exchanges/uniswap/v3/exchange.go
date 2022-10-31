package uniswapv3

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/utils/numbers"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func (u *dex) Exchange(o *entity.UserOrder, size, funds string) (string, error) {
	agent := u.agent("Exchange")

	bt := o.BC
	qt := o.QC
	side := o.Side

	pair, err := u.pairs.get(bt.CoinId, qt.CoinId)
	if err != nil {
		return "", err
	}

	sAddr := common.HexToAddress(o.Deposit.Addr)

	var tIn Token
	var tOut Token
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

	tx, pool, err := u.swap(tIn, tOut, amount, sAddr, sAddr)
	if err != nil {
		return "", err
	}

	o.MetaData["swap-pool"] = pool.address.String()

	o.ExchangeOrder.Funds = funds
	o.ExchangeOrder.Size = size
	o.ExchangeOrder.Symbol = pair.String()

	u.l.Debug(agent, fmt.Sprintf("order: `%d`, tx: `%s`", o.Id, tx.Hash().String()))
	return tx.Hash().String(), nil

}

func (u *dex) TrackExchangeOrder(o *entity.UserOrder, done chan<- struct{}, proccessed <-chan bool) {
	agent := u.agent("TrackExchangeOrder")
	pair, err := u.pairs.get(o.BC.CoinId, o.QC.CoinId)
	if err != nil {
		return
	}

	doneCh := make(chan struct{})
	tf := &ttFeed{
		txHash:   common.HexToHash(o.ExchangeOrder.ExId),
		receiver: &u.cfg.Router,
		needTx:   true,
		doneCh:   doneCh,
	}

	go u.tt.track(tf)

	<-doneCh

start:
	switch tf.status {
	case txSuccess:
		for _, log := range tf.Receipt.Logs {
			if len(log.Topics) == 3 && log.Topics[0] == erc20TransferSignature &&
				hashToAddress(log.Topics[2]) == common.HexToAddress(o.Deposit.Addr) {

				if o.Side == entity.SideBuy {
					d := pair.BT.Decimals
					o.ExchangeOrder.Size = numbers.BigIntToFloatString(new(big.Int).SetBytes(log.Data), d)
				} else {
					d := pair.QT.Decimals
					o.ExchangeOrder.Funds = numbers.BigIntToFloatString(new(big.Int).SetBytes(log.Data), d)
				}
				o.ExchangeOrder.Status = entity.ExOrderSucceed
				o.ExchangeOrder.Fee = computeTxFee(tf.tx.GasPrice(), tf.Receipt.GasUsed)
				o.ExchangeOrder.FeeCurrency = "ETH"
				u.l.Debug(agent, fmt.Sprintf("order: `%d`, tx: `%s`, confirm: `%d/%d`",
					o.Id, tf.txHash, tf.confirmed, tf.confirms))
				break start
			}

		}

		o.ExchangeOrder.Status = entity.ExOrderFailed
		o.ExchangeOrder.FailedDesc = "unable to parse tx logs"

	case txFailed:
		o.ExchangeOrder.Status = entity.ExOrderFailed
		o.ExchangeOrder.FailedDesc = tf.faildesc
	}

	done <- struct{}{}
	<-proccessed

}
