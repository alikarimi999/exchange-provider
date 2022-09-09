package uniswapv3

import (
	"math/big"
	"order_service/internal/entity"
	"order_service/pkg/utils/numbers"

	"github.com/ethereum/go-ethereum/common"
)

func (u *UniSwapV3) Exchange(o *entity.UserOrder, size, funds string) (string, error) {

	bt := o.BC
	qt := o.QC
	side := o.Side

	pair, err := u.pairs.get(bt.CoinId, qt.CoinId)
	if err != nil {
		return "", err
	}

	sAddr := common.HexToAddress(o.Deposit.Addr)
	dAddr := common.HexToAddress(o.Deposit.Addr)

	var tIn *token
	var tOut *token
	var amount string
	if side == entity.SideBuy {
		tIn = pair.qt
		tOut = pair.bt
		amount = funds
	} else {
		tIn = pair.qt
		tOut = pair.bt
		amount = size
	}

	tx, err := u.swap(tIn, tOut, amount, sAddr, dAddr)
	if err != nil {
		return "", err
	}

	o.ExchangeOrder.Symbol = pair.String()
	return tx.Hash().String(), nil

}

func (u *UniSwapV3) TrackExchangeOrder(o *entity.UserOrder, done chan<- struct{}, proccessed <-chan bool) {

	pair, err := u.pairs.get(o.BC.CoinId, o.QC.CoinId)
	if err != nil {
		return
	}

	doneCh := make(chan struct{})
	tf := &ttFeed{
		txHash:   common.HexToHash(o.ExchangeOrder.ExId),
		receiver: &routerV2,
		needTx:   true,
		doneCh:   doneCh,
	}

	u.tt.push(tf)

	<-doneCh

	for _, log := range tf.Receipt.Logs {
		if len(log.Topics) == 3 && log.Topics[0] == erc20TransferSignature &&
			hashToAddress(log.Topics[2]) == common.HexToAddress(o.Deposit.Addr) {

			if o.Side == entity.SideBuy {
				d := pair.bt.decimals
				o.ExchangeOrder.Size = numbers.BigIntToFloatString(new(big.Int).SetBytes(log.Data), d)
			} else {
				d := pair.qt.decimals
				o.ExchangeOrder.Funds = numbers.BigIntToFloatString(new(big.Int).SetBytes(log.Data), d)
			}
			break
		}

	}

	o.ExchangeOrder.Fee = computeTxFee(tf.tx.GasPrice(), tf.Receipt.GasUsed)
	o.ExchangeOrder.FeeCurrency = "ETH"

	done <- struct{}{}
	<-proccessed

}
