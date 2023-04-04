package dto

import (
	"exchange-provider/internal/entity"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

type evmTX struct {
	Type        uint8    `json:"type"`
	IsApproveTx bool     `json:"isApproveTx"`
	From        string   `json:"from"`
	To          string   `json:"to"`
	Data        string   `json:"data"`
	Value       *big.Int `json:"value"`
	Gas         uint64   `json:"gas"`
	GasPrice    *big.Int `json:"gasPrice"`
	GasFeeCap   *big.Int `json:"gasFeeCap"`
	GasTipCap   *big.Int `json:"gasTipCap"`
}

func evmTx(t entity.Tx, from string) *evmTX {
	et := t.(*entity.EvmTx)
	tx := et.Tx
	return &evmTX{
		Type:        tx.Type(),
		IsApproveTx: et.IsApproveTx,
		From:        from,
		To:          tx.To().Hex(),
		Data:        hexutil.Encode(tx.Data()),
		Value:       tx.Value(),
		Gas:         tx.Gas(),
		GasPrice:    tx.GasPrice(),
		GasFeeCap:   tx.GasFeeCap(),
		GasTipCap:   tx.GasTipCap(),
	}
}
