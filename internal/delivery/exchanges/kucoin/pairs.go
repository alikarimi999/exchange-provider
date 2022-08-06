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
	coinId              string
	chainId             string
	minOrderSize        string
	maxOrderSize        string
	minWithdrawalSize   string
	minWithdrawalFee    string
	withdrawalPrecision int
	orderPrecision      int

	needChain bool
}

type pair struct {
	id          string  // base.id + base.Chain.Id + quote.id + quote.Chain.Id
	symbol      string  // base.id + pairDelimiter + quote.id
	bc          *kuCoin // base coin
	qc          *kuCoin // quote coin
	feeCurrency string
}

func fromEntity(ep *entity.Pair) *pair {
	return &pair{
		id:     ep.BC.CoinId + ep.BC.ChainId + ep.QC.CoinId + ep.QC.ChainId,
		symbol: ep.BC.CoinId + pairDelimiter + ep.QC.CoinId,
		bc: &kuCoin{
			coinId:              ep.BC.CoinId,
			chainId:             ep.BC.ChainId,
			minOrderSize:        ep.BC.MinOrderSize,
			maxOrderSize:        ep.BC.MaxOrderSize,
			minWithdrawalSize:   ep.BC.MinWithdrawalSize,
			minWithdrawalFee:    ep.BC.WithdrawalMinFee,
			withdrawalPrecision: ep.BC.WithdrawalPrecision,
			orderPrecision:      ep.BC.OrderPrecision,
			needChain:           ep.BC.SetChain,
		},
		qc: &kuCoin{
			coinId:              ep.QC.CoinId,
			chainId:             ep.QC.ChainId,
			minOrderSize:        ep.QC.MinOrderSize,
			maxOrderSize:        ep.QC.MaxOrderSize,
			minWithdrawalSize:   ep.QC.MinWithdrawalSize,
			minWithdrawalFee:    ep.QC.WithdrawalMinFee,
			withdrawalPrecision: ep.QC.WithdrawalPrecision,
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
				CoinId:  p.bc.coinId,
				ChainId: p.bc.chainId,
			},
			MinOrderSize:        p.bc.minOrderSize,
			MaxOrderSize:        p.bc.maxOrderSize,
			MinWithdrawalSize:   p.bc.minWithdrawalSize,
			WithdrawalMinFee:    p.bc.minWithdrawalFee,
			WithdrawalPrecision: p.bc.withdrawalPrecision,
			OrderPrecision:      p.bc.orderPrecision,
			SetChain:            p.bc.needChain,
		},
		QC: &entity.PairCoin{
			Coin: &entity.Coin{
				CoinId:  p.qc.coinId,
				ChainId: p.qc.chainId,
			},
			MinOrderSize:        p.qc.minOrderSize,
			MaxOrderSize:        p.qc.maxOrderSize,
			MinWithdrawalSize:   p.qc.minWithdrawalSize,
			WithdrawalMinFee:    p.qc.minWithdrawalFee,
			WithdrawalPrecision: p.qc.withdrawalPrecision,
			OrderPrecision:      p.qc.orderPrecision,
			SetChain:            p.qc.needChain,
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
		sp.pairs[p.id] = p
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

func (sp *exPairs) remove(bc, qc *entity.Coin) {
	sp.mux.Lock()
	defer sp.mux.Unlock()

	delete(sp.pairs, bc.CoinId+bc.ChainId+qc.CoinId+qc.ChainId)
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
