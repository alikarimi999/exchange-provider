package panckakeswapv2

import (
	"exchange-provider/internal/entity"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (p *Panckakeswapv2) parseSwapLogs(o *entity.Order,
	tx *types.Transaction, r *types.Receipt) (*big.Int, error) {

	if r.Logs[len(r.Logs)-1].Topics[0] == p.abi.Events["Withdrawal"].ID {

		a, err := p.abi.Events["Withdrawal"].Inputs.Unpack(r.Logs[len(r.Logs)-1].Data)
		if err != nil {
			return nil, err
		}
		return a[0].(*big.Int), nil
	} else if r.Logs[len(r.Logs)-1].Topics[0] == p.abi.Events["Swap"].ID {
		amounts, err := p.abi.Events["Swap"].Inputs.Unpack(r.Logs[len(r.Logs)-1].Data)
		if err != nil {
			return nil, err
		}
		if amounts[2].(*big.Int).Cmp(common.Big0) == 1 {
			return amounts[2].(*big.Int), nil
		} else {
			return amounts[3].(*big.Int), nil
		}
	} else {
		return nil, fmt.Errorf("unable to parse tx logs")
	}
}
