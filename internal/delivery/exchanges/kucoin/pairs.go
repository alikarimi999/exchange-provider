package kucoin

import (
	"fmt"
	"exchange-provider/internal/delivery/exchanges/kucoin/dto"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"sync"
	"time"
)

const (
	pairDelimiter = "-"
)

type kuCoin struct {
	CoinId              string `json:"coin_id"`
	ChainId             string `json:"chain_id"`
	address             string
	tag                 string
	BlockTime           time.Duration `json:"block_time"`
	ConfirmBlocks       int64         `json:"confirm_blocks"`
	minOrderSize        string
	maxOrderSize        string
	minWithdrawalSize   string
	minWithdrawalFee    string
	WithdrawalPrecision int `json:"withdrawal_precision"`
	orderPrecision      int

	needChain bool
}

func (k *kuCoin) String() string {
	return fmt.Sprintf("%s-%s", k.CoinId, k.ChainId)
}

func coinFromEntity(c *entity.Coin) *kuCoin {
	return &kuCoin{
		CoinId:  c.CoinId,
		ChainId: c.ChainId,
	}
}

func (k *kuCoin) snapshot() *kuCoin {
	return &kuCoin{
		CoinId:              k.CoinId,
		ChainId:             k.ChainId,
		address:             k.address,
		tag:                 k.tag,
		BlockTime:           k.BlockTime,
		ConfirmBlocks:       k.ConfirmBlocks,
		minOrderSize:        k.maxOrderSize,
		maxOrderSize:        k.maxOrderSize,
		minWithdrawalSize:   k.minWithdrawalSize,
		minWithdrawalFee:    k.minWithdrawalFee,
		WithdrawalPrecision: k.WithdrawalPrecision,
		orderPrecision:      k.orderPrecision,
		needChain:           k.needChain,
	}
}

func (k *kuCoin) toEntityCoin() *entity.PairCoin {
	return &entity.PairCoin{

		Coin: &entity.Coin{
			CoinId:  k.CoinId,
			ChainId: k.ChainId,
		},
		Address:             k.address,
		Tag:                 k.tag,
		BlockTime:           k.BlockTime,
		MinOrderSize:        k.minOrderSize,
		MaxOrderSize:        k.maxOrderSize,
		MinWithdrawalSize:   k.minWithdrawalSize,
		WithdrawalMinFee:    k.minWithdrawalFee,
		WithdrawalPrecision: k.WithdrawalPrecision,
		OrderPrecision:      k.orderPrecision,
		SetChain:            k.needChain,
	}
}

type pair struct {
	Id          string  // base.id + base.Chain.Id + quote.id + quote.Chain.Id
	Symbol      string  // base.id + pairDelimiter + quote.id
	BC          *kuCoin // base coin
	QC          *kuCoin // quote coin
	feeCurrency string
}

func (p *pair) String() string {
	return fmt.Sprintf("%s/%s", p.BC.String(), p.QC.String())
}

func (p *pair) snapshot() *pair {
	return &pair{
		Id:          p.Id,
		Symbol:      p.Symbol,
		BC:          p.BC.snapshot(),
		QC:          p.QC.snapshot(),
		feeCurrency: p.feeCurrency,
	}
}

func fromDto(p *dto.Pair) *pair {
	return &pair{
		Id:     p.BC.CoinId + p.BC.ChainId + p.QC.CoinId + p.QC.ChainId,
		Symbol: p.BC.CoinId + pairDelimiter + p.QC.CoinId,
		BC: &kuCoin{
			CoinId:              p.BC.CoinId,
			ChainId:             p.BC.ChainId,
			BlockTime:           p.BC.BlockTime,
			WithdrawalPrecision: p.BC.WithdrawalPrecision,
		},
		QC: &kuCoin{
			CoinId:              p.QC.CoinId,
			ChainId:             p.QC.ChainId,
			BlockTime:           p.QC.BlockTime,
			WithdrawalPrecision: p.QC.WithdrawalPrecision,
		},
	}
}

func (p *pair) toEntity() *entity.Pair {
	return &entity.Pair{
		BC:          p.BC.toEntityCoin(),
		QC:          p.QC.toEntityCoin(),
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

func (sp *exPairs) add(pairs ...*pair) {
	sp.mux.Lock()
	defer sp.mux.Unlock()
	for _, p := range pairs {
		sp.pairs[p.Id] = p.snapshot()
	}
}

func (sp *exPairs) exists(bc, qc *kuCoin) bool {
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
		return p.snapshot(), nil
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
		pairs = append(pairs, p.snapshot())
	}
	return pairs
}
