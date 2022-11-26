package panckakeswapv2

import (
	"context"
	"exchange-provider/internal/delivery/exchanges/dex/pancakeswap/v2/contracts"
	ts "exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/internal/delivery/exchanges/dex/utils"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/logger"
	"exchange-provider/pkg/utils/numbers"
	"exchange-provider/pkg/wallet/eth"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Panckakeswapv2 struct {
	id     string
	wallet *eth.HDWallet
	router common.Address

	nt string
	tt *utils.TxTracker

	chainId int64
	abi     abi.ABI
	ps      []*ts.Provider

	l logger.Logger
}

func NewPanckakeswapV2(id, nt string, wallet *eth.HDWallet, tt *utils.TxTracker, router common.Address,
	ps []*ts.Provider, l logger.Logger) (*Panckakeswapv2, error) {
	p := &Panckakeswapv2{
		id:     id,
		wallet: wallet,
		router: router,

		nt: nt,
		tt: tt,

		ps: ps,
		l:  l,
	}

	abi, err := abi.JSON(strings.NewReader(string(contracts.EventsABI)))
	if err != nil {
		return nil, err
	}
	p.abi = abi
	c, err := p.provider().ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	p.chainId = c.Int64()
	return p, nil
}

func (p *Panckakeswapv2) Swap(o *entity.Order, tIn, tOut ts.Token, value string,
	source, dest common.Address) (*types.Transaction, *big.Int, error) {

	contract, err := contracts.NewContract(p.router, p.provider())
	if err != nil {
		return nil, nil, err
	}

	d := time.Now().Add(time.Minute * time.Duration(30)).Unix()
	amount, err := numbers.FloatStringToBigInt(value, tIn.Decimals)
	if err != nil {
		return nil, nil, err
	}
	val := common.Big0
	if tIn.IsNative() {
		val = amount
	}

	opts, err := p.wallet.NewKeyedTransactorWithChainID(source, val, p.chainId)
	if err != nil {
		return nil, opts.Nonce, err
	}

	if tIn.IsNative() {
		tx, err := contract.SwapExactETHForTokens(opts, common.Big0, []common.Address{tIn.Address, tOut.Address}, dest, big.NewInt(d))
		return tx, opts.Nonce, err

	} else if tOut.IsNative() {
		tx, err := contract.SwapExactTokensForETH(opts, amount, common.Big0, []common.Address{tIn.Address, tOut.Address}, dest, big.NewInt(d))
		if err != nil {
			return nil, opts.Nonce, err
		}
		o.Withdrawal.Unwrapped = true
		return tx, opts.Nonce, nil
	} else {
		tx, err := contract.SwapExactTokensForTokens(opts, amount, common.Big0, []common.Address{tIn.Address, tOut.Address}, dest, big.NewInt(d))
		return tx, opts.Nonce, err
	}
}
