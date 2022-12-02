package multichain

type ChainId string

// type pairId string
// type tokenId string
// type data struct {
// 	Token
// 	DestChains map[chainId]map[pairId]Data `json:"destChains`
// }
// type tokens = map[tokenId]*data

// func getinfos(cs map[chainId]*tokens, t1Symbol, t1Chain, t2Symbol, t2Chain string) (T1, T2 *Token) {

// 	t1 := &Token{CoinId: t1Symbol, ChainId: t1Chain}
// 	t2 := &Token{CoinId: t2Symbol, ChainId: t2Chain}

// 	wg := &sync.WaitGroup{}

// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		ts, ok := cs[chainId(t1Chain)]
// 		if !ok {
// 			return
// 		}

// 		for _, t := range *ts {
// 			if t.CoinId == t1Symbol {
// 				ts2, ok := t.DestChains[chainId(t2Chain)]
// 				if !ok {
// 					return
// 				}

// 				*t1 = t.Token
// 				for _, c := range ts2 {
// 					t1.cs = append(t1.cs, &c)
// 				}
// 			}
// 		}
// 	}()

// 	wg.Add(1)
// 	go func() {
// 		wg.Done()
// 		ts, ok := cs[chainId(t2Chain)]
// 		if !ok {
// 			return
// 		}
// 		for _, t := range *ts {
// 			if t.CoinId == t2Symbol {

// 				ts2, ok := t.DestChains[chainId(t1Chain)]
// 				if !ok {
// 					return
// 				}

// 				*t2 = t.Token
// 				for _, c := range ts2 {
// 					t2.cs = append(t2.cs, &c)
// 				}
// 			}
// 		}
// 	}()
// 	wg.Wait()
// 	return t1, t2
// }
