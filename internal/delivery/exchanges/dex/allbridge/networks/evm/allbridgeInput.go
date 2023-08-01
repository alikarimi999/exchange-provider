package evm

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

var (
	allBridgeArguments, _ = abi.NewType("tuple", "struct data", []abi.ArgumentMarshaling{
		{Name: "bridge", Type: "address"},
		{Name: "tokenAddress", Type: "bytes32"},
		{Name: "recipient", Type: "bytes32"},
		{Name: "destinationChainId", Type: "uint256"},
		{Name: "receiveTokenAddress", Type: "bytes32"},
		{Name: "nonce", Type: "uint256"},
		{Name: "messenger", Type: "uint8"},
		{Name: "feeTokenAmount", Type: "uint256"},
	})

	allBridgeArgs = abi.Arguments{
		{Type: allBridgeArguments},
	}
)

type allBridgeInput struct {
	Bridge              common.Address
	TokenAddress        [32]byte
	Recipient           [32]byte
	DestinationChainId  *big.Int
	ReceiveTokenAddress [32]byte
	Nonce               *big.Int
	Messenger           uint8
	FeeTokenAmount      *big.Int
}

func allBridgeInputBytes(in, out, recipient common.Address, nonce *big.Int,
	destChainId, messenger int64, bridge common.Address) []byte {
	input := allBridgeInput{
		Bridge:              bridge,
		TokenAddress:        convertAddress2Bytes32(in),
		Recipient:           convertAddress2Bytes32(recipient),
		DestinationChainId:  big.NewInt(destChainId),
		ReceiveTokenAddress: convertAddress2Bytes32(out),
		Nonce:               nonce,
		Messenger:           uint8(messenger),
		FeeTokenAmount:      common.Big0,
	}

	b, _ := allBridgeArgs.Pack(input)
	return b
}

func convertAddress2Bytes32(a common.Address) [32]byte {
	var fixedByte [32]byte
	copy(fixedByte[:], a.Hash().Bytes())
	return fixedByte
}
