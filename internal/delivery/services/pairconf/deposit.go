package pairconf

import (
	"context"
	"exchange-provider/internal/entity"

	"go.mongodb.org/mongo-driver/bson"
)

type PairDepositLimit struct {
	Pair string

	Coin1 string
	MinC1 float64

	Coin2 string
	MinC2 float64
}

func (c *PairConfigs) PairMinDeposit(c1, c2 string) (float64, float64) {

	f, s := c.getCoins(c1, c2)
	if f != nil {
		return f.Min, s.Min
	}
	return 0, 0
}

func (c *PairConfigs) ChangeMinDeposit(ps ...*entity.PairMinDeposit) error {

	for _, p := range ps {
		c1, c2 := c.getCoins(p.C1.Coin, p.C2.Coin)
		if c1 != nil {
			if c1.Coin == p.C1.Coin {
				if p.C1.Min != 0 {
					c1.Min = p.C1.Min
				}
				if p.C2.Min != 0 {
					c2.Min = p.C2.Min
				}
			} else {
				if p.C1.Min != 0 {
					c2.Min = p.C1.Min
				}
				if p.C2.Min != 0 {
					c1.Min = p.C2.Min
				}
			}
		} else {
			c1 = p.C1
			c2 = p.C2
		}
		if err := c.add(c1, c2); err != nil {
			return err
		}
	}
	return nil
}

func (r *PairConfigs) AllMinDeposit() []*entity.PairMinDeposit {
	return r.allMinDeposit()
}

func (r *PairConfigs) allMinDeposit() []*entity.PairMinDeposit {
	var m []*entity.PairMinDeposit
	for _, v := range r.minDpositCache {
		m = append(m, &entity.PairMinDeposit{
			C1: &entity.CoinMinDeposit{Coin: v.Coin1, Min: v.MinC1},
			C2: &entity.CoinMinDeposit{Coin: v.Coin2, Min: v.MinC2},
		})
	}
	return m
}

func (r *PairConfigs) retriveMinDeposits() error {
	var pairs []*PairDepositLimit
	cur, err := r.db.Find(context.Background(), bson.D{})
	if err != nil {
		return err
	}

	if err := cur.All(context.Background(), &pairs); err != nil {
		return err
	}
	for _, v := range pairs {
		r.minDpositCache[v.Pair] = v
	}
	return nil
}

func (c *PairConfigs) getCoins(c1, c2 string) (first, second *entity.CoinMinDeposit) {
	if v, ok := c.minDpositCache[pairId(c1, c2)]; ok {
		return &entity.CoinMinDeposit{Coin: v.Coin1, Min: v.MinC1},
			&entity.CoinMinDeposit{Coin: v.Coin2, Min: v.MinC2}
	} else if v, ok := c.minDpositCache[pairId(c2, c1)]; ok {
		return &entity.CoinMinDeposit{Coin: v.Coin2, Min: v.MinC2},
			&entity.CoinMinDeposit{Coin: v.Coin1, Min: v.MinC1}
	} else {
		return nil, nil
	}
}

func (c *PairConfigs) add(c1, c2 *entity.CoinMinDeposit) error {
	pId := pairId(c1.Coin, c2.Coin)
	p := &PairDepositLimit{
		Pair: pId,

		Coin1: c1.Coin,
		MinC1: c1.Min,

		Coin2: c2.Coin,
		MinC2: c2.Min,
	}

	if _, err := c.db.InsertOne(context.Background(), p); err != nil {
		return err
	}
	c.minDpositCache[pId] = p
	return nil
}
