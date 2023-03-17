package dex

import (
	"exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/internal/delivery/exchanges/dex/utils"
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
	txHash := common.HexToHash(f.d.TxId)

	doneCh := make(chan struct{})

	destAddress := common.HexToAddress("0")
	if f.token.IsNative() {
		destAddress = common.HexToAddress(f.d.Address.Addr)
	} else {
		destAddress = f.token.Address
	}

	tf := &utils.TtFeed{
		P:          d.provider(),
		TxHash:     txHash,
		Receiver:   &destAddress,
		NeedTx:     f.token.IsNative(),
		MaxRetries: 30,
		Confirms:   1,
		DoneCh:     doneCh,
	}

	go d.tt.Track(tf)

	<-doneCh
	switch tf.Status {
	case utils.TxFailed:
		f.d.Status = entity.DepositFailed
		f.d.FailedDesc = tf.Faildesc
		f.done <- struct{}{}
	case utils.TxSuccess:
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
			dAddress := utils.HashToAddress(log.Topics[2])
			if dAddress != common.HexToAddress(f.d.Address.Addr) {
				f.d.Status = entity.DepositFailed
				f.d.FailedDesc = fmt.Sprintf("invalid destination address `%s`", dAddress)
				f.done <- struct{}{}
				break
			}

			bn := new(big.Int).SetBytes(log.Data)
			bf := numbers.BigIntToFloat(bn, int(f.token.Decimals))
			f.d.Volume, _ = bf.Float64()
			f.d.Status = entity.DepositConfirmed
			f.done <- struct{}{}
			break
		}

		bf := numbers.BigIntToFloat(tf.Tx.Value(), f.token.Decimals)
		f.d.Volume, _ = bf.Float64()
		f.d.Status = entity.DepositConfirmed
		f.done <- struct{}{}

	}
	<-f.pCh

}
