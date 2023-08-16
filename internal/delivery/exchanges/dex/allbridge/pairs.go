package allbridge

import (
	"exchange-provider/internal/delivery/exchanges/dex/allbridge/types"
	"exchange-provider/internal/entity"
	"fmt"
)

type pairs struct {
	ps []*entity.Pair
	ex entity.EVMDex
}

type pair struct {
	p   *entity.Pair
	p1  *entity.Pair
	p2  *entity.Pair
	ex1 entity.EVMDex
	ex2 entity.EVMDex
}

func (ex *exchange) createPairs(exchangeFee, feeRate float64,
	tl tokenList, updateMin bool) ([]*entity.Pair, error) {
	agent := ex.agent("createPairs")
	exs := ex.exs.GetAll()
	mps0 := make(map[string]*pairs)
	for _, e := range exs {
		if e.Id() != ex.Id() {
			if e.IsEnable() && e.Type() == entity.EvmDEX {
				ps := ex.pairs.GetAll(e.Id())
				mps0[e.NID()] = &pairs{ex: e.(entity.EVMDex)}
				for _, p := range ps {
					if p.Enable && (tl.isTokenExists(p.T1.Id) || tl.isTokenExists(p.T2.Id)) {
						mps0[e.NID()].ps = append(mps0[e.NID()].ps, p)
					}
				}
			}
		}
	}
	tokenEfa := make(map[string]float64)
	eps := ex.pairs.GetAll(ex.Id())
	epsStr := []string{}
	for _, ep := range eps {
		epsStr = append(epsStr, pairId(ep.T1.String(), ep.T2.String()))
	}

	if updateMin {
		for _, p := range eps {
			updated := ex.updateMinAndMax(p, tl, exchangeFee, feeRate, tokenEfa)
			if updated {
				if err := ex.pairs.Update(ex.Id(), p, false); err != nil {
					ex.l.Debug(agent, err.Error())
				}
			}
		}
	}

	mps1 := make(map[string]*pair)
	for ex1, ps1 := range mps0 {
		for ex2, ps2 := range mps0 {
			if ex1 == ex2 {
				continue
			}

			for _, p1 := range ps1.ps {
				for _, p2 := range ps2.ps {
					var (
						t1, t2                   *entity.Token
						otherToken1, otherToken2 string
					)

					if !tl.isTokenExists(p1.T1.Id) {
						t1 = p1.T1.Snapshot()
						otherToken1 = p1.T2.Id.String()
					} else {
						t1 = p1.T2.Snapshot()
						otherToken1 = p1.T1.Id.String()
					}

					if !tl.isTokenExists(p2.T1.Id) {
						t2 = p2.T1.Snapshot()
						otherToken2 = p2.T2.Id.String()
					} else {
						t2 = p2.T2.Snapshot()
						otherToken2 = p2.T1.Id.String()
					}

					pId := pairId(t1.String(), t2.String())
					if _, ok := mps1[pId]; ok {
						continue
					}

					t1.ET = &types.EToken{ExtraExchange: ps1.ex.NID(), OtherToken: otherToken1}
					t2.ET = &types.EToken{ExtraExchange: ps2.ex.NID(), OtherToken: otherToken2}

					mps1[pId] = &pair{p: &entity.Pair{
						T1:       t1,
						T2:       t2,
						LP:       ex.Id(),
						Exchange: ex.NID(),
						Enable:   true,
						EP:       &types.ExchangePair{},
					},
						p1:  p1,
						p2:  p2,
						ex1: ps1.ex,
						ex2: ps2.ex,
					}
				}
			}
		}
	}

	mps2 := make(map[string]*entity.Pair)
	for net1, chain1 := range tl {
		for net2, chain2 := range tl {
			if net1 == net2 {
				continue
			}
			for _, t1 := range chain1.Tokens {
				T1 := &entity.Token{
					Id: entity.TokenId{Symbol: t1.Symbol,
						Standard: t1.Standard, Network: t1.Network},
					ContractAddress: t1.TokenAddress,
					Decimals:        uint64(t1.Decimals),
					Native:          t1.Symbol == t1.Standard,
					ET:              &types.EToken{},
				}
				for _, t2 := range chain2.Tokens {
					T2 := &entity.Token{
						Id: entity.TokenId{Symbol: t2.Symbol,
							Standard: t2.Standard, Network: t2.Network},
						ContractAddress: t2.TokenAddress,
						Decimals:        uint64(t2.Decimals),
						Native:          t2.Symbol == t2.Standard,
						ET:              &types.EToken{},
					}

					pId := pairId(T1.String(), T2.String())
					if _, ok := mps2[pId]; ok {
						continue
					}

					mps2[pId] = &entity.Pair{
						T1:       T1,
						T2:       T2,
						LP:       ex.Id(),
						Exchange: ex.NID(),
						Enable:   true,
						EP:       &types.ExchangePair{},
					}
				}
			}
		}
	}

	for _, ps := range mps0 {
		for _, p0 := range ps.ps {
			for _, p2 := range mps2 {
				var (
					t1, t2     *entity.Token
					otherToken string
				)
				if p0.T2.String() == p2.T1.String() {
					t1 = p0.T1.Snapshot()
					otherToken = p0.T2.Id.String()
					t2 = p2.T2
				} else if p0.T1.String() == p2.T1.String() {
					t1 = p0.T2.Snapshot()
					otherToken = p0.T1.Id.String()
					t2 = p2.T2

				} else if p0.T2.String() == p2.T2.String() {
					t1 = p0.T1.Snapshot()
					otherToken = p0.T2.Id.String()
					t2 = p2.T1
				} else if p0.T1.String() == p2.T2.String() {
					t1 = p0.T2.Snapshot()
					otherToken = p0.T1.Id.String()
					t2 = p2.T1
				} else {
					continue
				}
				if t1.Id.Network == t2.Id.Network {
					continue
				}
				t1.ET = &types.EToken{ExtraExchange: ps.ex.NID(), OtherToken: otherToken}
				t2.ET = &types.EToken{}

				mps1[pairId(t1.String(), t2.String())] = &pair{
					p: &entity.Pair{
						T1:       t1,
						T2:       t2,
						LP:       ex.Id(),
						Exchange: ex.NID(),
						Enable:   true,
						EP:       &types.ExchangePair{},
					},
					p1:  p0,
					ex1: ps.ex,
					ex2: nil,
				}
			}
		}
	}

	ps := []*pair{}
	for _, p := range mps1 {
		if !isIn(pairId(p.p.T1.String(), p.p.T2.String()), epsStr) {
			ps = append(ps, p)
		}
	}

	for _, p := range mps2 {
		if !isIn(pairId(p.T1.String(), p.T2.String()), epsStr) {
			ps = append(ps, &pair{p: p, ex1: nil, ex2: nil})
		}
	}

	ps2 := []*entity.Pair{}
	for _, p := range ps {
		if p.p.T1.Id.Network != p.p.T2.Id.Network {
			var min1, min2 float64
			tEfa, ok := tokenEfa[p.p.T1.String()]
			if !ok {
				if p.ex1 == nil {
					tEfa = exchangeFee
					tokenEfa[p.p.T1.String()] = tEfa

				} else {
					efa, _, err := p.ex1.ExchangeFeeAmount(p.p.T1.Id, p.p1, exchangeFee)
					if err != nil {
						ex.l.Debug(agent, err.Error())
					} else {
						tEfa = efa
						tokenEfa[p.p.T1.String()] = tEfa
					}
				}
			}
			if tEfa > 0 {
				min1 = ex.minAndMax(tEfa, feeRate)
			}

			tEfa, ok = tokenEfa[p.p.T2.String()]
			if !ok {
				if p.ex2 == nil {
					tEfa = exchangeFee
					tokenEfa[p.p.T2.String()] = tEfa
				} else {
					efa, _, err := p.ex2.ExchangeFeeAmount(p.p.T2.Id, p.p2, exchangeFee)
					if err != nil {
						ex.l.Debug(agent, err.Error())
					} else {
						tEfa = efa
						tokenEfa[p.p.T2.String()] = tEfa
					}
				}
			}
			if tEfa > 0 {
				min2 = ex.minAndMax(tEfa, feeRate)
			}
			if min1 > p.p.T1.Min {
				p.p.T1.Min = min1
			}
			if min2 > p.p.T2.Min {
				p.p.T2.Min = min2
			}
			p.p.ExchangeFee = exchangeFee
			p.p.FeeRate1 = feeRate
			p.p.FeeRate2 = feeRate
			ps2 = append(ps2, p.p)
		}
	}

	return ps2, nil
}

func pairId(t1, t2 string) string {
	if t1 < t2 {
		return fmt.Sprintf("%s/%s", t1, t2)
	} else {
		return fmt.Sprintf("%s/%s", t2, t1)
	}
}

func isIn(p string, ps []string) bool {
	for _, p0 := range ps {
		if p0 == p {
			return true
		}
	}
	return false
}
