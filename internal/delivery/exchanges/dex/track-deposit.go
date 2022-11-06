package dex

import (
	"exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/utils/numbers"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type dtFeed struct {
	d     *entity.Deposit
	token *types.Token
	done  chan<- struct{}
	pCh   <-chan bool
}

func (d *dex) trackDeposit(f *dtFeed) {
	agent := d.agent("trackDeposit")
	txHash := common.HexToHash(f.d.TxId)

	doneCh := make(chan struct{})

	destAddress := common.HexToAddress("0")
	if f.token.IsNative() {
		destAddress = common.HexToAddress(f.d.Addr)
	} else {
		destAddress = f.token.Address
	}

	tf := &ttFeed{
		txHash:     txHash,
		receiver:   &destAddress,
		needTx:     f.token.IsNative(),
		maxRetries: 20,
		confirms:   3,
		doneCh:     doneCh,
	}

	go d.tt.track(tf)

	<-doneCh
	switch tf.status {
	case txFailed:
		f.d.Status = entity.DepositFailed
		f.d.FailedDesc = tf.faildesc
		f.done <- struct{}{}
	case txSuccess:
		if !f.token.IsNative() {
			if len(tf.Logs) == 0 {
				f.d.Status = entity.DepositFailed
				f.d.FailedDesc = "invalid transaction"
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
			d.l.Debug(agent, fmt.Sprintf("order: `%d`, tx: `%s`, confirm: `%d/%d`",
				f.d.OrderId, tf.txHash, tf.confirmed, tf.confirms))
			f.done <- struct{}{}
			break
		}
		f.d.Volume = numbers.BigIntToFloatString(tf.tx.Value(), f.token.Decimals)
		f.d.Status = entity.DepositConfirmed
		d.l.Debug(agent, fmt.Sprintf("order: `%d`, tx: `%s`, confirm: `%d/%d`",
			f.d.OrderId, tf.txHash, tf.confirmed, tf.confirms))
		f.done <- struct{}{}

	}
	<-f.pCh

}
