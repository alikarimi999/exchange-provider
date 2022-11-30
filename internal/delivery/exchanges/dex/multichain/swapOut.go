package multichain

import (
	"exchange-provider/internal/delivery/exchanges/dex/multichain/contracts/tokenBridge"
	ts "exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/pkg/wallet/eth"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func swapOut(in, out *Token, amount *big.Int, src common.Address, pr *ts.Provider,
	w *eth.HDWallet) (*types.Transaction, *big.Int, error) {

	srcChain, err := strconv.Atoi(in.ChainId)
	if err != nil {
		return nil, nil, err
	}

	c, err := tokenBridge.NewContracts(common.HexToAddress(in.Data.Router), pr.Client)
	if err != nil {
		return nil, nil, err
	}
	opts, err := w.NewKeyedTransactorWithChainID(src, common.Big0, int64(srcChain))
	if err != nil {
		return nil, opts.Nonce, err
	}
	tx, err := c.Swapout(opts, amount, src)
	return tx, opts.Nonce, err
}
