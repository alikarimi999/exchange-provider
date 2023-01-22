package dto

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

type evmTX struct {
	Type      uint8    `json:"type"`
	From      string   `json:"from"`
	To        string   `json:"to"`
	Data      string   `json:"data"`
	Value     *big.Int `json:"value"`
	Gas       uint64   `json:"gas"`
	GasPrice  *big.Int `json:"gasPrice"`
	GasFeeCap *big.Int `json:"gasFeeCap"`
	GasTipCap *big.Int `json:"gasTipCap"`
}

func evmTx(tx *types.Transaction, from string) *evmTX {
	return &evmTX{
		Type:      tx.Type(),
		From:      from,
		To:        tx.To().Hex(),
		Data:      hexutil.Encode(tx.Data()),
		Value:     tx.Value(),
		Gas:       tx.Gas(),
		GasPrice:  tx.GasPrice(),
		GasFeeCap: tx.GasFeeCap(),
		GasTipCap: tx.GasTipCap(),
	}
}
