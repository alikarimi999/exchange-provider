package evm

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	bridgeArguments, _ = abi.NewType("tuple", "struct data", []abi.ArgumentMarshaling{
		{Name: "bridge", Type: "address"},
		{Name: "tokenIn", Type: "address"},
		{Name: "sender", Type: "address"},
		{Name: "bridgeFee", Type: "uint256"},
		{Name: "afterSwap", Type: "bool"},
		{Name: "amountIn", Type: "uint256"},
		{Name: "feeAmount", Type: "uint256"},
		{Name: "bridgeData", Type: "bytes"},
	})

	bridgeArgs = abi.Arguments{
		{Type: bridgeArguments},
	}
)

type iBridgeAggregatorbridgeInput struct {
	Bridge     common.Address
	TokenIn    common.Address
	Sender     common.Address
	BridgeFee  *big.Int
	AfterSwap  bool
	AmountIn   *big.Int
	FeeAmount  *big.Int
	BridgeData []byte
}

func bridgeInputBytes(bridge, in, sender common.Address, bridgeFee,
	amountIn, feeAmount *big.Int, afterSwap bool, bridgeData []byte) (iBridgeAggregatorbridgeInput, error) {
	return iBridgeAggregatorbridgeInput{
		Bridge:     bridge,
		TokenIn:    in,
		Sender:     sender,
		BridgeFee:  bridgeFee,
		AfterSwap:  afterSwap,
		AmountIn:   amountIn,
		FeeAmount:  feeAmount,
		BridgeData: bridgeData,
	}, nil
}

func sign(b []byte, prvKey *ecdsa.PrivateKey) ([]byte, error) {
	return crypto.Sign(crypto.Keccak256Hash(b).Bytes(), prvKey)
}
