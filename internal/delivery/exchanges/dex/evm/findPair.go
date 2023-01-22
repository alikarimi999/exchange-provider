package evm

import (
	"exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/pkg/errors"
	"fmt"
	"strings"
	"sync"
)

func (d *EvmDex) findAllPairs() {
	agent := d.agent("findAllPairs")
	guard := make(chan struct{}, 50)
	proccessed := []string{}
	pMux := &sync.Mutex{}

	for _, t1 := range d.ts.Tokens {
		if t1.ChainId != int64(d.ChainId) {
			continue
		}

		go func(t1 types.Token) {

		start:
			for _, t2 := range d.ts.Tokens {
				if t2.ChainId != int64(d.ChainId) || t2.Symbol == t1.Symbol {
					continue
				}

				p := pairId(t1.Symbol, t2.Symbol)
				pMux.Lock()
				for _, pr := range proccessed {
					if pr == p {
						pMux.Unlock()
						continue start
					}
				}
				proccessed = append(proccessed, p)
				pMux.Unlock()

				guard <- struct{}{}
				go func(t2 types.Token) {
					defer func() {
						<-guard
					}()

					p, err := d.Pair(t1, t2)
					if err != nil {
						if errors.ErrorCode(err) == errors.ErrNotFound {
							return
						}
						d.l.Error(agent, err.Error())
						return
					}
					ep := p.ToEntity(d.Id(), d.NativeToken, d.TokenStandard)
					ep.Price1 = ""
					ep.Price2 = ""
					d.pairs.Add(d, ep.Snapshot())

					if ep.T1.TokenId == d.WrappedNativeToken {
						ep.T1.TokenId = d.NativeToken
						ep.T1.Native = true
						d.pairs.Add(d, ep.Snapshot())
					} else if ep.T2.TokenId == d.WrappedNativeToken {
						ep.T2.TokenId = d.NativeToken
						ep.T2.Native = true
						d.pairs.Add(d, ep.Snapshot())
					}

				}(t2)
			}
		}(t1)
	}
}

func pairId(t1, t2 string) string {
	if strings.Compare(t1, t2) == -1 {
		return fmt.Sprintf("%s%s%s", t1, types.Delimiter, t2)
	}
	return fmt.Sprintf("%s%s%s", t2, types.Delimiter, t1)
}
