package multichain

import (
	"exchange-provider/internal/delivery/exchanges/dex/utils"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/utils/numbers"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var erc20TransferSignature = common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")

type dtFeed struct {
	d    *entity.Deposit
	t    *token
	done chan<- struct{}
	pCh  <-chan bool
}

func (m *Multichain) trackDeposit(f *dtFeed) {
	agent := "Multichain-trackDeposit"

	txHash := common.HexToHash(f.d.TxId)

	doneCh := make(chan struct{})

	destAddress := common.HexToAddress("0")
	if f.t.Native {
		destAddress = common.HexToAddress(f.d.Addr)
	} else {
		destAddress = common.HexToAddress(f.t.Address)
	}

	tf := &utils.TtFeed{
		P:          m.cs[chainId(f.t.Chain)].provider(),
		TxHash:     txHash,
		Receiver:   &destAddress,
		NeedTx:     f.t.Native,
		MaxRetries: 30,
		Confirms:   1,
		DoneCh:     doneCh,
	}

	go m.tt.Track(tf)

	<-doneCh
	switch tf.Status {
	case utils.TxFailed:
		f.d.Status = entity.DepositFailed
		f.d.FailedDesc = tf.Faildesc
		f.done <- struct{}{}
	case utils.TxSuccess:
		if !f.t.Native {
			if len(tf.Logs) == 0 {
				f.d.Status = entity.DepositFailed
				f.d.FailedDesc = "invalid transaction"
				f.done <- struct{}{}
				break
			}
			log := tf.Logs[0]
			topic := log.Topics[0]

			if log.Address.String() != f.t.Address {
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
			if dAddress != common.HexToAddress(f.d.Addr) {
				f.d.Status = entity.DepositFailed
				f.d.FailedDesc = fmt.Sprintf("invalid destination address `%s`", dAddress)
				f.done <- struct{}{}
				break
			}

			bn := new(big.Int).SetBytes(log.Data)
			f.d.Volume = numbers.BigIntToFloatString(bn, int(f.t.Decimals))
			f.d.Status = entity.DepositConfirmed
			m.l.Debug(agent, fmt.Sprintf("order: `%d`, tx: `%s`, confirm: `%d/%d`",
				f.d.OrderId, tf.TxHash, tf.Confirmed, tf.Confirms))
			f.done <- struct{}{}
			break
		}
		f.d.Volume = numbers.BigIntToFloatString(tf.Tx.Value(), f.t.Decimals)
		f.d.Status = entity.DepositConfirmed
		m.l.Debug(agent, fmt.Sprintf("order: `%d`, tx: `%s`, confirm: `%d/%d`",
			f.d.OrderId, tf.TxHash, tf.Confirmed, tf.Confirms))
		f.done <- struct{}{}

	}
	<-f.pCh

}
