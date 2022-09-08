package uniswapv3

import (
	"order_service/internal/entity"
	"order_service/pkg/errors"
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
		return errors.Wrap(errors.ErrNotFound)
	}

	return u.pairs.remove(bc.CoinId, qc.CoinId)
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
