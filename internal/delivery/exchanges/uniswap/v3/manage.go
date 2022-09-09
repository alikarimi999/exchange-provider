package uniswapv3

import (
	"fmt"
	"order_service/internal/entity"
	"time"
)

func (u *UniSwapV3) Stop() {
	op := fmt.Sprintf("%s.Stop", u.NID())
	close(u.stopCh)
	u.stoppedAt = time.Now()
	u.l.Debug(string(op), "stopped")
}

func (u *UniSwapV3) Configs() interface{} {
	return u.cfg
}

func (u *UniSwapV3) GetAllPairs() []*entity.Pair {
	ps := u.pairs.getAll()
	pairs := []*entity.Pair{}

	for _, p := range ps {
		pairs = append(pairs, p.ToEntity())
	}

	return pairs
}

func (u *UniSwapV3) StartAgain() (*entity.StartAgainResult, error) {
	return &entity.StartAgainResult{}, nil
}

func (u *UniSwapV3) GetPair(bc, qc *entity.Coin) (*entity.Pair, error) {
	if bc.ChainId != chainId || qc.ChainId != chainId {
		return nil, fmt.Errorf("unexpected chain id %v and chain id %v", bc.ChainId, qc.ChainId)
	}

	p, err := u.pairs.get(bc.CoinId, qc.CoinId)
	if err != nil {
		return nil, err
	}

	return p.ToEntity(), nil
}