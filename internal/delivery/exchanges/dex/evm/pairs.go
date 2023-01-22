package evm

import (
	"exchange-provider/internal/delivery/exchanges/dex/evm/dto"
	"exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

var errPairNotSupport = fmt.Errorf("pair not supported")

func (d *EvmDex) AddPairs(data interface{}) (*entity.AddPairsResult, error) {
	req := data.(*dto.AddPairsRequest)

	res := &entity.AddPairsResult{}
	mux := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	for _, p := range req.Pairs {
		if d.pairs.Exists(d.Id(), p.T1, p.T2) {
			mux.Lock()
			res.Existed = append(res.Existed, p.String())
			mux.Unlock()
			continue
		}

		wg.Add(1)
		go func(p *dto.Pair) {
			defer wg.Done()
			in := types.Token{
				Symbol:   p.T1.TokenId,
				Address:  common.HexToAddress(p.T1.Address),
				Decimals: int(p.T1.Decimals),
				ChainId:  d.ChainId,
			}
			out := types.Token{
				Symbol:   p.T2.TokenId,
				Address:  common.HexToAddress(p.T2.Address),
				Decimals: int(p.T2.Decimals),
				ChainId:  d.ChainId,
			}
			pair, err := d.Pair(in, out)
			mux.Lock()
			defer mux.Unlock()
			if err != nil {
				res.Failed = append(res.Failed, &entity.PairsErr{Pair: pair.String(), Err: err})
				return
			}
			ep := pair.ToEntity(d.Id(), d.NativeToken, d.TokenStandard)
			d.pairs.Add(d, ep)
			res.Added = append(res.Added, *ep)
		}(p)
	}
	wg.Wait()
	return res, nil
}

func (d *EvmDex) Price(t1, t2 *entity.Token) (*entity.Pair, error) {
	if t1.ChainId != d.TokenStandard || t2.ChainId != d.TokenStandard {
		return nil, errPairNotSupport
	}

	T1, ok := d.get(t1.TokenId)
	if !ok {
		return nil, errors.Wrap(errors.ErrNotFound)
	}
	T2, ok := d.get(t2.TokenId)
	if !ok {
		return nil, errors.Wrap(errors.ErrNotFound)
	}

	p, err := d.Pair(T1, T2)
	if err != nil {
		return nil, err
	}

	if T1.IsNative() {
		p.T1.Symbol = d.NativeToken
	}
	if T2.IsNative() {
		p.T2.Symbol = d.NativeToken
	}
	return p.ToEntity(d.Id(), d.NativeToken, d.TokenStandard), nil
}

// func (d *EvmDex) GetAllPairs() []*entity.Pair {
// 	agent := d.agent("GetAllPairs")
// 	pairs := d.pairs.getAll()
// 	ps := []*entity.Pair{}
// 	psMux := &sync.Mutex{}
// 	guard := make(chan struct{}, 20)

// 	wg := &sync.WaitGroup{}
// 	for i, p := range pairs {
// 		guard <- struct{}{}
// 		wg.Add(1)
// 		go func(pair types.Pair, i int) {
// 			defer func() {
// 				<-guard
// 				wg.Done()
// 			}()

// 			p, err := d.Pair(pair.T1, pair.T2)
// 			if err != nil {
// 				d.l.Debug(agent, err.Error())
// 				return
// 			}
// 			psMux.Lock()
// 			ps = append(ps, p)
// 			psMux.Unlock()
// 		}(p, i)
// 	}
// 	wg.Wait()
// 	return ps
// }

func (d *EvmDex) Support(t1, t2 *entity.Token) bool {
	if t1.ChainId != d.TokenStandard || t2.ChainId != d.TokenStandard {
		return false
	}
	return d.pairs.Exists(d.Id(), t1, t2)
}

func (d *EvmDex) RemovePair(t1, t2 *entity.Token) error {
	if t1.ChainId != d.TokenStandard || t2.ChainId != d.TokenStandard {
		return errors.New("pair not found")
	}
	return d.pairs.Remove(d.Id(), t1, t2)
}
