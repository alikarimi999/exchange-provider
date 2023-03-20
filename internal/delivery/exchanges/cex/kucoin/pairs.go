package kucoin

import (
	"exchange-provider/internal/delivery/exchanges/cex/kucoin/dto"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
	"sync"
	"time"
)

type kuToken struct {
	TokenId string  `json:"tokenId"`
	ChainId chainId `json:"chainId"`

	address string
	tag     string

	BlockTime     time.Duration `json:"block_time"`
	ConfirmBlocks int64         `json:"confirm_blocks"`

	minOrderSize string
	maxOrderSize string

	minWithdrawalSize string
	minWithdrawalFee  string

	WithdrawalPrecision int `json:"withdrawal_precision"`
	orderPrecision      int

	needChain bool
}

func (k *kuToken) String() string {
	return fmt.Sprintf("%s-%s", k.TokenId, k.ChainId)
}

func (k *kuToken) snapshot() *kuToken {
	return &kuToken{
		TokenId:             k.TokenId,
		ChainId:             k.ChainId,
		address:             k.address,
		tag:                 k.tag,
		BlockTime:           k.BlockTime,
		ConfirmBlocks:       k.ConfirmBlocks,
		minOrderSize:        k.minOrderSize,
		maxOrderSize:        k.maxOrderSize,
		minWithdrawalSize:   k.minWithdrawalSize,
		minWithdrawalFee:    k.minWithdrawalFee,
		WithdrawalPrecision: k.WithdrawalPrecision,
		orderPrecision:      k.orderPrecision,
		needChain:           k.needChain,
	}
}

func (k *kuToken) toEntityCoin() *entity.Token {
	return &entity.Token{
		Symbol:   k.TokenId,
		Standard: string(k.ChainId),
		Address:  k.address,

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
	BC          *kuToken // base coin
	QC          *kuToken // quote coin
	feeCurrency string
}

func (p *pair) Id() string {
	return p.BC.TokenId + string(p.BC.ChainId) +
		p.QC.TokenId + string(p.QC.ChainId)
}

func (p *pair) String() string {
	return fmt.Sprintf("%s/%s", p.BC.String(), p.QC.String())
}

func (p *pair) Symbol() string {
	return fmt.Sprintf("%s-%s", p.BC.TokenId, p.QC.TokenId)
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
		BC: &kuToken{
			TokenId:             p.T1.TokenId,
			ChainId:             chainId(p.T1.ChainId),
			BlockTime:           p.T1.BlockTime,
			WithdrawalPrecision: p.T1.WithdrawalPrecision,
		},
		QC: &kuToken{
			TokenId:             p.T2.TokenId,
			ChainId:             chainId(p.T2.ChainId),
			BlockTime:           p.T2.BlockTime,
			WithdrawalPrecision: p.T2.WithdrawalPrecision,
		},
	}
}

func (p *pair) toEntity() *entity.Pair {
	return &entity.Pair{
		T1:          p.BC.toEntityCoin(),
		T2:          p.QC.toEntityCoin(),
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

func (sp *exPairs) get(c1, c2 *entity.Token) (*pair, error) {
	sp.mux.Lock()
	defer sp.mux.Unlock()

	if p, exist := sp.pairs[pId(c1, c2)]; exist {
		return p, nil
	} else if p, exist = sp.pairs[pId(c2, c1)]; exist {
		return p, nil
	}

	return nil, errors.Wrap(errors.ErrNotFound, errors.NewMesssage("pair not found"))
}

func (sp *exPairs) remove(id string) {
	sp.mux.Lock()
	defer sp.mux.Unlock()
	delete(sp.pairs, id)
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
func pId(bc, qc *entity.Token) string {
	return bc.Symbol + bc.Standard + qc.Symbol + qc.Standard
}
