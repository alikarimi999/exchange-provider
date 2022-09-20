package uniswapv3

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/utils/numbers"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type dtFeed struct {
	d     *entity.Deposit
	token *token
	done  chan<- struct{}
	pCh   <-chan bool
}

func (u *UniSwapV3) trackDeposit(f *dtFeed) {
	agent := u.agent("trackDeposit")
	txHash := common.HexToHash(f.d.TxId)

	doneCh := make(chan struct{})

	destAddress := common.HexToAddress("0")
	if f.token.isNative() {
		destAddress = common.HexToAddress(f.d.Addr)
	} else {
		destAddress = f.token.Address
	}
	tf := &ttFeed{
		txHash:     txHash,
		receiver:   &destAddress,
		needTx:     f.token.isNative(),
		effortRate: 2,
		confirms:   3,
		doneCh:     doneCh,
	}

	go u.tt.track(tf)

	<-doneCh
	switch tf.status {
	case txNotFound:
		f.d.Status = entity.DepositFailed
		f.d.FailedDesc = "transaction not found"
		f.done <- struct{}{}
	case txFailed:
		f.d.Status = entity.DepositFailed
		f.d.FailedDesc = tf.faildesc
		f.done <- struct{}{}
	case txSuccess:
		if !f.token.isNative() {
			if len(tf.Logs) != 1 {
				f.d.Status = entity.DepositFailed
				f.d.FailedDesc = fmt.Sprintf("invalid transaction with `%d` logs", len(tf.Logs))
				f.done <- struct{}{}
				break
			}
			log := tf.Logs[0]
			topic := log.Topics[0]

			if log.Address != f.token.Address {
				f.d.Status = entity.DepositFailed
				f.d.FailedDesc = fmt.Sprintf("contract address of this transaction's log is incorrect `%s`", log.Address)
				f.done <- struct{}{}
				break
			}

			if topic != erc20TransferSignature {
				f.d.Status = entity.DepositFailed
				f.d.FailedDesc = fmt.Sprintf("invalid transaction log topic `%s`", topic.String())
				f.done <- struct{}{}
				break
			}
			dAddress := hashToAddress(log.Topics[2])
			if dAddress != common.HexToAddress(f.d.Addr) {
				f.d.Status = entity.DepositFailed
				f.d.FailedDesc = fmt.Sprintf("invalid destination address `%s`", dAddress)
				f.done <- struct{}{}
				break
			}

			bn := new(big.Int).SetBytes(log.Data)
			f.d.Volume = numbers.BigIntToFloatString(bn, int(f.token.Decimals))
			f.d.Status = entity.DepositConfirmed
			u.l.Debug(agent, fmt.Sprintf("order: `%d`, tx: `%s`, confirm: `%d/%d`",
				f.d.OrderId, tf.txHash, tf.confirmed, tf.confirms))
			f.done <- struct{}{}
			break
		}
		f.d.Volume = numbers.BigIntToFloatString(tf.tx.Value(), f.token.Decimals)
		f.d.Status = entity.DepositConfirmed
		u.l.Debug(agent, fmt.Sprintf("order: `%d`, tx: `%s`, confirm: `%d/%d`",
			f.d.OrderId, tf.txHash, tf.confirmed, tf.confirms))
		f.done <- struct{}{}

	}
	<-f.pCh

}
