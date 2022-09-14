package uniswapv3

import (
	"fmt"
	"order_service/internal/entity"
	"order_service/pkg/errors"
	"strings"
)

func (u *UniSwapV3) Support(bc, qc *entity.Coin) bool {
	if bc.ChainId != chainId || qc.ChainId != chainId {
		return false
	}
	_, err := u.pairs.get(bc.CoinId, qc.CoinId)
	return err == nil
}

func (u *UniSwapV3) RemovePair(bc, qc *entity.Coin) error {
	if bc.ChainId != chainId || qc.ChainId != chainId {
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

func (u *UniSwapV3) GetAddress(c *entity.Coin) (*entity.Address, error) {
	if c.ChainId != chainId {
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
