package kucoin

import (
	"exchange-provider/internal/delivery/exchanges/kucoin/dto"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
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
	BC          *kuCoin // base coin
	QC          *kuCoin // quote coin
	feeCurrency string
}

func (p *pair) Id() string     { return p.BC.CoinId + p.BC.ChainId + p.QC.CoinId + p.QC.ChainId }
func (p *pair) Symbol() string { return fmt.Sprintf("%s/%s", p.BC.String(), p.QC.String()) }

func (p *pair) String() string {
	return fmt.Sprintf("%s/%s", p.BC.String(), p.QC.String())
}

func (p *pair) snapshot() *pair {
	return &pair{
		BC:          p.BC.snapshot(),
		QC:          p.QC.snapshot(),
		feeCurrency: p.feeCurrency,
	}
}

func fromDto(p *dto.Pair) *pair {
	return &pair{
		BC: &kuCoin{
			CoinId:              p.C1.CoinId,
			ChainId:             p.C1.ChainId,
			BlockTime:           p.C1.BlockTime,
			WithdrawalPrecision: p.C1.WithdrawalPrecision,
		},
		QC: &kuCoin{
			CoinId:              p.C2.CoinId,
			ChainId:             p.C2.ChainId,
			BlockTime:           p.C2.BlockTime,
			WithdrawalPrecision: p.C2.WithdrawalPrecision,
		},
	}
}

func (p *pair) toEntity() *entity.Pair {
	return &entity.Pair{
		C1:          p.BC.toEntityCoin(),
		C2:          p.QC.toEntityCoin(),
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
		sp.pairs[p.Id()] = p.snapshot()
	}
}

func (sp *exPairs) get(c1, c2 *entity.Coin) (*pair, error) {
	sp.mux.Lock()
	defer sp.mux.Unlock()

	if p, exist := sp.pairs[pId(c1, c2)]; exist {
		return p, nil
	} else if p, exist = sp.pairs[pId(c2, c1)]; exist {
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
		pairs = append(pairs, p.snapshot())
	}
	return pairs
}
func pId(bc, qc *entity.Coin) string {
	return bc.CoinId + bc.ChainId + qc.CoinId + qc.ChainId
}
