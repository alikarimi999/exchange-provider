package multichain

import "sync"

type chainId string
type pairId string
type tokenId string
type data struct {
	token
	DestChains map[chainId]map[pairId]pairInfo `json:"destChains`
}
type tokens = map[tokenId]*data

type pairInfo struct {
	token
	AnyToken       token  `json:"anytoken"`
	FromAnyToken   token  `json:"fromanytoken"`
	Router         string `json:"router"`
	RouterABI      string `json:"routerABI"`
	DepositAddress string `json:"DepositAddress"`
	IsApprove      bool   `json:"isApprove"`
}

func getinfos(cs map[chainId]*tokens, t1Symbol, t1Chain, t2Symbol, t2Chain string) (T1, T2 *token) {

	t1 := &token{Symbol: t1Symbol, Chain: t1Chain}
	t2 := &token{Symbol: t2Symbol, Chain: t2Chain}

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		ts, ok := cs[chainId(t1Chain)]
		if !ok {
			return
		}

		for _, t := range *ts {
			if t.Symbol == t1Symbol {
				ts2, ok := t.DestChains[chainId(t2Chain)]
				if !ok {
					return
				}

				*t1 = t.token
				for _, c := range ts2 {
					t1.cs = append(t1.cs, &c)
				}
			}
		}
	}()

	wg.Add(1)
	go func() {
		wg.Done()
		ts, ok := cs[chainId(t2Chain)]
		if !ok {
			return
		}
		for _, t := range *ts {
			if t.Symbol == t2Symbol {

				ts2, ok := t.DestChains[chainId(t1Chain)]
				if !ok {
					return
				}

				*t2 = t.token
				for _, c := range ts2 {
					t2.cs = append(t2.cs, &c)
				}
			}
		}
	}()
	wg.Wait()
	return t1, t2
}
