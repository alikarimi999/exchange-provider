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

func (m *Multichain) Exchange(o *entity.Order, i int) (string, error) {

	ch := m.cs[chainId(o.Routes[i].In.ChainId)]
	src := common.HexToAddress(o.Deposit.Addr)
	tx, nonce, err := m.swap(o, src, ch, i)
	if err != nil {
		if nonce != nil {
			ch.w.ReleaseNonce(src, nonce.Uint64())
		}
		return "", err
	}
	if nonce != nil {
		ch.w.BurnNonce(src, nonce.Uint64())
	}

	return tx.Hash().String(), nil
}

func (m *Multichain) swap(o *entity.Order, src common.Address, chain *Chain, i int) (tx *types.Transaction, nonce *big.Int, err error) {

	in := c2T(o.Routes[i].In)
	out := c2T(o.Routes[i].Out)
	p, err := m.pairs.get(in, out)
	if err != nil {
		return nil, nil, err
	}

	var tIn, tOut *token
	if p.t1.Symbol == in.Symbol {
		tIn = p.t1
		tOut = p.t2
	} else {
		tOut = p.t1
		tIn = p.t2
	}

	srcChain, err := strconv.Atoi(tIn.Chain)
	if err != nil {
		return nil, nil, err
	}
	destChain, err := strconv.Atoi(tOut.Chain)
	if err != nil {
		return nil, nil, err
	}

	ch := tIn.cs[len(tIn.cs)-1]
	if ch.IsApprove {
		if err != nil {
			return nil, nil, err
		}
		if errs := chain.am.InfinitApproves(ts.Token{Symbol: tIn.Symbol,
			Address: common.HexToAddress(tIn.Address), Decimals: tIn.Decimals},
			common.HexToAddress(ch.Router), src); len(errs) > 0 {
			e := errors.New("")
			for _, err := range errs {
				e = fmt.Errorf("%s\n%s", e.Error(), err.Error())
			}
			return nil, nil, e
		}
	}

	amount, err := numbers.FloatStringToBigInt(o.Swaps[i].InAmount, tIn.Decimals)
	if err != nil {
		return nil, nil, err
	}

	pr := chain.provider()
	if strings.Contains(ch.RouterABI, "Swapout") {

		c, err := contracts.NewContracts(common.HexToAddress(ch.Router), pr.Client)
		if err != nil {
			return nil, nil, err
		}
		opts, err := chain.w.NewKeyedTransactorWithChainID(src, common.Big0, int64(srcChain))
		if err != nil {
			return nil, opts.Nonce, err
		}
		tx, err := c.Swapout(opts, amount, src)
		return tx, opts.Nonce, err

	} else if strings.Contains(ch.RouterABI, "anySwapOutUnderlying") {

		c, err := contracts.NewContracts(common.HexToAddress(ch.Router), pr.Client)
		if err != nil {
			return nil, nil, err
		}
		opts, err := chain.w.NewKeyedTransactorWithChainID(src, common.Big0, int64(srcChain))
		if err != nil {
			return nil, opts.Nonce, err
		}
		tx, err := c.AnySwapOutUnderlying(opts, common.HexToAddress(ch.FromAnyToken.Address),
			src, amount, big.NewInt(int64(destChain)))
		return tx, opts.Nonce, err

	} else if strings.Contains(ch.RouterABI, "anySwapOutNative") {

		c, err := contracts.NewContracts(common.HexToAddress(ch.Router), pr.Client)
		if err != nil {
			return nil, nil, err
		}
		opts, err := chain.w.NewKeyedTransactorWithChainID(src, amount, int64(srcChain))
		if err != nil {
			return nil, opts.Nonce, err
		}
		tx, err := c.AnySwapOutNative(opts, common.HexToAddress(ch.FromAnyToken.Address),
			src, big.NewInt(int64(destChain)))
		return tx, opts.Nonce, err

	} else if strings.Contains(ch.RouterABI, "transfer") {

		c, err := contracts.NewContracts(common.HexToAddress(tIn.Address), pr.Client)
		if err != nil {
			return nil, nil, err
		}
		opts, err := chain.w.NewKeyedTransactorWithChainID(src, common.Big0, int64(srcChain))
		if err != nil {
			return nil, opts.Nonce, err
		}

		tx, err := c.Transfer(opts, common.HexToAddress(ch.DepositAddress), amount)
		return tx, opts.Nonce, err

	} else if strings.Contains(ch.RouterABI, "anySwapOut") {

		c, err := contracts.NewContracts(common.HexToAddress(ch.Router), pr.Client)
		if err != nil {
			return nil, nil, err
		}
		opts, err := chain.w.NewKeyedTransactorWithChainID(src, common.Big0, int64(srcChain))
		if err != nil {
			return nil, opts.Nonce, err
		}
		tx, err := c.AnySwapOut(opts, common.HexToAddress(ch.FromAnyToken.Address),
			src, amount, big.NewInt(int64(destChain)))
		return tx, opts.Nonce, err

	} else if strings.Contains(ch.RouterABI, "sendTransaction") {

		tx, err := utils.TransferNative(chain.w, src, common.HexToAddress(ch.DepositAddress),
			int64(srcChain), amount, pr)
		return tx, nil, err

	} else {
		return nil, nil, fmt.Errorf("unknown RouterABI '%s'", ch.RouterABI)
	}

}
