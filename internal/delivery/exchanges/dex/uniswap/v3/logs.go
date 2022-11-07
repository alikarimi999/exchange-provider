package uniswapv3

import (
	ts "exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/internal/delivery/exchanges/dex/utils"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/utils/numbers"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (u *UniswapV3) ParseSwapLogs(o *entity.UserOrder, tx *types.Transaction, pair *ts.Pair, receipt *types.Receipt) (amount, fee string, err error) {
	for _, log := range receipt.Logs {
		if len(log.Topics) == 3 && log.Topics[0] == erc20TransferSignature &&
			utils.HashToAddress(log.Topics[2]) == common.HexToAddress(o.Deposit.Addr) {

			var decimals int
			if o.Side == entity.SideBuy {
				decimals = pair.BT.Decimals
			} else {
				decimals = pair.QT.Decimals
			}

			return numbers.BigIntToFloatString(new(big.Int).SetBytes(log.Data), decimals), utils.TxFee(tx.GasPrice(), receipt.GasUsed), nil

		}
	}

	return "", "", errors.New("unable to parse tx logs")
}
