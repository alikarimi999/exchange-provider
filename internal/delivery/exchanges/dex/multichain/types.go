package multichain

type token struct {
	Address  string
	Name     string
	Symbol   string
	Decimals int
}

type chainId string
type pairId string
type tokenId string
type data struct {
	*token
	DestChains map[chainId]map[pairId]chain `json:"destChains`
}
type tokens = map[tokenId]*data
type chains = map[chainId]*tokens
type chain struct {
	token
	AnyToken       *token `json:"anytoken"`
	FromAnyToken   *token `json:"fromanytoken"`
	Underlying     *token
	Router         string `json:"router"`
	RouterABI      string `json:"routerABI"`
	DepositAddress string `json:"DepositAddress"`
	IsApprove      bool   `json:"isApprove"`
}
