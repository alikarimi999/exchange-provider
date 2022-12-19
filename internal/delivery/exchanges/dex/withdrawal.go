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
	t, err := u.tokens.get(o.Withdrawal.TokenId)
	if err != nil {
		return "", err
	}

	value, err := numbers.FloatStringToBigInt(o.Withdrawal.Volume, t.Decimals)
	if err != nil {
		return "", err
	}

	pr := u.provider()
	contract, err := contracts.NewMain(t.Address, pr)
	if err != nil {
		return "", err
	}

	sender := common.HexToAddress(o.Deposit.Addr)
	reciever := common.HexToAddress(o.Withdrawal.Addr)

	// unwrap
	if t.IsNative() {

		if !o.Withdrawal.Unwrapped {

			unwrapAmount, err := numbers.FloatStringToBigInt(o.Withdrawal.Volume, t.Decimals)
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
				P:        pr,
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
				o.Withdrawal.Fee = utils.TxFee(tf.Tx.GasPrice(), tf.Receipt.GasUsed)
				o.Withdrawal.FeeCurrency = u.cfg.NativeToken + "-" + u.cfg.chainId
				u.l.Debug(agent, fmt.Sprintf("order: `%d`, unwrap-tx: `%s`, confirm: `%d/%d`",
					o.Id, tf.TxHash, tf.Confirmed, tf.Confirms))
			}
		}

		tx, err := transferNative(u.wallet, sender, reciever, int64(u.cfg.ChainId), value, pr)
		if err != nil {
			return "", err
		}
		o.Withdrawal.TxId = tx.Hash().String()
		u.l.Debug(agent, fmt.Sprintf("order: `%d`, tx: `%s`", o.Id, tx.Hash()))
		return tx.Hash().String(), nil
	}

	opts, err := u.wallet.NewKeyedTransactorWithChainID(sender, big.NewInt(0), int64(u.cfg.ChainId))
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

	o.Withdrawal.TxId = tx.Hash().String()
	u.l.Debug(agent, fmt.Sprintf("order: `%d`, tx: `%s`", o.Id, tx.Hash()))

	return tx.Hash().String(), nil
}
