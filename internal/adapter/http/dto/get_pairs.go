package dto

import "order_service/internal/entity"

type GetPair struct {
	BaseCoin   string `json:"base_coin"`
	BaseChain  string `json:"base_chain"`
	QuoteCoin  string `json:"quote_coin"`
	QuoteChain string `json:"quote_chain"`
}

type GetExchangesPairsRequest struct {
	Exchanges []string `json:"exchanges"`
}

type GetExchangesPairsResponse struct {
	Exchanges map[string][]*GetPair `json:"exchanges"`
	Messages  []string              `json:"messages"`
}

func ToDTO(p *entity.Pair) *GetPair {
	return &GetPair{
		BaseCoin:   p.BC.Id,
		BaseChain:  p.BC.Chain.Id,
		QuoteCoin:  p.QC.Id,
		QuoteChain: p.QC.Chain.Id,
	}
}
