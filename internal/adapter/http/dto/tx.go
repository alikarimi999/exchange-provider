package dto

import (
	"exchange-provider/internal/entity"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

type evmTX struct {
	Type        uint8  `json:"type"`
	IsApproveTx bool   `json:"isApproveTx"`
	From        string `json:"from"`
	To          string `json:"to"`
	Data        string `json:"data"`
	Value       string `json:"value"`
	Gas         string `json:"gas"`
	GasPrice    string `json:"gasPrice"`
	GasFeeCap   string `json:"gasFeeCap"`
	GasTipCap   string `json:"gasTipCap"`
}

func evmTx(et *entity.EvmTx) *evmTX {
	tx := et.Tx
	return &evmTX{
		Type:        tx.Type(),
		IsApproveTx: et.IsApproveTx,
		From:        et.From(),
		To:          tx.To().Hex(),
		Data:        hexutil.Encode(tx.Data()),
		Value:       hexutil.EncodeBig(tx.Value()),
		Gas:         hexutil.EncodeUint64(tx.Gas()),
		GasPrice:    hexutil.EncodeBig(tx.GasPrice()),
		GasFeeCap:   hexutil.EncodeBig(tx.GasFeeCap()),
		GasTipCap:   hexutil.EncodeBig(tx.GasTipCap()),
	}
}
