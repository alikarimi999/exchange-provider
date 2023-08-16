package pairsRepo

import (
	"context"
	"exchange-provider/internal/delivery/exchanges/cex/binance"
	"exchange-provider/internal/delivery/exchanges/cex/kucoin"
	at "exchange-provider/internal/delivery/exchanges/dex/allbridge/types"
	et "exchange-provider/internal/delivery/exchanges/dex/evm/types"

	"exchange-provider/internal/entity"
	"strings"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
)

type exchangePairs struct {
	NID    string `bson:"_id"`
	ExId   uint
	ExType entity.ExType
	Pairs  []pair
}

type pair struct {
	Id string
	T1 *token
	T2 *token

	FeeRate1    float64
	FeeRate2    float64
	ExchangeFee float64
	Spreads     map[uint]float64
	Enable      bool
	EP          bson.Raw
}

func pFromEntity(p *entity.Pair) pair {
	ep, _ := bson.Marshal(p.EP)
	return pair{
		Id:          pairId(p.T1.String(), p.T2.String()),
		T1:          fromEntity(p.T1),
		T2:          fromEntity(p.T2),
		FeeRate1:    p.FeeRate1,
		FeeRate2:    p.FeeRate2,
		ExchangeFee: p.ExchangeFee,
		Spreads:     p.Spreads,
		Enable:      p.Enable,
		EP:          ep,
	}
}

func (p *pair) toEntity(exType entity.ExType, exNID string, exId uint) *entity.Pair {
	pair := &entity.Pair{
		Exchange:    exNID,
		LP:          exId,
		FeeRate1:    p.FeeRate1,
		FeeRate2:    p.FeeRate2,
		Spreads:     p.Spreads,
		ExchangeFee: p.ExchangeFee,
		Enable:      p.Enable,
	}

	var t entity.ExchangeToken
	fn := func(et bson.Raw) entity.ExchangeToken {
		bson.Unmarshal(et, t)
		return t.Snapshot()
	}

	switch exType {
	case entity.CEX:
		switch strings.Split(exNID, "-")[0] {
		case "kucoin":
			ep := &kucoin.ExchangePair{}
			bson.Unmarshal(p.EP, ep)
			pair.EP = ep
			t = &kucoin.Token{}

		case "binance":
			ep := &binance.ExchangePair{}
			bson.Unmarshal(p.EP, ep)
			pair.EP = ep
			t = &binance.Token{}
		}
	case entity.EvmDEX:
		ep := &et.ExchangePair{}
		bson.Unmarshal(p.EP, ep)
		pair.EP = ep
		t = &et.EToken{}
	case entity.CrossDex:
		if strings.Split(exNID, "-")[0] == "allbridge" {
			ep := &at.ExchangePair{}
			bson.Unmarshal(p.EP, ep)
			pair.EP = ep
			t = &at.EToken{}
		}
	}
	pair.T1 = p.T1.toEntity(fn)
	pair.T2 = p.T2.toEntity(fn)
	return pair
}

func (pr *pairsRepo) retrievePairs() error {
	// agent := pr.agent("retrievePairs")
	cur, err := pr.c.Find(context.Background(), bson.D{{}})
	if err != nil {
		return err
	}
	eps := []*exchangePairs{}
	if err := cur.All(context.Background(), &eps); err != nil {
		return err
	}

	for _, ep := range eps {
		pr.eps[ep.NID] = &exPairs{
			mux:   &sync.RWMutex{},
			exId:  ep.ExId,
			exNID: ep.NID,
			pairs: make(map[string]*entity.Pair),
		}

		for _, p := range ep.Pairs {
			pr.eps[ep.NID].pairs[p.Id] = p.toEntity(ep.ExType, ep.NID, ep.ExId)
		}
	}
	return nil
}
