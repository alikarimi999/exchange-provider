package uniswapv3

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
	"strings"
)

func (u *dex) Support(bc, qc *entity.Coin) bool {
	if bc.ChainId != u.cfg.TokenStandard || qc.ChainId != u.cfg.TokenStandard {
		return false
	}
	_, err := u.pairs.get(bc.CoinId, qc.CoinId)
	return err == nil
}

func (u *dex) RemovePair(bc, qc *entity.Coin) error {
	if bc.ChainId != u.cfg.TokenStandard || qc.ChainId != u.cfg.TokenStandard {
		return errors.Wrap(errors.ErrNotFound, errors.NewMesssage("pair not found"))
	}

	if u.pairs.existsExactly(bc.CoinId, qc.CoinId) {
		id := pairId(bc.CoinId, qc.CoinId)
		delete(u.v.Get(fmt.Sprintf("%s.pairs", u.NID())).(map[string]interface{}), strings.ToLower(id))
		if err := u.v.WriteConfig(); err != nil {
			return err
		}
		return u.pairs.remove(bc.CoinId, qc.CoinId)
	}
	return errors.Wrap(errors.ErrNotFound, errors.NewMesssage("pair not found"))
}

func (u *dex) GetAddress(c *entity.Coin) (*entity.Address, error) {
	if c.ChainId != u.cfg.TokenStandard {
		return nil, errors.Wrap(errors.ErrBadRequest)
	}

	_, err := u.tokens.get(c.CoinId)
	if err != nil {
		return nil, err
	}

	addr, err := u.wallet.RandAddress()
	if err != nil {
		return nil, err
	}
	return &entity.Address{Addr: addr.String()}, nil
}
