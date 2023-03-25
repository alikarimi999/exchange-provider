package dto

type EToken struct {
	Symbol   string `json:"symbol"`
	Standard string `json:"standard"`
	Network  string `json:"network"`

	Address  string `json:"address"`
	Decimals uint64 `json:"decimals"`
	Native   bool   `json:"native"`
	ET       Token  `json:"exchangeToken"`
}

type Token struct{}
