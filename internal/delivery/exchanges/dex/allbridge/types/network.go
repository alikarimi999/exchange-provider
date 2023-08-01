package types

import (
	"exchange-provider/internal/entity"
)

type NetworkType string

const (
	EvmNetwork  NetworkType = "EVM"
	TronNetwork NetworkType = "TRON"
	SolNetwork  NetworkType = "SOLANA"
)

type Network interface {
	Type() NetworkType
	NeedApproval(in *entity.Token, owner string, minAmount float64) (bool, error)
	ApproveTx(in *entity.Token, owner string, step int) (entity.Tx, error)
	BridgeData(input interface{}) ([]byte, error)
	DownloadLogs(fromBlock, toBlock uint64) ([]*TokensReceivedLog, uint64, error)
}
