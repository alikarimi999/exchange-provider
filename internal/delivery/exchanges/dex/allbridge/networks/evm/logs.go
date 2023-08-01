package evm

import (
	"context"
	"exchange-provider/internal/delivery/exchanges/dex/allbridge/types"
	"exchange-provider/pkg/bind"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

var (
	contractsMetaData = &bind.MetaData{
		ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"messenger\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"}],\"name\":\"TokensReceived\",\"type\":\"event\"}]",
	}
	contractABI, _ = contractsMetaData.GetAbi()
)

type logData struct {
	Amount    *big.Int
	Recipient common.Hash
	Nonce     *big.Int
	Messenger uint8
	Message   [32]uint8
}

func (n *net) DownloadLogs(fromBlock, toBlock uint64) ([]*types.TokensReceivedLog, uint64, error) {
	var tb *big.Int
	if toBlock > 0 {
		tb = big.NewInt(int64(toBlock))
	}
	f := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(fromBlock)),
		ToBlock:   tb,
		Addresses: []common.Address{n.allbridgeContract},
		Topics:    [][]common.Hash{{common.HexToHash("0xe9d840d27ab4032a839c20760fb995af8e3ad1980b9428980ca1c7e072acd87a")}},
	}
	ls, err := n.provider.FilterLogs(context.Background(), f)
	if err != nil {
		return nil, 0, err
	}

	var lb uint64
	td := time.Now()
	trls := []*types.TokensReceivedLog{}
	for _, l := range ls {
		if l.Removed {
			continue
		}
		if l.BlockNumber > lb {
			lb = l.BlockNumber
		}
		ld := &logData{}
		if err := contractABI.UnpackIntoInterface(ld, "TokensReceived", l.Data); err != nil {
			continue
		}
		trls = append(trls, &types.TokensReceivedLog{
			TxId:       l.TxHash.Hex(),
			Recipient:  common.BytesToAddress(ld.Recipient.Bytes()).Hex(),
			Amount:     ld.Amount,
			Nonce:      ld.Nonce,
			DownloadAt: td,
		})
	}
	return trls, lb, nil
}
