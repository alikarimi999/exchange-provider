package panckakeswapv2

import (
	ts "exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/internal/delivery/exchanges/dex/utils"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/utils/numbers"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (p *Panckakeswapv2) ParseSwapLogs(o *entity.UserOrder, tx *types.Transaction, pair *ts.Pair, r *types.Receipt) (amount, fee string, err error) {

	var vol *big.Int
	if (o.Side == entity.SideBuy && pair.BT.IsNative() ||
		o.Side == entity.SideSell && pair.QT.IsNative()) &&
		r.Logs[len(r.Logs)-1].Topics[0] == p.abi.Events["Withdrawal"].ID {

		a, err := p.abi.Events["Withdrawal"].Inputs.Unpack(r.Logs[len(r.Logs)-1].Data)
		if err != nil {
			return "", "", err
		}
		vol = a[0].(*big.Int)
	} else if r.Logs[len(r.Logs)-1].Topics[0] == p.abi.Events["Swap"].ID {
		amounts, err := p.abi.Events["Swap"].Inputs.Unpack(r.Logs[len(r.Logs)-1].Data)
		if err != nil {
			return "", "", err
		}
		if amounts[2].(*big.Int).Cmp(common.Big0) == 1 {
			vol = amounts[2].(*big.Int)
		} else {
			vol = amounts[3].(*big.Int)
		}
	} else {
		return "", "", fmt.Errorf("unable to parse tx logs")
	}

	var decimals int
	if o.Side == entity.SideBuy {
		decimals = pair.BT.Decimals
	} else {
		decimals = pair.QT.Decimals
	}

	return numbers.BigIntToFloatString(vol, decimals), utils.TxFee(tx.GasPrice(), r.GasUsed), nil
}
