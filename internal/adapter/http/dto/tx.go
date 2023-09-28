package dto

import (
	"exchange-provider/internal/entity"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

type evmTX struct {
	Network     string            `json:"network"`
	IsApproveTx bool              `json:"isApproveTx"`
	From        string            `json:"from"`
	To          string            `json:"to"`
	Data        string            `json:"data"`
	Value       string            `json:"value"`
	Developer   *entity.Developer `json:"developer"`
}

func evmTx(tx *entity.EvmTx) *evmTX {
	return &evmTX{
		Network:     tx.Network,
		IsApproveTx: tx.IsApproveTx,
		From:        tx.From,
		To:          tx.To,
		Data:        hexutil.Encode(tx.TxData),
		Value:       hexutil.EncodeBig(tx.Value),
		Developer:   tx.Developer,
	}
}
