package uniswapv3

import (
	"log"
	"math/big"
	"order_service/internal/delivery/exchanges/uniswap/v3/contracts"
	"order_service/pkg/utils/numbers"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (u *UniSwapV3) swap(tIn, tOut *token, value string, source, dest common.Address) (*types.Transaction, error) {
	pool, err := u.bestPool(tIn, tOut)
	if err != nil {
		return nil, err
	}

	amount, err := numbers.FloatStringToBigInt(value, tIn.decimals)
	if err != nil {
		return nil, err
	}

	val := big.NewInt(0)
	if tIn.isNative {
		val = amount
	}
	opts, err := u.newKeyedTransactorWithChainID(source, val)
	if err != nil {
		return nil, err
	}

	data := [][]byte{}
	abi, err := contracts.RouteMetaData.GetAbi()
	if err != nil {
		log.Fatal(err)
	}

	route, err := contracts.NewRoute(routerV2, u.dp)
	if err != nil {
		log.Fatal(err)
	}

	params := contracts.IV3SwapRouterExactInputSingleParams{
		TokenIn:           tIn.address,
		TokenOut:          tOut.address,
		Fee:               pool.feeTier,
		Recipient:         dest,
		AmountIn:          amount,
		AmountOutMinimum:  big.NewInt(0),
		SqrtPriceLimitX96: big.NewInt(0),
	}
	es, err := abi.Pack("exactInputSingle", params)
	if err != nil {
		log.Fatal(err)
	}

	data = append(data, es)

	deadline := big.NewInt(time.Now().Add(time.Minute * time.Duration(30)).Unix())
	tx, err := route.Multicall0(opts, deadline, data)
	if err != nil {
		u.wallet.ReleaseNonce(source, opts.Nonce.Uint64())
		return nil, err
	}
	u.wallet.BurnNonce(source, tx.Nonce())
	return tx, err
}
