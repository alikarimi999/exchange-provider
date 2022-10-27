package uniswapv3

import (
	"context"
	"math/big"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (u *dex) transferNative(from, to common.Address, value *big.Int) (*types.Transaction, error) {

	p := u.provider()
	var err error
	head, err := p.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	gasPrice, err := p.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	gasTipCap, err := p.SuggestGasTipCap(context.Background())
	if err != nil {
		return nil, err
	}

	gasFeeCap := new(big.Int).Add(
		gasTipCap,
		new(big.Int).Mul(head.BaseFee, big.NewInt(2)),
	)

	gas, err := p.EstimateGas(context.Background(), ethereum.CallMsg{
		From:      from,
		To:        &to,
		GasPrice:  gasPrice,
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Value:     value,
		Data:      []byte{},
	})
	if err != nil {
		return nil, err
	}

	nonce, err := u.wallet.Nonce(from)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			u.wallet.ReleaseNonce(from, nonce)
		} else {
			u.wallet.BurnNonce(from, nonce)

		}
	}()

	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   big.NewInt(int64(u.cfg.ChianId)),
		Nonce:     nonce,
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Gas:       gas,
		To:        &to,
		Value:     value,
		Data:      []byte{},
	})

	prvKey, err := u.wallet.PrivateKey(from)
	if err != nil {
		return nil, err
	}
	tx, err = types.SignTx(tx, types.NewLondonSigner(big.NewInt(int64(u.cfg.ChianId))), prvKey)
	if err != nil {
		return nil, err
	}
	err = p.SendTransaction(context.Background(), tx)
	if err != nil {
		return nil, err
	}

	return tx, nil
}
