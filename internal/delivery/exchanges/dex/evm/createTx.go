package evm

import (
	"exchange-provider/internal/delivery/exchanges/dex/evm/contracts"
	"exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/internal/entity"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common/math"
	ts "github.com/ethereum/go-ethereum/core/types"

	"exchange-provider/pkg/bind"
	"exchange-provider/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

var (
	arguments, _ = abi.NewType("tuple", "struct data", []abi.ArgumentMarshaling{
		{Name: "input", Type: "address"},
		{Name: "totalAmount", Type: "uint256"},
		{Name: "feeAmount", Type: "uint256"},
		{Name: "swapper", Type: "address"},
		{Name: "data", Type: "bytes"},
		{Name: "sender", Type: "address"},
	})

	args = abi.Arguments{
		{Type: arguments},
	}
)

func (d *EvmDex) createTx(r *entity.Route, sender, receiver common.Address,
	amount, feeRate float64) (*ts.Transaction, error) {
	agent := d.agent("createTx")

	var (
		tx  *ts.Transaction
		err error
	)

	T1, ok := d.get(r.In.TokenId)
	if !ok {
		return nil, errors.Wrap(errors.ErrNotFound)
	}
	T2, ok := d.get(r.Out.TokenId)
	if !ok {
		return nil, errors.Wrap(errors.ErrNotFound)
	}

	var in, out types.Token
	if T1.Symbol == r.In.TokenId {
		in = T1
		out = T2
	} else {
		in = T2
		out = T1
	}

	decF := big.NewFloat(0).SetInt(math.BigPow(10, int64(in.Decimals)))
	totalAmountF := big.NewFloat(0).Mul(big.NewFloat(amount), decF)
	totalAmountI, _ := totalAmountF.Int(nil)

	feeAmountF := big.NewFloat(0).Mul(big.NewFloat(amount*feeRate), decF)
	feeAmountI, _ := feeAmountF.Int(nil)

	swapAmountF := big.NewFloat(0).Sub(totalAmountF, feeAmountF)
	swapAmountI, _ := swapAmountF.Int(nil)

	input, err := d.TxData(in, out, sender, receiver, swapAmountI)
	if err != nil {
		d.l.Debug(agent, err.Error())
		return nil, err
	}

	c, err := contracts.NewContracts(common.HexToAddress(d.Contract), d.provider().Client)
	if err != nil {
		d.l.Debug(agent, err.Error())
		return nil, err
	}

	opts, err := bind.NewKeyedTransactorWithChainID(d.privateKey, big.NewInt(d.ChainId))
	if err != nil {
		d.l.Debug(agent, err.Error())
		return nil, err
	}
	opts.NoSend = true

	data := contracts.ExchangeAggregatorswapData{
		Input:       in.Address,
		TotalAmount: totalAmountI,
		FeeAmount:   feeAmountI,
		Swapper:     d.Router(),
		Data:        input,
		Sender:      sender,
	}

	sig, err := d.sign(data)
	if err != nil {
		d.l.Debug(agent, err.Error())
		return nil, err
	}

	opts.From = sender
	opts.Sign = false

	if in.IsNative() {
		opts.Value = totalAmountI
		tx, err = c.SwapNativeIn(opts, data, sig)
	} else {
		tx, err = c.Swap(opts, data, sig)
	}

	if err != nil {
		if err.Error() == c.ErrSTF().Error() {
			return nil, errors.Wrap(errors.ErrBadRequest,
				errors.NewMesssage("you don't have enough token or previous step was not successful"))
		}
		if strings.Contains(err.Error(), "insufficient funds") {
			return nil, errors.Wrap(errors.ErrBadRequest,
				errors.NewMesssage(err.Error()))
		}
		d.l.Debug(agent, err.Error())
		return nil, err
	}

	return tx, nil
}
