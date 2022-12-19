package multichain

import (
	"exchange-provider/internal/delivery/exchanges/dex/multichain/contracts/AnyswapV6Router"
	"exchange-provider/internal/delivery/exchanges/dex/multichain/contracts/MultichainV7Router"
	ts "exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/pkg/wallet/eth"

	"fmt"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func anySwapOut(in, out *Token, amount *big.Int, src common.Address,
	pr *ts.EthProvider, w *eth.HDWallet) (tx *types.Transaction, nonce *big.Int, err error) {

	srcChain, err := strconv.Atoi(in.ChainId)
	if err != nil {
		return nil, nil, err
	}
	destChain, err := strconv.Atoi(out.ChainId)
	if err != nil {
		return nil, nil, err
	}

	switch in.Data.RouterName {
	case "MultichainV7Router":
		c, err := MultichainV7Router.NewContract(common.HexToAddress(in.Data.Router), pr.Client)
		if err != nil {
			return nil, nil, err
		}

		opts, err := w.NewKeyedTransactorWithChainID(src, common.Big0, int64(srcChain))
		if err != nil {
			return nil, opts.Nonce, err
		}

		tx, err := c.AnySwapOut(opts, common.HexToAddress(in.Address),
			src.String(), amount, big.NewInt(int64(destChain)))
		return tx, opts.Nonce, err

	case "AnyswapV6Router", "AnyswapV3Router", "AnyswapV4Router":
		c, err := AnyswapV6Router.NewContracts(common.HexToAddress(in.Data.Router), pr.Client)
		if err != nil {
			return nil, nil, err
		}

		opts, err := w.NewKeyedTransactorWithChainID(src, common.Big0, int64(srcChain))
		if err != nil {
			return nil, opts.Nonce, err
		}

		tx, err := c.AnySwapOut(opts, common.HexToAddress(in.Address),
			src, amount, big.NewInt(int64(destChain)))
		return tx, opts.Nonce, err
	default:
		return nil, nil, fmt.Errorf("unknown RouterName: %s", in.Data.RouterName)
	}

}
