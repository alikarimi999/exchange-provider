package pairsRepo

import (
	"context"
	"exchange-provider/internal/delivery/exchanges/cex/kucoin"
	"exchange-provider/internal/delivery/exchanges/cex/swapspace"
	"exchange-provider/internal/delivery/exchanges/dex/evm"
	"exchange-provider/internal/entity"
	"fmt"
	"strings"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
)

type exchangePairs struct {
	NID   string `bson:"_id"`
	ExId  uint
	Pairs []pair
}

type pair struct {
	Id string `bson:"_id"`
	T1 *token `bson:"t1"`
	T2 *token

	FeeRate    float64
	SpreadRate float64
}

func pFromEntity(p *entity.Pair) pair {
	return pair{
		Id:         pairId(p.T1.String(), p.T2.String()),
		T1:         fromEntity(p.T1),
		T2:         fromEntity(p.T2),
		FeeRate:    p.FeeRate,
		SpreadRate: p.SpreadRate,
	}
}

func (p *pair) toEntity(exNID string, exId uint) *entity.Pair {
	ep := &entity.Pair{
		Exchange:   exNID,
		LP:         exId,
		FeeRate:    p.FeeRate,
		SpreadRate: p.SpreadRate,
	}
	var t entity.ExchangeToken

	fn := func(et bson.Raw) entity.ExchangeToken {
		bson.Unmarshal(et, t)
		return t.Snapshot()
	}

	switch strings.Split(exNID, "-")[0] {
	case "swapspace":
		t = &swapspace.Token{}
	case "uniswapv3", "uniswapv2", "panckakeswapv2":
		t = &evm.Token{}
	case "kucoin":
		t = &kucoin.Token{}
	}
	ep.T1 = p.T1.toEntity(fn)
	ep.T2 = p.T2.toEntity(fn)
	return ep
}

func (pr *pairsRepo) retrievePairs() error {
	agent := pr.agent("retrievePairs")
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
			pr.eps[ep.NID].pairs[p.Id] = p.toEntity(ep.NID, ep.ExId)
			pr.l.Debug(agent, fmt.Sprintf("pair '%s' added to exchange '%s'", p.Id, ep.NID))
		}
	}
	return nil
}
