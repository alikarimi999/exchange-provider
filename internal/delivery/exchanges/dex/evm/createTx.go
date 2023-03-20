package evm

import (
	"exchange-provider/internal/delivery/exchanges/dex/evm/contracts"
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

func errSTF() error {
	return errors.New("execution reverted: ExchangeAggregator::TransferHelper:safeTransferFrom")
}

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

	in, err := d.get(r.In.Symbol)
	if err != nil {
		return nil, err
	}

	out, err := d.get(r.Out.Symbol)
	if err != nil {
		return nil, err
	}

	_, fee, err := d.dex.EstimateAmountOut(in, out, amount)
	if err != nil {
		return nil, err
	}

	decF := big.NewFloat(0).SetInt(math.BigPow(10, int64(in.Decimals)))
	totalAmountF := big.NewFloat(0).Mul(big.NewFloat(amount), decF)
	totalAmountI, _ := totalAmountF.Int(nil)

	feeAmountF := big.NewFloat(0).Mul(big.NewFloat(amount*feeRate), decF)
	feeAmountI, _ := feeAmountF.Int(nil)

	swapAmountF := big.NewFloat(0).Sub(totalAmountF, feeAmountF)
	swapAmountI, _ := swapAmountF.Int(nil)

	input, err := d.dex.TxData(in, out, sender, receiver, swapAmountI, int64(fee))
	if err != nil {
		d.l.Debug(agent, err.Error())
		return nil, err
	}

	c, err := contracts.NewContracts(d.contractAddress, d.provider().Client)
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

	data := contracts.IExchangeAggregatorswapData{
		Input:       in.Address,
		TotalAmount: totalAmountI,
		FeeAmount:   feeAmountI,
		Swapper:     d.dex.Router(),
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

	if in.Native {
		opts.Value = totalAmountI
		tx, err = c.SwapNativeIn(opts, data, sig)
	} else {
		tx, err = c.Swap(opts, data, sig)
	}

	if err != nil {
		if err.Error() == errSTF().Error() {
			return nil, errors.Wrap(errors.ErrBadRequest,
				errors.NewMesssage("insufficient funds or the previous step was unsuccessful"))
		}
		if strings.Contains(err.Error(), "insufficient funds") {
			return nil, errors.Wrap(errors.ErrBadRequest,
				errors.NewMesssage(err.Error()))
		}
		d.l.Debug(agent, err.Error())
		return nil, err
	}

	// fmt.Println()
	// fmt.Println("input: ", data)
	// fmt.Println("data: 	", hexutil.Encode(data.Data))
	// fmt.Println("sig: 	", hexutil.Encode(sig))
	// fmt.Println()
	return tx, nil
}
