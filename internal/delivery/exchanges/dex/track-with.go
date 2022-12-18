package dex

import (
	"exchange-provider/internal/delivery/exchanges/dex/utils"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/utils/numbers"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func (u *dex) TrackWithdrawal(o *entity.Order, done chan<- struct{},
	proccessedCh <-chan bool) {

	agent := u.agent("TrackWithdrawal")

	w := o.Withdrawal
	t, err := u.tokens.get(w.TokenId)
	if err != nil {
		w.Status = entity.WithdrawalFailed
		w.FailedDesc = err.Error()
		done <- struct{}{}
		<-proccessedCh
		return
	}

	var r common.Address
	if t.IsNative() {
		r = common.HexToAddress(w.Addr)
	} else {
		r = t.Address
	}

	doneCh := make(chan struct{})
	tf := &utils.TtFeed{
		P:        u.provider(),
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
		fee, _ := numbers.FloatStringToBigInt(f, t.Decimals)

		unwrapFee := new(big.Int)
		var err error
		if w.Fee != "" {
			unwrapFee, err = numbers.FloatStringToBigInt(w.Fee, t.Decimals)
			if err != nil {
				unwrapFee = big.NewInt(0)
			}
		}
		w.Fee = numbers.BigIntToFloatString(new(big.Int).Add(fee, unwrapFee), t.Decimals)
		w.FeeCurrency = u.cfg.NativeToken
		w.Status = entity.WithdrawalSucceed
		u.l.Debug(agent, fmt.Sprintf("order: `%d`, tx: `%s`, confirm: `%d/%d`",
			w.OrderId, tf.TxHash, tf.Confirmed, tf.Confirms))

	default:
		w.Status = entity.WithdrawalFailed
		w.FailedDesc = tf.Faildesc
	}

	done <- struct{}{}
	<-proccessedCh

}
