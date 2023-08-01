package types

import "exchange-provider/internal/entity"

type TokenInfo struct {
	entity.TokenId
	Name         string   `json:"name"`
	PoolAddress  string   `json:"poolAddress"`
	TokenAddress string   `json:"tokenAddress"`
	Decimals     int      `json:"decimals"`
	PoolInfo     PoolInfo `json:"poolInfo"`
	TransferTime map[string]TransferTime
	FeeShare     string `json:"feeShare"`
	Apr          string `json:"apr"`
	LpRate       string `json:"lpRate"`
	Chain        string
	ChainId      int
}
type TransferTime struct {
	Allbridge int `json:"allbridge"`
	Wormhole  int `json:"wormhole"`
}

type TxCostAmount struct {
	Swap      int64  `json:"swap"`
	Transfer  int64  `json:"transfer"`
	MaxAmount string `json:"maxAmount"`
}
type Chain struct {
	Tokens        []*TokenInfo            `json:"tokens"`
	ChainID       int                     `json:"chainId"`
	BridgeAddress string                  `json:"bridgeAddress"`
	SwapAddress   string                  `json:"swapAddress"`
	TransferTime  map[string]TransferTime `json:"transferTime"`
	Confirmations int                     `json:"confirmations"`
	TxCostAmount  TxCostAmount            `json:"txCostAmount"`
}
