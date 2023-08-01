package evm

import (
	"crypto/ecdsa"
	"exchange-provider/internal/delivery/exchanges/dex/evm/contracts"
	"exchange-provider/internal/delivery/exchanges/dex/evm/types"
	"exchange-provider/internal/entity"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
)

func (d *exchange) CreateSwapBytes(in, out entity.TokenId, tokenOwner, sender, receiver,
	mainContract common.Address, amount, feeAmount float64,
	prvKey *ecdsa.PrivateKey) ([]byte, error) {

	p, err := d.pairs.Get(d.Id(), in.String(), out.String())
	if err != nil {
		return nil, err
	}

	var In, Out *entity.Token
	if p.T1.String() == in.String() {
		In = p.T1
		Out = p.T2
	} else {
		In = p.T2
		Out = p.T1
	}

	var feeTier uint64
	inT := types.TokenFromEntity(In)
	outT := types.TokenFromEntity(Out)

	if d.cfg.Version == 3 {
		_, feeTier, err = d.dex.EstimateAmountOut(inT, outT, amount)
		if err != nil {
			return nil, err
		}
	}
	decF := big.NewFloat(0).SetInt(math.BigPow(10, int64(In.Decimals)))
	totalAmountF := big.NewFloat(0).Mul(big.NewFloat(amount), decF)
	totalAmountI, _ := totalAmountF.Int(nil)

	feeAmountF := big.NewFloat(0).Mul(big.NewFloat(feeAmount), decF)
	feeAmountI, _ := feeAmountF.Int(nil)
	swapAmountI := big.NewInt(0).Sub(totalAmountI, feeAmountI)

	input, err := d.dex.TxData(inT, outT, receiver, swapAmountI, int64(feeTier))
	if err != nil {
		return nil, err
	}

	data := contracts.IExchangeAggregatorswapInput{
		TokenIn:      common.HexToAddress(In.ContractAddress),
		TokenOut:     common.HexToAddress(Out.ContractAddress),
		TotalAmount:  totalAmountI,
		FeeAmount:    feeAmountI,
		AmountIn:     swapAmountI,
		FromContract: tokenOwner.Hash().Big().Cmp(mainContract.Hash().Big()) == 0,
		Swapper:      d.dex.Router(),
		SwapperData:  input,
		Sender:       sender,
		Receiver:     receiver,
		Native:       In.Native,
	}

	sig, err := d.sign(data, prvKey)
	if err != nil {
		return nil, err
	}

	b, err := d.abi.Pack("Swap", data, sig)
	if err != nil {
		return nil, err
	}

	return b, nil
}
