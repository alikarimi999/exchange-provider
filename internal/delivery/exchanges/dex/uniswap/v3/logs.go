package uniswapv3

import (
	"exchange-provider/internal/delivery/exchanges/dex/utils"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (u *UniswapV3) parseSwapLogs(o *entity.CexOrder,
	tx *types.Transaction, receipt *types.Receipt) (*big.Int, error) {
	for _, log := range receipt.Logs {
		if len(log.Topics) == 3 && log.Topics[0] == erc20TransferSignature &&
			utils.HashToAddress(log.Topics[2]) == common.HexToAddress(o.Deposit.Address.Addr) {

			return new(big.Int).SetBytes(log.Data), nil

		}
	}

	return nil, errors.New("unable to parse tx logs")
}
