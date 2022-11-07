package uniswapv3

import (
	ts "exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/internal/delivery/exchanges/dex/uniswap/v3/contracts"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/utils/numbers"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (u *UniswapV3) Swap(o *entity.UserOrder, tIn, tOut ts.Token, value string, source, dest common.Address) (*types.Transaction, *big.Int, error) {

	var err error
	pair, err := u.Pair(tIn, tOut)
	if err != nil {
		return nil, nil, err
	}

	amount, err := numbers.FloatStringToBigInt(value, tIn.Decimals)
	if err != nil {
		return nil, nil, err
	}

	val := big.NewInt(0)
	if tIn.IsNative() {
		val = amount
	}
	opts, err := u.wallet.NewKeyedTransactorWithChainID(source, val, u.chaindId)
	if err != nil {
		return nil, nil, err
	}

	defer func() {
		if err != nil {
			u.wallet.ReleaseNonce(source, opts.Nonce.Uint64())
		} else {
			u.wallet.BurnNonce(source, opts.Nonce.Uint64())

		}
	}()

	data := [][]byte{}
	abi, err := contracts.RouteMetaData.GetAbi()
	if err != nil {
		return nil, opts.Nonce, err
	}

	route, err := contracts.NewRoute(u.router, u.provider())
	if err != nil {
		return nil, opts.Nonce, err
	}

	params := contracts.IV3SwapRouterExactInputSingleParams{
		TokenIn:           tIn.Address,
		TokenOut:          tOut.Address,
		Fee:               pair.FeeTier,
		Recipient:         dest,
		AmountIn:          amount,
		AmountOutMinimum:  big.NewInt(0),
		SqrtPriceLimitX96: big.NewInt(0),
	}
	es, err := abi.Pack("exactInputSingle", params)
	if err != nil {
		return nil, opts.Nonce, err
	}

	data = append(data, es)

	deadline := big.NewInt(time.Now().Add(time.Minute * time.Duration(30)).Unix())
	tx, err := route.Multicall0(opts, deadline, data)
	if err != nil {
		return nil, opts.Nonce, err
	}

	return tx, opts.Nonce, err
}
