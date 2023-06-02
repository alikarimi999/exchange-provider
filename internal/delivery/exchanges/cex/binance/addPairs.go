package binance

import (
	"exchange-provider/internal/delivery/exchanges/cex/binance/dto"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/adshao/go-binance/v2"
)

func (ex *exchange) AddPairs(data interface{}) (*entity.AddPairsResult, error) {
	agent := fmt.Sprintf("%s.AddPairs", ex.NID())
	req, ok := data.(*dto.AddPairsRequest)
	if !ok {
		return nil, errors.Wrap(agent, errors.ErrBadRequest)
	}

	res := &entity.AddPairsResult{}
	ps := []*entity.Pair{}

	wg := &sync.WaitGroup{}
	mux := sync.Mutex{}
	waitChan := make(chan struct{}, max_conccurrent_jobs)

	for _, p := range req.Pairs {
		p.BC.ToUpper()
		p.QC.ToUpper()
		if ex.pairs.Exists(ex.Id(), p.BC.String(), p.QC.String()) {
			res.Existed = append(res.Existed, p.String())
			continue
		}

		bc, ok := ex.cl.getCoin(p.BC.ET.Coin, p.BC.ET.Network)
		if !ok {
			res.Failed = append(res.Failed, &entity.PairsErr{Pair: p.String(),
				Err: fmt.Errorf("token with '%s-%s' not found in binance",
					p.BC.ET.Coin, p.BC.ET.Network)})
			continue
		}

		if !bc.DepositEnable {
			res.Failed = append(res.Failed, &entity.PairsErr{Pair: p.String(),
				Err: fmt.Errorf("token with '%s-%s' deposit is disable in binance",
					p.BC.ET.Coin, p.BC.ET.Network)})
			continue
		}

		if !bc.WithdrawEnable {
			res.Failed = append(res.Failed, &entity.PairsErr{Pair: p.String(),
				Err: fmt.Errorf("token with '%s-%s' withdraw is disable in binance",
					p.BC.ET.Coin, p.BC.ET.Network)})
			continue
		}

		qc, ok := ex.cl.getCoin(p.QC.ET.Coin, p.QC.ET.Network)
		if !ok {
			res.Failed = append(res.Failed, &entity.PairsErr{Pair: p.String(),
				Err: fmt.Errorf("token with '%s-%s' not found in binance",
					p.QC.ET.Coin, p.QC.ET.Network)})
			continue
		}

		if !qc.DepositEnable {
			res.Failed = append(res.Failed, &entity.PairsErr{Pair: p.String(),
				Err: fmt.Errorf("token with '%s-%s' deposit is disable in binance",
					p.QC.ET.Coin, p.QC.ET.Network)})
			continue
		}

		if !qc.WithdrawEnable {
			res.Failed = append(res.Failed, &entity.PairsErr{Pair: p.String(),
				Err: fmt.Errorf("token with '%s-%s' withdraw is disable in binance",
					p.QC.ET.Coin, p.QC.ET.Network)})
			continue
		}

		fn := func(t dto.Token, n binance.Network) (entity.ExchangeToken, error) {
			bt, err := time.ParseDuration(t.BlockTime)
			if err != nil {
				return nil, err
			}
			if t.StableToken == "" {
				return nil, fmt.Errorf("stableToken cannot be empty")
			}

			token := &Token{
				Coin:        t.Coin,
				Network:     t.Network,
				StableToken: t.StableToken,
				BlockTime:   bt,
			}
			token.setInfos(n)
			return token, nil
		}

		ep, err := p.ToEntity(fn, bc, qc)
		if err != nil {
			res.Failed = append(res.Failed, &entity.PairsErr{
				Pair: p.String(),
				Err:  err,
			})
			continue
		}

		ep.EP = &ExchangePair{}
		if p.IC != "" {
			ic := Token{Coin: strings.ToUpper(p.IC)}
			ep.EP = &ExchangePair{
				HasIntermediaryCoin: true,
				IC1:                 ic.Snapshot().(*Token),
				IC2:                 ic.Snapshot().(*Token),
			}
		}

		waitChan <- struct{}{}
		wg.Add(1)
		go func(p *entity.Pair) {
			defer func() {
				<-waitChan
				wg.Done()
			}()
			if err := ex.infos(ep); err != nil {
				mux.Lock()
				res.Failed = append(res.Failed, &entity.PairsErr{Pair: p.String(), Err: err})
				mux.Unlock()
				return
			}

			mux.Lock()
			ps = append(ps, ep)
			res.Added = append(res.Added, p.String())
			mux.Unlock()
		}(ep)
	}
	wg.Wait()
	return res, ex.pairs.Add(ex, ps...)
}

func (k *exchange) RemovePair(t1, t2 entity.TokenId) error {
	return k.pairs.Remove(k.Id(), t1.String(), t2.String(), true)
}
