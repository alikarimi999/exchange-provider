package uniswapv3

import (
	"context"
	"errors"
	"exchange-provider/pkg/logger"
	"fmt"
	"time"

	"exchange-provider/pkg/try"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	txFailed = iota
	txSuccess
	// txNotFound
)

var errTxNotFound = "not found"

type ttFeed struct {
	txHash   common.Hash
	receiver *common.Address
	needTx   bool

	maxRetries uint64

	confirms  uint64
	confirmed uint64

	status   int
	faildesc string
	tx       *types.Transaction
	*types.Receipt

	doneCh chan<- struct{}
}

type txTracker struct {
	us *dex
	l  logger.Logger

	ctx context.Context
}

func newTxTracker(us *dex) *txTracker {
	return &txTracker{
		us: us,
		l:  us.l,

		ctx: context.Background(),
	}
}

func (tr *txTracker) track(f *ttFeed) {
	agent := tr.us.agent("txTracker.track")

	p := tr.us.provider()

	if f.maxRetries == 0 {
		f.maxRetries = 100
	}
	if f.confirms == 0 {
		f.confirms = tr.us.confirms
	}

	err := try.Do(f.maxRetries, func(attempt uint64) (bool, error) {
		if f.needTx {
			tx, pending, err := p.TransactionByHash(tr.ctx, f.txHash)
			if err != nil {
				if err.Error() == errTxNotFound {
					if attempt < f.maxRetries {
						time.Sleep(tr.us.blockTime)
						return true, err
					}
					f.status = txFailed
					f.faildesc = "tx not found"
					f.doneCh <- struct{}{}
					return false, nil
				}
				tr.l.Error(agent, err.Error())
				time.Sleep(tr.us.blockTime)
				return true, err
			}
			if pending {
				time.Sleep(tr.us.blockTime)
				return true, errors.New("pending")
			}

			if *tx.To() != *f.receiver {

				f.status = txFailed
				f.faildesc = fmt.Sprintf("invalid destination address `%s`", tx.To())
				f.doneCh <- struct{}{}
				return false, nil
			}
			f.tx = tx
		}

		receipt, err := p.TransactionReceipt(tr.ctx, f.txHash)
		if err != nil {
			if err.Error() == errTxNotFound {
				if attempt < f.maxRetries {
					time.Sleep(tr.us.blockTime * time.Duration(f.confirms))
					return true, err
				}
				f.status = txFailed
				f.faildesc = "tx not found"
				f.doneCh <- struct{}{}
				return false, nil
			}
			tr.l.Error(agent, err.Error())
			time.Sleep(tr.us.blockTime)
			return true, err
		}

		f.Receipt = receipt
		switch f.Receipt.Status {
		case txSuccess:
			bn := receipt.BlockNumber.Uint64()
			cn, err := p.BlockNumber(tr.ctx)
			if err != nil {
				tr.l.Error(agent, err.Error())
				return true, err
			}

			confirmed := cn - bn
			if confirmed >= f.confirms {
				f.confirmed = confirmed
				f.status = txSuccess
				f.doneCh <- struct{}{}
				return false, nil
			}

			t := f.confirms - confirmed
			time.Sleep(tr.us.blockTime * time.Duration(t))
			return true, fmt.Errorf("confirmed `%d` blocks", confirmed)
		default:
			f.status = txFailed
			f.doneCh <- struct{}{}
			return false, nil
		}
	})
	if err != nil {
		f.status = txFailed
		f.faildesc = err.Error()
		f.doneCh <- struct{}{}
	}

}
