package types

type PoolInfo struct {
	AValue             string `json:"aValue"`
	DValue             string `json:"dValue"`
	TokenBalance       string `json:"tokenBalance"`
	VUsdBalance        string `json:"vUsdBalance"`
	TotalLpAmount      string `json:"totalLpAmount"`
	AccRewardPerShareP string `json:"accRewardPerShareP"`
	P                  int    `json:"p"`
}
