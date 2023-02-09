package evm

import (
	"exchange-provider/internal/delivery/exchanges/dex/types"
)

func (d *EvmDex) findAllPairs() {
	agent := d.agent("findAllPairs")
	ps, err := d.retreivePairs()
	tps := []types.Pair{}
	if err != nil || len(ps) == 0 {
		if err != nil {
			d.l.Error(agent, err.Error())
		}
		for _, tA := range d.ts.Tokens {
			if tA.ChainId != int64(d.ChainId) {
				continue
			}

			for _, tB := range d.ts.Tokens {
				if tB.ChainId != int64(d.ChainId) || tB.Symbol == tA.Symbol {
					continue
				}

				t1 := types.Token{}
				t2 := types.Token{}
				if tA.Address.Hash().Big().Cmp(tB.Address.Hash().Big()) == -1 {
					t1 = tA
					t2 = tB
				} else {
					t2 = tA
					t1 = tB
				}

				tp := types.Pair{T1: t1, T2: t2}
				tps = append(tps, tp)
			}
		}

		d.SaveAvailablePairs(tps, d.PairsFile)
		ps, err = d.retreivePairs()
		if err != nil {
			return
		}
	}

	ps, err = d.price(ps...)
	if err != nil {
		return
	}
	d.pairsRepo.Add(d, ps...)
}
