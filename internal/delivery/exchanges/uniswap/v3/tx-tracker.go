package uniswapv3

import (
	"context"
	"errors"
	"fmt"
	"order_service/pkg/logger"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/matryer/try"
)

const (
	txFailed = iota
	txSuccess
	txNotFound
)

var errTxNotFound = errors.New("not found")

type ttFeed struct {
	txHash   common.Hash
	receiver *common.Address
	needTx   bool

	status   int
	faildesc string
	tx       *types.Transaction
	*types.Receipt

	doneCh chan<- struct{}
}

type txTracker struct {
	us *UniSwapV3
	c  *ethclient.Client
	l  logger.Logger

	maxRetries int

	fCh chan *ttFeed
	ctx context.Context
}

func newTxTracker(us *UniSwapV3) *txTracker {
	return &txTracker{
		us: us,
		c:  us.dp,
		l:  us.l,

		maxRetries: 4,
		fCh:        make(chan *ttFeed),
		ctx:        context.Background(),
	}
}

func (tr *txTracker) run(wg *sync.WaitGroup, stopCh chan struct{}) {
	agent := tr.us.agent("txTracker")
	defer wg.Done()
	try.MaxRetries = tr.maxRetries

	for {
		select {
		case feed := <-tr.fCh:
			go func(f *ttFeed) {
				err := try.Do(func(attempt int) (bool, error) {
					tr.l.Debug(agent, fmt.Sprintf("attempt: `%d`, txId: `%s`", attempt, f.txHash))

					if f.needTx {
						tx, pending, err := tr.c.TransactionByHash(tr.ctx, f.txHash)
						if err != nil {
							if err == errTxNotFound {
								if attempt == 1 {
									time.Sleep(tr.us.blockTime)
									return true, err
								}
								f.status = txNotFound
								f.doneCh <- struct{}{}
								return false, nil
							}
							tr.l.Error(agent, err.Error())
							return true, err
						}
						if pending {
							time.Sleep(tr.us.blockTime)
							return true, errors.New("pending")
						}
						if tx.To() != f.receiver {
							f.status = txFailed
							f.faildesc = fmt.Sprintf("invalid destination address `%s`", tx.To())
							f.doneCh <- struct{}{}
							return false, nil
						}
						f.tx = tx
					}

					receipt, err := tr.c.TransactionReceipt(tr.ctx, f.txHash)

					if err != nil {
						if err == errTxNotFound {
							if attempt == 1 {
								time.Sleep(tr.us.blockTime)
								return true, err
							}
							f.status = txNotFound
							f.doneCh <- struct{}{}
							return false, nil
						}
						tr.l.Error(agent, err.Error())
						return true, err
					}

					f.Receipt = receipt
					if receipt.Status == txSuccess {
						bn := receipt.BlockNumber.Uint64()
						cn, err := tr.c.BlockNumber(tr.ctx)
						if err != nil {
							tr.l.Error(agent, err.Error())
							return true, err
						}

						confirmed := cn - bn
						if confirmed >= tr.us.confirms {
							f.status = txSuccess
							f.doneCh <- struct{}{}
							return false, nil
						}

						t := tr.us.confirms - confirmed
						time.Sleep(tr.us.blockTime * time.Duration(t))
						return true, fmt.Errorf("confirmed `%d` blocks", confirmed)
					}
					f.status = txFailed
					f.doneCh <- struct{}{}
					return false, nil
				})
				if err != nil {
					f.status = txFailed
					f.faildesc = err.Error()
					f.doneCh <- struct{}{}
				}
			}(feed)

		case <-stopCh:
			tr.l.Info(agent, "stopped")
			return
		}
	}
}

func (tr *txTracker) push(feed *ttFeed) {
	tr.fCh <- feed
}
