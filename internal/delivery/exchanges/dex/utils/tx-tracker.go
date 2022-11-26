package utils

import (
	"context"
	"errors"
	"fmt"
	"time"

	ts "exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/pkg/logger"
	"exchange-provider/pkg/try"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	TxFailed = iota
	TxSuccess
	// txNotFound
)

var ErrTxNotFound = "not found"

type TtFeed struct {
	P        *ts.Provider
	TxHash   common.Hash
	Receiver *common.Address
	NeedTx   bool

	MaxRetries uint64

	Confirms  uint64
	Confirmed uint64

	Status   int
	Faildesc string
	Tx       *types.Transaction
	*types.Receipt

	DoneCh chan<- struct{}
}

type TxTracker struct {
	exId     string
	confirms uint64
	bt       time.Duration

	l   logger.Logger
	ctx context.Context
}

func NewTxTracker(id string, bt time.Duration, confirms uint64, l logger.Logger) *TxTracker {
	return &TxTracker{
		exId:     id,
		confirms: confirms,
		bt:       bt,

		l:   l,
		ctx: context.Background(),
	}
}

func (tr *TxTracker) Track(f *TtFeed) {
	agent := tr.exId + "txTracker.track"

	p := f.P

	if f.MaxRetries == 0 {
		f.MaxRetries = 100
	}
	if f.Confirms == 0 {
		f.Confirms = tr.confirms
	}

	err := try.Do(f.MaxRetries, func(attempt uint64) (bool, error) {
		if f.NeedTx {
			tx, pending, err := p.TransactionByHash(tr.ctx, f.TxHash)
			if err != nil {
				if err.Error() == ErrTxNotFound {
					if attempt < f.MaxRetries {
						time.Sleep(tr.bt)
						return true, err
					}
					f.Status = TxFailed
					f.Faildesc = "tx not found"
					f.DoneCh <- struct{}{}
					return false, nil
				}
				tr.l.Error(agent, err.Error())
				time.Sleep(tr.bt)
				return true, err
			}
			if pending {
				time.Sleep(tr.bt)
				return true, errors.New("pending")
			}

			if *tx.To() != *f.Receiver {

				f.Status = TxFailed
				f.Faildesc = fmt.Sprintf("invalid destination address `%s`", tx.To())
				f.DoneCh <- struct{}{}
				return false, nil
			}
			f.Tx = tx
		}

		receipt, err := p.TransactionReceipt(tr.ctx, f.TxHash)
		if err != nil {
			if err.Error() == ErrTxNotFound {
				if attempt < f.MaxRetries {
					time.Sleep(tr.bt * time.Duration(f.Confirms))
					return true, err
				}
				f.Status = TxFailed
				f.Faildesc = "tx not found"
				f.DoneCh <- struct{}{}
				return false, nil
			}
			tr.l.Error(agent, err.Error())
			time.Sleep(tr.bt)
			return true, err
		}

		f.Receipt = receipt
		switch f.Receipt.Status {
		case TxSuccess:
			bn := receipt.BlockNumber.Uint64()
			cn, err := p.BlockNumber(tr.ctx)
			if err != nil {
				tr.l.Error(agent, err.Error())
				return true, err
			}

			confirmed := cn - bn
			if confirmed >= f.Confirms {
				f.Confirmed = confirmed
				f.Status = TxSuccess
				f.DoneCh <- struct{}{}
				return false, nil
			}

			t := f.Confirms - confirmed
			time.Sleep(tr.bt * time.Duration(t))
			return true, fmt.Errorf("confirmed `%d` blocks", confirmed)
		default:
			f.Status = TxFailed
			f.DoneCh <- struct{}{}
			return false, nil
		}
	})
	if err != nil {
		f.Status = TxFailed
		f.Faildesc = err.Error()
		f.DoneCh <- struct{}{}
	}

}
