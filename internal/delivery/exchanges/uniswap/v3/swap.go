package uniswapv3

import (
	"log"
	"math/big"
	"order_service/pkg/utils/numbers"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (u *UniSwapV3) swap(tIn, tOut *token, value int, source, dest common.Address) (*types.Transaction, error) {
	pool, err := u.bestPool(tIn.address, tOut.address)
	if err != nil {
		return nil, err
	}

	prvKey, err := u.wallet.PrivateKey(source)
	if err != nil {
		return nil, err
	}

	amount := numbers.FloatStringToBigInt(strconv.Itoa(value), tIn.decimals)

	opts, err := bind.NewKeyedTransactorWithChainID(prvKey, u.chainId)
	if err != nil {
		return nil, err
	}

	n, err := u.wallet.Nonce(source)
	if err != nil {
		return nil, err
	}

	opts.Nonce = big.NewInt(int64(n))
	opts.Value = big.NewInt(0)
	opts.GasLimit = 0

	data := [][]byte{}
	abi, err := RouteMetaData.GetAbi()
	if err != nil {
		log.Fatal(err)
	}

	route, err := NewRoute(routerV2, u.dp)
	if err != nil {
		log.Fatal(err)
	}

	params := IV3SwapRouterExactInputSingleParams{
		TokenIn:           tIn.address,
		TokenOut:          tOut.address,
		Fee:               big.NewInt(int64(pool.feeTier)),
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
	return route.Multicall0(opts, deadline, data)
}
