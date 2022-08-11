package kucoin

import (
	"order_service/internal/entity"
	"order_service/pkg/errors"
	"sync"
)

const (
	pairDelimiter = "-"
)

type kuCoin struct {
	CoinId              string `json:"coin_id"`
	ChainId             string `json:"chain_id"`
	minOrderSize        string
	maxOrderSize        string
	minWithdrawalSize   string
	minWithdrawalFee    string
	WithdrawalPrecision int `json:"withdrawal_precision"`
	orderPrecision      int

	needChain bool
}

type pair struct {
	Id          string  // base.id + base.Chain.Id + quote.id + quote.Chain.Id
	Symbol      string  // base.id + pairDelimiter + quote.id
	Bc          *kuCoin // base coin
	Qc          *kuCoin // quote coin
	feeCurrency string
}

func fromEntity(ep *entity.Pair) *pair {
	return &pair{
		Id:     ep.BC.CoinId + ep.BC.ChainId + ep.QC.CoinId + ep.QC.ChainId,
		Symbol: ep.BC.CoinId + pairDelimiter + ep.QC.CoinId,
		Bc: &kuCoin{
			CoinId:              ep.BC.CoinId,
			ChainId:             ep.BC.ChainId,
			minOrderSize:        ep.BC.MinOrderSize,
			maxOrderSize:        ep.BC.MaxOrderSize,
			minWithdrawalSize:   ep.BC.MinWithdrawalSize,
			minWithdrawalFee:    ep.BC.WithdrawalMinFee,
			WithdrawalPrecision: ep.BC.WithdrawalPrecision,
			orderPrecision:      ep.BC.OrderPrecision,
			needChain:           ep.BC.SetChain,
		},
		Qc: &kuCoin{
			CoinId:              ep.QC.CoinId,
			ChainId:             ep.QC.ChainId,
			minOrderSize:        ep.QC.MinOrderSize,
			maxOrderSize:        ep.QC.MaxOrderSize,
			minWithdrawalSize:   ep.QC.MinWithdrawalSize,
			minWithdrawalFee:    ep.QC.WithdrawalMinFee,
			WithdrawalPrecision: ep.QC.WithdrawalPrecision,
			orderPrecision:      ep.QC.OrderPrecision,
			needChain:           ep.QC.SetChain,
		},
		feeCurrency: ep.FeeCurrency,
	}
}

func (p *pair) toEntity() *entity.Pair {
	return &entity.Pair{
		BC: &entity.PairCoin{
			Coin: &entity.Coin{
				CoinId:  p.Bc.CoinId,
				ChainId: p.Bc.ChainId,
			},
			MinOrderSize:        p.Bc.minOrderSize,
			MaxOrderSize:        p.Bc.maxOrderSize,
			MinWithdrawalSize:   p.Bc.minWithdrawalSize,
			WithdrawalMinFee:    p.Bc.minWithdrawalFee,
			WithdrawalPrecision: p.Bc.WithdrawalPrecision,
			OrderPrecision:      p.Bc.orderPrecision,
			SetChain:            p.Bc.needChain,
		},
		QC: &entity.PairCoin{
			Coin: &entity.Coin{
				CoinId:  p.Qc.CoinId,
				ChainId: p.Qc.ChainId,
			},
			MinOrderSize:        p.Qc.minOrderSize,
			MaxOrderSize:        p.Qc.maxOrderSize,
			MinWithdrawalSize:   p.Qc.minWithdrawalSize,
			WithdrawalMinFee:    p.Qc.minWithdrawalFee,
			WithdrawalPrecision: p.Qc.WithdrawalPrecision,
			OrderPrecision:      p.Qc.orderPrecision,
			SetChain:            p.Qc.needChain,
		},
		FeeCurrency: p.feeCurrency,
	}
}

type exPairs struct {
	mux   *sync.Mutex
	pairs map[string]*pair // map[id]*pair
}

func newExPairs() *exPairs {
	return &exPairs{
		mux:   &sync.Mutex{},
		pairs: make(map[string]*pair),
	}
}

func (sp *exPairs) add(pairs []*pair) {
	sp.mux.Lock()
	defer sp.mux.Unlock()
	for _, p := range pairs {
		sp.pairs[p.Id] = p
	}
}

func (sp *exPairs) exists(bc, qc *entity.Coin) bool {
	sp.mux.Lock()
	defer sp.mux.Unlock()

	_, exist := sp.pairs[bc.CoinId+bc.ChainId+qc.CoinId+qc.ChainId]
	if exist {
		return true
	}

	return exist
}

func (sp *exPairs) get(bc, qc *entity.Coin) (*pair, error) {
	sp.mux.Lock()
	defer sp.mux.Unlock()

	p, exist := sp.pairs[bc.CoinId+bc.ChainId+qc.CoinId+qc.ChainId]
	if exist {
		return p, nil
	}

	return nil, errors.New("pair not found")
}

func (sp *exPairs) remove(id string) {
	sp.mux.Lock()
	defer sp.mux.Unlock()
	delete(sp.pairs, id)
}

func (sp *exPairs) purge() {
	sp.mux.Lock()
	defer sp.mux.Unlock()
	sp.pairs = make(map[string]*pair)
}

func (sp *exPairs) snapshot() []*pair {
	sp.mux.Lock()
	defer sp.mux.Unlock()

	pairs := make([]*pair, 0)
	for _, p := range sp.pairs {
		pairs = append(pairs, p)
	}
	return pairs
}
