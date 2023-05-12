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
	return errors.New("execution reverted: TransferHelper:safeTransferFrom")
}

var (
	arguments, _ = abi.NewType("tuple", "struct data", []abi.ArgumentMarshaling{
		{Name: "tokenIn", Type: "address"},
		{Name: "tokenOut", Type: "address"},
		{Name: "totalAmount", Type: "uint256"},
		{Name: "feeAmount", Type: "uint256"},
		{Name: "amountIn", Type: "uint256"},
		{Name: "fromContract", Type: "bool"},
		{Name: "swapper", Type: "address"},
		{Name: "swapperData", Type: "bytes"},
		{Name: "sender", Type: "address"},
		{Name: "native", Type: "bool"},
	})

	args = abi.Arguments{
		{Type: arguments},
	}
)

func (d *evmDex) createTx(in, out *entity.Token, tokenOwner, sender, receiver common.Address,
	amount, feeRate float64) (*ts.Transaction, error) {
	agent := d.agent("createTx")

	var (
		tx  *ts.Transaction
		err error
	)

	var feeTier uint64
	if d.Version == 3 {
		_, feeTier, err = d.dex.EstimateAmountOut(in, out, amount)
		if err != nil {
			return nil, err
		}
	}

	decF := big.NewFloat(0).SetInt(math.BigPow(10, int64(in.Decimals)))
	totalAmountF := big.NewFloat(0).Mul(big.NewFloat(amount), decF)
	totalAmountI, _ := totalAmountF.Int(nil)

	feeAmountF := big.NewFloat(0).Mul(totalAmountF, big.NewFloat(feeRate/100))
	feeAmountI, _ := feeAmountF.Int(nil)

	swapAmountI := big.NewInt(0).Sub(totalAmountI, feeAmountI)

	input, err := d.dex.TxData(in, out, receiver, swapAmountI, int64(feeTier))
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
	opts.From = sender
	opts.Sign = false

	data := contracts.IExchangeAggregatorswapInput{
		TokenIn:      common.HexToAddress(in.ContractAddress),
		TokenOut:     common.HexToAddress(out.ContractAddress),
		TotalAmount:  totalAmountI,
		FeeAmount:    feeAmountI,
		AmountIn:     swapAmountI,
		FromContract: tokenOwner.Hash().Big().Cmp(d.contractAddress.Hash().Big()) == 1,
		Swapper:      d.dex.Router(),
		SwapperData:  input,
		Sender:       sender,
		Native:       in.Native,
	}

	sig, err := d.sign(data)
	if err != nil {
		d.l.Debug(agent, err.Error())
		return nil, err
	}

	if in.Native {
		opts.Value = totalAmountI
	}
	tx, err = c.Swap(opts, data, sig)

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

	return tx, nil
}
