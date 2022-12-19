package dex

import (
	"context"
	ts "exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/pkg/wallet/eth"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func transferNative(w *eth.HDWallet, from, to common.Address,
	chainId int64, value *big.Int, p *ts.EthProvider) (*types.Transaction, error) {

	head, err := p.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	var tx *types.Transaction
	var nonce *big.Int
	if head.BaseFee != nil {
		tx, nonce, err = sendDynamicTx(w, from, to, chainId, value, head, p)
	} else {
		// Chain is not London ready -> use legacy transaction
		tx, nonce, err = sendLegacyTx(w, from, to, chainId, value, p)
	}

	if err != nil {
		if nonce != nil {
			w.ReleaseNonce(from, nonce.Uint64())
		}
	} else {
		w.BurnNonce(from, nonce.Uint64())
	}

	return tx, nil
}

func sendLegacyTx(w *eth.HDWallet, from, to common.Address, chainId int64,
	value *big.Int, p *ts.EthProvider) (*types.Transaction, *big.Int, error) {

	gasPrice, err := p.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, nil, err
	}

	gas, err := p.EstimateGas(context.Background(), ethereum.CallMsg{
		From:      from,
		To:        &to,
		GasPrice:  gasPrice,
		GasFeeCap: nil,
		GasTipCap: nil,
		Value:     value,
		Data:      []byte{},
	})
	if err != nil {
		return nil, nil, err
	}

	n, err := w.Nonce(from)
	if err != nil {
		return nil, nil, err
	}
	bn := big.NewInt(int64(n))

	tx := types.NewTx(&types.LegacyTx{
		To:       &to,
		Nonce:    n,
		GasPrice: gasPrice,
		Gas:      gas,
		Value:    value,
		Data:     []byte{},
	})

	prvKey, err := w.PrivateKey(from)
	if err != nil {
		return nil, bn, err
	}
	tx, err = types.SignTx(tx, types.LatestSignerForChainID(big.NewInt(chainId)), prvKey)
	if err != nil {
		return nil, bn, err
	}

	err = p.SendTransaction(context.Background(), tx)
	if err != nil {
		return nil, bn, err
	}

	return tx, bn, nil

}

func sendDynamicTx(wallet *eth.HDWallet, from, to common.Address, chainId int64, value *big.Int,
	head *types.Header, p *ts.EthProvider) (*types.Transaction, *big.Int, error) {

	gasPrice, err := p.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, nil, err
	}

	gasTipCap, err := p.SuggestGasTipCap(context.Background())
	if err != nil {
		return nil, nil, err
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
		return nil, nil, err
	}

	n, err := wallet.Nonce(from)
	if err != nil {
		return nil, nil, err
	}
	bn := big.NewInt(int64(n))

	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   big.NewInt(chainId),
		Nonce:     n,
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Gas:       gas,
		To:        &to,
		Value:     value,
		Data:      []byte{},
	})

	prvKey, err := wallet.PrivateKey(from)
	if err != nil {
		return nil, bn, err
	}

	signer := types.LatestSignerForChainID(big.NewInt(chainId))

	signature, err := crypto.Sign(signer.Hash(tx).Bytes(), prvKey)
	if err != nil {
		return nil, bn, err
	}
	tx, err = tx.WithSignature(signer, signature)
	if err != nil {
		return nil, bn, err
	}

	err = p.SendTransaction(context.Background(), tx)
	if err != nil {
		return nil, bn, err
	}

	return tx, bn, nil
}
