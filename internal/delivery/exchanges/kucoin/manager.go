package kucoin

import (
	"order_service/internal/entity"
)

func (k *kucoinExchange) AddPairs(pairs []*entity.ExchangePair) {

	ps := []*pair{}
	for _, p := range pairs {
		ps = append(ps, newPair(p))
	}

	k.exchangePairs.add(ps)

	cs := map[string]*withdrawalCoin{}
	for _, p := range pairs {
		cs[p.BC.Id+p.BC.Chain.Id] = &withdrawalCoin{
			needChain: p.BC.SetChain,
		}

		cs[p.QC.Id+p.QC.Chain.Id] = &withdrawalCoin{
			needChain: p.QC.SetChain,
		}
	}

	k.withdrawalCoins.add(cs)

}

func (k *kucoinExchange) GetPairs() []*entity.Pair {
	pairs := []*entity.Pair{}
	ps := k.exchangePairs.snapshot()
	for _, p := range ps {
		pairs = append(pairs, &entity.Pair{
			BaseCoin:  p.b.Coin,
			QuoteCoin: p.q.Coin,
		})

	}
	return pairs
}

func (k *kucoinExchange) Support(c1, c2 *entity.Coin) bool {
	return k.exchangePairs.exists(c1, c2)
}
