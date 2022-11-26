package multichain

import (
	"errors"
	"exchange-provider/internal/delivery/exchanges/dex/multichain/contracts"
	ts "exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/internal/delivery/exchanges/dex/utils"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/utils/numbers"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (ex *Multichain) Swap(o *entity.Order, tIn, tOut ts.Token,
	value string, source, dest common.Address) (tx *types.Transaction, nonce *big.Int, err error) {

	if ts, ok := ex.c[chainId(strconv.Itoa(int(tIn.ChainId)))]; ok {
		cs := destChains(*ts, tOut.ChainId, tIn.Address)
		if len(cs) > 0 {

			ch := cs[len(cs)-1]

			if ch.IsApprove {
				if errs := ex.am.InfinitApproves(tIn, common.HexToAddress(ch.Router),
					tIn.ChainId, source); len(errs) > 0 {
					e := errors.New("")
					for _, err := range errs {
						e = fmt.Errorf("%s\n%s", e.Error(), err.Error())
					}
					return nil, nil, e
				}
			}

			amount, err := numbers.FloatStringToBigInt(value, tIn.Decimals)
			if err != nil {
				return nil, nil, err
			}

			if strings.Contains(ch.RouterABI, "Swapout") {

				c, err := contracts.NewContracts(common.HexToAddress(ch.Router), ex.provider().Client)
				if err != nil {
					return nil, nil, err
				}
				opts, err := ex.wallet.NewKeyedTransactorWithChainID(source, common.Big0, tIn.ChainId)
				if err != nil {
					return nil, opts.Nonce, err
				}
				tx, err := c.Swapout(opts, amount, source)
				return tx, opts.Nonce, err

			} else if strings.Contains(ch.RouterABI, "anySwapOutUnderlying") {

				c, err := contracts.NewContracts(common.HexToAddress(ch.Router), ex.provider().Client)
				if err != nil {
					return nil, nil, err
				}
				opts, err := ex.wallet.NewKeyedTransactorWithChainID(source, common.Big0, tIn.ChainId)
				if err != nil {
					return nil, opts.Nonce, err
				}
				tx, err := c.AnySwapOutUnderlying(opts, common.HexToAddress(ch.FromAnyToken.Address),
					source, amount, big.NewInt(tOut.ChainId))
				return tx, opts.Nonce, err

			} else if strings.Contains(ch.RouterABI, "anySwapOutNative") {

				c, err := contracts.NewContracts(common.HexToAddress(ch.Router), ex.provider().Client)
				if err != nil {
					return nil, nil, err
				}
				opts, err := ex.wallet.NewKeyedTransactorWithChainID(source, amount, tIn.ChainId)
				if err != nil {
					return nil, opts.Nonce, err
				}

				tx, err := c.AnySwapOutNative(opts, common.HexToAddress(ch.FromAnyToken.Address),
					source, big.NewInt(tOut.ChainId))
				return tx, opts.Nonce, err

			} else if strings.Contains(ch.RouterABI, "transfer") {

				c, err := contracts.NewContracts(tIn.Address, ex.provider().Client)
				if err != nil {
					return nil, nil, err
				}
				opts, err := ex.wallet.NewKeyedTransactorWithChainID(source, common.Big0, tIn.ChainId)
				if err != nil {
					return nil, opts.Nonce, err
				}

				tx, err := c.Transfer(opts, common.HexToAddress(ch.DepositAddress), amount)
				return tx, opts.Nonce, err

			} else if strings.Contains(ch.RouterABI, "anySwapOut") {

				c, err := contracts.NewContracts(common.HexToAddress(ch.Router), ex.provider().Client)
				if err != nil {
					return nil, nil, err
				}
				opts, err := ex.wallet.NewKeyedTransactorWithChainID(source, common.Big0, tIn.ChainId)
				if err != nil {
					return nil, opts.Nonce, err
				}
				tx, err := c.AnySwapOut(opts, common.HexToAddress(ch.FromAnyToken.Address),
					source, amount, big.NewInt(tOut.ChainId))
				return tx, opts.Nonce, err

			} else if strings.Contains(ch.RouterABI, "sendTransaction") {

				tx, err := utils.TransferNative(ex.wallet, source, common.HexToAddress(ch.DepositAddress), tIn.ChainId, amount, ex.provider())
				return tx, nil, err

			} else {
				return nil, nil, fmt.Errorf("unknown RouterABI '%s'", ch.RouterABI)
			}
		}
	}
	return nil, nil, fmt.Errorf("pair not supported by multichain")
}

func destChains(ts map[tokenId]*data, destId int64, srcToken common.Address) []chain {
	c := make([]chain, 0)

	for _, ts := range ts {
		if ts.Address == srcToken.String() {
			chains := ts.DestChains[chainId(strconv.Itoa(int(destId)))]
			for _, v := range chains {
				c = append(c, v)
			}
		}
	}

	return c
}
