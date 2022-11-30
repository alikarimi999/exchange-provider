package multichain

import (
	"exchange-provider/internal/delivery/exchanges/dex/utils"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/utils/numbers"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func (u *Multichain) TrackWithdrawal(o *entity.Order, done chan<- struct{},
	proccessedCh <-chan bool) {

	agent := "TrackWithdrawal"

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
		P:        u.cs[chainId(t.ChainId)].provider(),
		TxHash:   common.HexToHash(w.WId),
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
		if w.ExchangeFee != "" {
			unwrapFee, err = numbers.FloatStringToBigInt(w.ExchangeFee, t.Decimals)
			if err != nil {
				unwrapFee = big.NewInt(0)
			}
		}
		w.ExchangeFee = numbers.BigIntToFloatString(new(big.Int).Add(fee, unwrapFee), t.Decimals)
		// w.ExchangeFeeCurrency = u.cs[chainId(t.Chain)].nativeToken
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
