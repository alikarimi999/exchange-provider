package multichain

import (
	"exchange-provider/internal/delivery/exchanges/dex/uniswap/v3/contracts"
	"exchange-provider/internal/delivery/exchanges/dex/utils"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/utils/numbers"
	"fmt"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
)

func (m *Multichain) Withdrawal(o *entity.CexOrder) (string, error) {
	agent := "Withdrawal"

	in := c2T(o.Routes[len(o.Routes)-1].In)
	out := c2T(o.Routes[len(o.Routes)-1].Out)

	var err error
	p, err := m.pairs.get(in, out)
	if err != nil {
		return "", err
	}

	var t *Token
	if p.T1.ChainId == out.ChainId {
		t = p.T1
	} else {
		t = p.T2
	}
	cid, _ := strconv.Atoi(t.ChainId)

	value, err := numbers.FloatStringToBigInt(o.Withdrawal.Volume, t.Decimals)
	if err != nil {
		return "", err
	}

	pr := m.cs[ChainId(out.ChainId)].provider()
	contract, err := contracts.NewMain(common.HexToAddress(t.Address), pr)
	if err != nil {
		return "", err
	}

	sender := common.HexToAddress(o.Deposit.Address.Addr)
	reciever := common.HexToAddress(o.Withdrawal.Addr)

	// unwrap
	if t.Native {

		if !o.Withdrawal.Unwrapped {

			unwrapAmount, err := numbers.FloatStringToBigInt(o.Withdrawal.Volume, t.Decimals)
			if err != nil {
				return "", err
			}
			tx, err := m.cs[ChainId(t.ChainId)].unwrap(sender, common.HexToAddress(t.Address), unwrapAmount)
			if err != nil {
				return "", err
			}
			m.l.Debug(agent, fmt.Sprintf("order: `%d`, unwrap-tx: `%s`", o.Id, tx.Hash()))
			o.MetaData["unwrap-txId"] = tx.Hash().String()
			done := make(chan struct{})
			r := common.HexToAddress(t.Address)
			tf := &utils.TtFeed{
				P:        pr,
				TxHash:   tx.Hash(),
				Receiver: &r,
				NeedTx:   true,

				DoneCh: done,
			}
			go m.tt.Track(tf)
			<-done

			switch tf.Status {
			case utils.TxFailed:
				return "", errors.Wrap(errors.NewMesssage(fmt.Sprintf("unwrap-tx `%s` failed (%s)", tx.Hash(), tf.Faildesc)))
			case utils.TxSuccess:
				o.Withdrawal.Unwrapped = true
				o.Withdrawal.Fee = utils.TxFee(tf.Tx.GasPrice(), tf.Receipt.GasUsed)
				// o.Withdrawal.ExchangeFeeCurrency = m.cs[chainId(t.Chain)].nativeToken
				m.l.Debug(agent, fmt.Sprintf("order: `%d`, unwrap-tx: `%s`, confirm: `%d/%d`",
					o.Id, tf.TxHash, tf.Confirmed, tf.Confirms))
			}
		}

		tx, err := transferNative(m.cs[ChainId(t.ChainId)].w, sender,
			reciever, int64(cid), value, pr)
		if err != nil {
			return "", err
		}
		o.Withdrawal.TxId = tx.Hash().String()
		m.l.Debug(agent, fmt.Sprintf("order: `%d`, tx: `%s`", o.Id, tx.Hash()))
		return tx.Hash().String(), nil
	}

	opts, err := m.cs[ChainId(t.ChainId)].w.NewKeyedTransactorWithChainID(sender,
		big.NewInt(0), int64(cid))
	if err != nil {
		return "", err
	}

	defer func() {
		if err != nil {
			m.cs[ChainId(t.ChainId)].w.ReleaseNonce(sender, opts.Nonce.Uint64())
		} else {
			m.cs[ChainId(t.ChainId)].w.BurnNonce(sender, opts.Nonce.Uint64())

		}
	}()

	tx, err := contract.Transfer(opts, reciever, value)
	if err != nil {
		return "", err
	}

	o.Withdrawal.TxId = tx.Hash().String()
	m.l.Debug(agent, fmt.Sprintf("order: `%d`, tx: `%s`", o.Id, tx.Hash()))

	return tx.Hash().String(), nil
}
