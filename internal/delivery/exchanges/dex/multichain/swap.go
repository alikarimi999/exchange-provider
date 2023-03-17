package multichain

import (
	"errors"
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

func (m *Multichain) Swap(o *entity.CexOrder, i int) (string, error) {
	agent := "multichain.Swap"

	ch := m.cs[ChainId(o.Routes[i].In.ChainId)]
	src := common.HexToAddress(o.Deposit.Address.Addr)
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

	m.l.Debug(agent, fmt.Sprintf("order %d tx %s", o.Id, tx.Hash().String()))
	return tx.Hash().String(), nil
}

func (m *Multichain) swap(o *entity.CexOrder, src common.Address, chain *Chain,
	i int) (tx *types.Transaction, nonce *big.Int, err error) {

	in := c2T(o.Routes[i].In)
	out := c2T(o.Routes[i].Out)
	p, err := m.pairs.get(in, out)
	if err != nil {
		return nil, nil, err
	}

	var tIn, tOut *Token
	if p.T1.ChainId == in.ChainId {
		tIn = p.T1
		tOut = p.T2
	} else {
		tOut = p.T1
		tIn = p.T2
	}

	srcChain, err := strconv.Atoi(tIn.ChainId)
	if err != nil {
		return nil, nil, err
	}

	d := tIn.Data
	if d.IsApprove {
		if err != nil {
			return nil, nil, err
		}
		if errs := chain.am.InfinitApproves(ts.Token{Symbol: tIn.TokenId,
			Address: common.HexToAddress(tIn.Address), Decimals: tIn.Decimals},
			common.HexToAddress(d.Router), src); len(errs) > 0 {
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
	if strings.Contains(d.RouterABI, "Swapout") {
		return swapOut(tIn, tOut, amount, src, pr, chain.w)

	} else if strings.Contains(d.RouterABI, "anySwapOutUnderlying") {
		return anySwapOutUnderlying(tIn, tOut, amount, src, pr, chain.w)

	} else if strings.Contains(d.RouterABI, "anySwapOutNative") {
		return anySwapOutNative(tIn, tOut, amount, src, pr, chain.w)

	} else if strings.Contains(d.RouterABI, "transfer") {
		return transfer(tIn, tOut, amount, src, pr, chain.w)

	} else if strings.Contains(d.RouterABI, "anySwapOut") {
		return anySwapOut(tIn, tOut, amount, src, pr, chain.w)

	} else if strings.Contains(d.RouterABI, "sendTransaction") {
		tx, err := utils.TransferNative(chain.w, src, common.HexToAddress(d.DepositAddress),
			int64(srcChain), amount, pr)
		return tx, nil, err

	}

	return nil, nil, fmt.Errorf("unknown RouterABI '%s'", d.RouterABI)
}
