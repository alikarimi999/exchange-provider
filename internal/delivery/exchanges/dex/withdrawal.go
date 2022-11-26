package dex

import (
	"exchange-provider/internal/delivery/exchanges/dex/uniswap/v3/contracts"
	"exchange-provider/internal/delivery/exchanges/dex/utils"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/utils/numbers"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func (u *dex) Withdrawal(o *entity.Order) (string, error) {
	agent := u.agent("Withdrawal")

	var err error
	t, err := u.tokens.get(o.Withdrawal.CoinId)
	if err != nil {
		return "", err
	}

	value, err := numbers.FloatStringToBigInt(o.Withdrawal.Total, t.Decimals)
	if err != nil {
		return "", err
	}

	contract, err := contracts.NewMain(t.Address, u.provider())
	if err != nil {
		return "", err
	}

	sender := common.HexToAddress(o.Deposit.Addr)
	reciever := common.HexToAddress(o.Withdrawal.Addr)

	// unwrap
	if t.IsNative() {

		if !o.Withdrawal.Unwrapped {

			unwrapAmount, err := numbers.FloatStringToBigInt(o.Withdrawal.Total, t.Decimals)
			if err != nil {
				return "", err
			}
			tx, err := u.unwrap(sender, t.Address, unwrapAmount)
			if err != nil {
				return "", err
			}
			u.l.Debug(agent, fmt.Sprintf("order: `%d`, unwrap-tx: `%s`", o.Id, tx.Hash()))
			o.MetaData["unwrap-txId"] = tx.Hash().String()
			done := make(chan struct{})
			tf := &utils.TtFeed{
				P:        u.provider(),
				TxHash:   tx.Hash(),
				Receiver: &t.Address,
				NeedTx:   true,

				DoneCh: done,
			}
			go u.tt.Track(tf)
			<-done

			switch tf.Status {
			case utils.TxFailed:
				return "", errors.Wrap(errors.NewMesssage(fmt.Sprintf("unwrap-tx `%s` failed (%s)", tx.Hash(), tf.Faildesc)))
			case utils.TxSuccess:
				o.Withdrawal.Unwrapped = true
				o.Withdrawal.ExchangeFee = utils.TxFee(tf.Tx.GasPrice(), tf.Receipt.GasUsed)
				o.Withdrawal.ExchangeFeeCurrency = u.cfg.NativeToken
				u.l.Debug(agent, fmt.Sprintf("order: `%d`, unwrap-tx: `%s`, confirm: `%d/%d`",
					o.Id, tf.TxHash, tf.Confirmed, tf.Confirms))
			}
		}

		tx, err := TransferNative(u.wallet, sender, reciever, int64(u.cfg.ChianId), value, u.provider())
		if err != nil {
			return "", err
		}
		o.Withdrawal.Executed = o.Withdrawal.Total
		o.Withdrawal.TxId = tx.Hash().String()
		u.l.Debug(agent, fmt.Sprintf("order: `%d`, tx: `%s`", o.Id, tx.Hash()))
		return tx.Hash().String(), nil
	}

	opts, err := u.wallet.NewKeyedTransactorWithChainID(sender, big.NewInt(0), int64(u.cfg.ChianId))
	if err != nil {
		return "", err
	}

	defer func() {
		if err != nil {
			u.wallet.ReleaseNonce(sender, opts.Nonce.Uint64())
		} else {
			u.wallet.BurnNonce(sender, opts.Nonce.Uint64())

		}
	}()

	tx, err := contract.Transfer(opts, reciever, value)
	if err != nil {
		return "", err
	}

	o.Withdrawal.Executed = o.Withdrawal.Total
	o.Withdrawal.TxId = tx.Hash().String()
	u.l.Debug(agent, fmt.Sprintf("order: `%d`, tx: `%s`", o.Id, tx.Hash()))

	return tx.Hash().String(), nil
}

func (u *dex) TrackWithdrawal(w *entity.Withdrawal, done chan<- struct{},
	proccessedCh <-chan bool) {

	agent := u.agent("TrackWithdrawal")

	t, err := u.tokens.get(w.CoinId)
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
		fee, _ := numbers.FloatStringToBigInt(f, ethDecimals)

		unwrapFee := new(big.Int)
		var err error
		if w.ExchangeFee != "" {
			unwrapFee, err = numbers.FloatStringToBigInt(w.ExchangeFee, ethDecimals)
			if err != nil {
				unwrapFee = big.NewInt(0)
			}
		}
		w.ExchangeFee = numbers.BigIntToFloatString(new(big.Int).Add(fee, unwrapFee), ethDecimals)
		w.ExchangeFeeCurrency = u.cfg.NativeToken
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
