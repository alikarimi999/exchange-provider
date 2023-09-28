package evm

import (
	"crypto/ecdsa"
	"exchange-provider/internal/delivery/exchanges/dex/evm/contracts"
	"exchange-provider/internal/delivery/exchanges/dex/evm/types"
	"exchange-provider/internal/entity"
	"math/big"

	"github.com/ethereum/go-ethereum/common/math"

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
		{Name: "receiver", Type: "address"},
		{Name: "native", Type: "bool"},
	})

	args = abi.Arguments{
		{Type: arguments},
	}
)

func (d *exchange) createTx(in, out *entity.Token, tokenOwner, sender, receiver,
	contractAddress common.Address, amount, feeAmount float64,
	prvKey *ecdsa.PrivateKey) ([]byte, [][]byte, *big.Int, error) {

	agent := d.agent("createTx")
	var (
		err error
	)

	var feeTier uint64
	inT := types.TokenFromEntity(in)
	outT := types.TokenFromEntity(out)

	if d.cfg.Version == 3 {
		_, feeTier, err = d.dex.EstimateAmountOut(inT, outT, amount)
		if err != nil {
			return nil, nil, nil, err
		}
	}
	decF := big.NewFloat(0).SetInt(math.BigPow(10, int64(in.Decimals)))
	totalAmountF := big.NewFloat(0).Mul(big.NewFloat(amount), decF)
	totalAmountI, _ := totalAmountF.Int(nil)

	feeAmountF := big.NewFloat(0).Mul(big.NewFloat(feeAmount), decF)
	feeAmountI, _ := feeAmountF.Int(nil)
	swapAmountI := big.NewInt(0).Sub(totalAmountI, feeAmountI)

	input, err := d.dex.TxData(inT, outT, receiver, swapAmountI, int64(feeTier))
	if err != nil {
		d.l.Debug(agent, err.Error())
		return nil, nil, nil, err
	}

	data := contracts.IExchangeAggregatorswapInput{
		TokenIn:      common.HexToAddress(in.ContractAddress),
		TokenOut:     common.HexToAddress(out.ContractAddress),
		TotalAmount:  totalAmountI,
		FeeAmount:    feeAmountI,
		AmountIn:     swapAmountI,
		FromContract: tokenOwner.Hash().Big().Cmp(contractAddress.Hash().Big()) == 0,
		Swapper:      d.dex.Router(),
		SwapperData:  input,
		Sender:       sender,
		Receiver:     receiver,
		Native:       in.Native,
	}

	sig, err := d.sign(data, d.cfg.prvKey)
	if err != nil {
		return nil, nil, nil, err
	}

	val := big.NewInt(0)

	if in.Native {
		val = totalAmountI
	}

	b, err := d.abi.Pack("Swap", data, sig)
	if err != nil {
		return nil, nil, nil, err
	}

	bs := [][]byte{b}
	b, err = d.abi.Pack("multicall", bs)
	return b, bs, val, err
}
