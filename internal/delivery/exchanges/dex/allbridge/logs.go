package allbridge

import (
	"context"
	"exchange-provider/internal/entity"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var errInvalidTx = fmt.Errorf("tx is invalid")

func (ex *exchange) getTokenTransferAmount(txId string, t *entity.Token,
	receiver common.Address) (value *big.Int,
	isPending bool, err error) {
	txHash := common.HexToHash(txId)
	cl := ex.cfg.Networks.network(t.Id.Network).client
	tx, pending, err := cl.TransactionByHash(context.Background(), txHash)
	if err != nil {
		return nil, false, err
	}
	if pending {
		return nil, pending, nil
	}
	receipt, err := cl.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		log.Fatal(err)
	}

	e := ex.erc20.Events["Transfer"]
	for _, l := range receipt.Logs {
		if len(l.Topics) == 3 && l.Address.Big().Cmp(common.HexToAddress(t.ContractAddress).Big()) == 0 &&
			e.ID.Big().Cmp(l.Topics[0].Big()) == 0 &&
			common.HexToAddress(l.Topics[2].Hex()).Big().Cmp(receiver.Big()) == 0 {
			i, err := ex.erc20.Unpack(e.Name, l.Data)
			if err != nil {
				fmt.Println(err)
			}
			v, ok := i[0].(*big.Int)
			if ok {
				return v, false, nil
			}
		}
	}

	return nil, false, errInvalidTx
}

func (ex *exchange) getNativeTransferAmount(txId string, t *entity.Token) (value *big.Int,
	isPending bool, err error) {

	txHash := common.HexToHash(txId)
	cl := ex.cfg.Networks.network(t.Id.Network).client
	tx, pending, err := cl.TransactionByHash(context.Background(), txHash)
	if err != nil {
		return nil, false, err
	}
	if pending {
		return nil, pending, nil
	}

	receipt, err := cl.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		return nil, false, err
	}

	e := ex.erc20.Events["Withdrawal"]
	for _, l := range receipt.Logs {
		if len(l.Topics) == 2 && l.Address.Big().Cmp(common.HexToAddress(t.ContractAddress).Big()) == 0 &&
			e.ID.Big().Cmp(l.Topics[0].Big()) == 0 {
			i, err := ex.erc20.Unpack(e.Name, l.Data)
			if err != nil {
				return nil, false, err
			}

			v, ok := i[0].(*big.Int)
			if ok {
				return v, false, nil
			}
		}
	}

	return nil, false, errInvalidTx
}
