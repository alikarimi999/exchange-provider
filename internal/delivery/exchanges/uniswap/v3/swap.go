package uniswapv3

import (
	"fmt"
	"math/big"
	"order_service/internal/delivery/exchanges/uniswap/v3/contracts"
	"order_service/pkg/utils/numbers"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (u *UniSwapV3) swap(tIn, tOut token, value string, source, dest common.Address) (*types.Transaction, *pair, error) {
	agent := u.agent("swap")
	pool, err := u.setBestPrice(tIn, tOut)
	if err != nil {
		return nil, nil, err
	}

	amount, err := numbers.FloatStringToBigInt(value, tIn.Decimals)
	if err != nil {
		return nil, nil, err
	}

	val := big.NewInt(0)
	if tIn.isNative() {
		val = amount
	}
	opts, err := u.newKeyedTransactorWithChainID(source, val)
	if err != nil {
		return nil, nil, err
	}

	data := [][]byte{}
	abi, err := contracts.RouteMetaData.GetAbi()
	if err != nil {
		return nil, nil, err
	}

	route, err := contracts.NewRoute(routerV2, u.provider)
	if err != nil {
		return nil, nil, err
	}

	params := contracts.IV3SwapRouterExactInputSingleParams{
		TokenIn:           tIn.Address,
		TokenOut:          tOut.Address,
		Fee:               pool.feeTier,
		Recipient:         dest,
		AmountIn:          amount,
		AmountOutMinimum:  big.NewInt(0),
		SqrtPriceLimitX96: big.NewInt(0),
	}
	es, err := abi.Pack("exactInputSingle", params)
	if err != nil {
		return nil, nil, err
	}

	data = append(data, es)

	deadline := big.NewInt(time.Now().Add(time.Minute * time.Duration(30)).Unix())
	tx, err := route.Multicall0(opts, deadline, data)
	if err != nil {
		u.wallet.ReleaseNonce(source, opts.Nonce.Uint64())
		return nil, nil, err
	}
	u.wallet.BurnNonce(source, tx.Nonce())

	u.l.Debug(agent, fmt.Sprintf("swap `%+v`", params))
	return tx, pool, err
}
