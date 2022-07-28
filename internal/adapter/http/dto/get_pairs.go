package dto

import "order_service/internal/entity"

type GetPair struct {
	BC     string `json:"base_coin"`
	BChain string `json:"base_chain"`
	QC     string `json:"quote_coin"`
	QChain string `json:"quote_chain"`
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
		BC:     p.BC.CoinId,
		BChain: p.BC.ChainId,
		QC:     p.QC.CoinId,
		QChain: p.QC.ChainId,
	}
}
