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

func (ex *exchange) createPairs() error {
	agent := ex.agent("createPairs")
	exs := ex.exs.GetAll()
	mps0 := make(map[string]*pairs)
	for _, e := range exs {
		if e.Id() != ex.Id() {
			if e.IsEnable() && e.Type() == entity.EvmDEX {
				ps := ex.pairs.GetAll(e.Id())
				mps0[e.NID()] = &pairs{ex: e.(entity.EVMDex)}
				for _, p := range ps {
					if p.Enable && (ex.tl.isTokenExists(p.T1.Id) || ex.tl.isTokenExists(p.T2.Id)) {
						mps0[e.NID()].ps = append(mps0[e.NID()].ps, p)
					}
				}
			}
		}
	}

	eps := ex.pairs.GetAll(ex.Id())
	epsStr := []string{}

	for _, ep := range eps {
		epsStr = append(epsStr, pairId(ep.T1.String(), ep.T2.String()))
	}

	tokenEfa := make(map[string]float64)
	mps1 := make(map[string]*entity.Pair)
	for ex1, ps1 := range mps0 {
		for ex2, ps2 := range mps0 {
			if ex1 == ex2 {
				continue
			}

			for _, p1 := range ps1.ps {
				for _, p2 := range ps2.ps {
					var (
						t1, t2 *entity.Token
					)

					if !ex.tl.isTokenExists(p1.T1.Id) {
						t1 = p1.T1
					} else {
						t1 = p1.T2
					}

					if !ex.tl.isTokenExists(p2.T1.Id) {
						t2 = p2.T1
					} else {
						t2 = p2.T2
					}

					pId := pairId(t1.String(), t2.String())
					if _, ok := mps1[pId]; ok {
						continue
					}

					min, err := ex.minAndMax(t1, p1, ps1.ex, tokenEfa)
					if err != nil {
						ex.l.Debug(agent, err.Error())
						continue
					}
					if min > t1.Min {
						t1.Min = min
					}

					min, err = ex.minAndMax(t2, p2, ps2.ex, tokenEfa)
					if err != nil {
						ex.l.Debug(agent, err.Error())
						continue
					}
					if min > t2.Min {
						t2.Min = min
					}

					mps1[pId] = &entity.Pair{
						T1:       t1,
						T2:       t2,
						LP:       ex.Id(),
						Exchange: ex.NID(),
						Enable:   true,
						EP:       &types.ExchangePair{},
					}
				}
			}
		}
	}

	mps2 := make(map[string]*entity.Pair)
	for net1, chain1 := range ex.tl {
		for net2, chain2 := range ex.tl {
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
					if _, ok := mps1[pId]; ok {
						continue
					}

					min := ex.cfg.ExchangeFee + (ex.cfg.ExchangeFee * ex.cfg.FeeRate * 2)
					min = min + (min * 0.1)
					T1.Min = min
					T2.Min = min

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
				var t1, t2 *entity.Token
				if p0.T2.String() == p2.T1.String() {
					t1 = p0.T1
					t2 = p2.T1
				} else if p0.T1.String() == p2.T1.String() {
					t1 = p0.T2
					t2 = p2.T2
				} else if p0.T2.String() == p2.T2.String() {
					t1 = p0.T1
					t2 = p2.T1
				} else if p0.T1.String() == p2.T2.String() {
					t1 = p0.T2
					t2 = p2.T1
				} else {
					continue
				}
				t1.ET = &types.EToken{}
				t2.ET = &types.EToken{}

				min := ex.cfg.ExchangeFee + (ex.cfg.ExchangeFee * ex.cfg.FeeRate * 2)
				min = min + (min * 0.1)
				t2.Min = min
				pid := pairId(t1.String(), t2.String())
				mps1[pid] = &entity.Pair{
					T1:       t1,
					T2:       t2,
					LP:       ex.Id(),
					Exchange: ex.NID(),
					Enable:   true,
					EP:       &types.ExchangePair{},
				}
			}
		}
	}

	ps := []*entity.Pair{}
	for _, p := range mps1 {
		if !isIn(pairId(p.T1.String(), p.T2.String()), epsStr) {
			ps = append(ps, p)
		}
	}

	for _, p := range mps2 {
		if !isIn(pairId(p.T1.String(), p.T2.String()), epsStr) {
			ps = append(ps, p)
		}
	}

	ps2 := []*entity.Pair{}
	for _, p := range ps {
		p.T1.Min *= 2
		p.T2.Min *= 2

		if p.T1.Id.Network != p.T2.Id.Network {
			ps2 = append(ps2, p)
			p.ExchangeFee = ex.cfg.ExchangeFee
			p.FeeRate1 = ex.cfg.FeeRate
			p.FeeRate2 = ex.cfg.FeeRate
		}
	}

	if len(ps2) > 0 {
		return ex.pairs.Add(ex, ps2...)
	}
	return nil
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
