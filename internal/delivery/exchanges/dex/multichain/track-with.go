package multichain

import (
	"exchange-provider/internal/delivery/exchanges/dex/utils"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/utils/numbers"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func (u *Multichain) TrackWithdrawal(o *entity.CexOrder, done chan<- struct{},
	proccessedCh <-chan bool) {

	w := o.Withdrawal
	in := c2T(o.Routes[len(o.Routes)-1].In)
	out := c2T(o.Routes[len(o.Routes)-1].Out)

	p, err := u.pairs.get(in, out)
	if err != nil {
		o.Withdrawal.Status = entity.WithdrawalFailed
		done <- struct{}{}
		<-proccessedCh
		return
	}

	var t *Token
	if p.T1.ChainId == out.ChainId {
		t = p.T1
	} else {
		t = p.T2
	}

	var r common.Address
	if t.Native {
		r = common.HexToAddress(w.Addr)
	} else {
		r = common.HexToAddress(t.Address)
	}

	doneCh := make(chan struct{})
	tf := &utils.TtFeed{
		P:        u.cs[ChainId(t.ChainId)].provider(),
		TxHash:   common.HexToHash(w.TxId),
		Receiver: &r,
		NeedTx:   true,
		DoneCh:   doneCh,
	}
	go u.tt.Track(tf)
	<-doneCh

	switch tf.Status {
	case utils.TxSuccess:
		f := utils.TxFee(tf.Tx.GasPrice(), tf.Receipt.GasUsed)
		fee, _ := numbers.StringToBigFloat(f)
		unwrapFee, _ := numbers.StringToBigFloat(w.Fee)

		w.Fee = new(big.Float).Add(fee, unwrapFee).Text('f', utils.EthDecimals)
		w.Status = entity.WithdrawalSucceed
	default:
		w.Status = entity.WithdrawalFailed
		w.FailedDesc = tf.Faildesc
	}

	done <- struct{}{}
	<-proccessedCh

}
