package utils

import (
	"exchange-provider/internal/delivery/exchanges/dex/uniswap/v3/contracts"

	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/logger"
	"exchange-provider/pkg/wallet/eth"
	"fmt"
	"math/big"
	"sync"
	"time"

	ts "exchange-provider/internal/delivery/exchanges/dex/types"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var max96 = new(big.Int).Sub(new(big.Int).Lsh(common.Big1, 96), common.Big1)

type ApproveManager struct {
	chainId int64
	tt      *TxTracker
	w       *eth.HDWallet

	pending  *amQueue
	approved *amQueue

	ps []*ts.EthProvider
	l  logger.Logger
}

func NewApproveManager(chainId int64, tt *TxTracker, w *eth.HDWallet,
	l logger.Logger, ps []*ts.EthProvider) *ApproveManager {
	return &ApproveManager{

		tt: tt,
		w:  w,

		pending:  newAMQueue(),
		approved: newAMQueue(),

		ps: ps,
		l:  l,
	}
}

func (m *ApproveManager) exists(t ts.Token, owner, spender common.Address) bool {
	if m.pending.exists(t, owner, spender, m.chainId) || m.approved.exists(t, owner, spender, m.chainId) {
		return true
	}
	return false
}

func (am *ApproveManager) notifyApproved(t ts.Token, owner, spender common.Address) bool {
	if am.approved.exists(t, owner, spender, am.chainId) {
		return true
	}

	if am.pending.exists(t, owner, spender, am.chainId) {
		time.Sleep(time.Second * 1)
	} else {
		return false
	}

	return am.notifyApproved(t, owner, spender)
}

func (m *ApproveManager) add(t ts.Token, owner, spender common.Address, approved bool) {
	if approved {
		m.approved.add(t, owner, spender, m.chainId)
		return
	}
	m.pending.add(t, owner, spender, m.chainId)
}

func (m *ApproveManager) remove(t ts.Token, owner, spender common.Address, fromPending bool) {
	if fromPending {
		m.pending.remove(t, owner, spender, m.chainId)
		return
	}
	m.approved.remove(t, owner, spender, m.chainId)
}

func (am *ApproveManager) InfinitApproves(t ts.Token, spender common.Address,
	owners ...common.Address) []error {
	// agent := am.u.agent("approveManager.infinitApproves")

	var errs []error

	wg := &sync.WaitGroup{}
	for _, owner := range owners {
		wg.Add(1)
		go func(o common.Address) {
			defer wg.Done()
			if am.exists(t, o, spender) {
				if !am.notifyApproved(t, o, spender) {
					errs = append(errs, fmt.Errorf("%s-%s didn't receive approval", t.Symbol, o))
				}
				return
			}

			am.add(t, o, spender, false)
			err := am.infinitApprove(t, o, spender)
			if err == nil {
				am.add(t, o, spender, true)
			}
			am.remove(t, o, spender, true)

			if err != nil {
				errs = append(errs, err)
			}
		}(owner)
	}
	wg.Wait()
	return errs
}

func (am *ApproveManager) infinitApprove(t ts.Token, owner, spender common.Address) error {
	agent := "approveManager.infinitApprove"

	prefix := fmt.Sprintf("%s-%s", t.Symbol, owner)

	amount, err := am.allowance(t, owner, spender)
	if err != nil {
		return err
	}

	if amount.Cmp(max96) == -1 {
		tx, err := am.approve(t, owner, spender, abi.MaxUint256)
		if err != nil {
			return errors.Wrap(errors.Op(agent), err)
		}
		am.l.Debug(agent, fmt.Sprintf("`%s` approving tx:`%s`", prefix, tx.Hash().String()))
		doneCh := make(chan struct{})
		tf := &TtFeed{
			P:        am.provider(),
			TxHash:   tx.Hash(),
			Receiver: &t.Address,
			NeedTx:   false,
			DoneCh:   doneCh,
		}

		go am.tt.Track(tf)
		<-doneCh
		if tf.Status == TxSuccess {
			am.l.Debug(agent, fmt.Sprintf("`%s` approve done", prefix))
			return nil
		}

		if tf.Faildesc != "" {
			return errors.Wrap(errors.Op(agent), fmt.Sprintf("`%s` tx `%s` failed `%s`", prefix, tx.Hash().String(), tf.Faildesc))
		}
		return errors.Wrap(errors.Op(agent), fmt.Sprintf("`%s` approve tx failed `%s`", prefix, tx.Hash().String()))
	}

	am.l.Debug(agent, fmt.Sprintf("`%s` approved before", prefix))
	return nil
}

func (am *ApproveManager) allowance(token ts.Token, owner, spender common.Address) (*big.Int, error) {
	p := am.provider()
	c, err := contracts.NewMain(token.Address, p.Client)
	if err != nil {
		return nil, err
	}
	return c.Allowance(nil, owner, spender)
}

func (am *ApproveManager) approve(token ts.Token, owner, spender common.Address,
	amount *big.Int) (*types.Transaction, error) {
	p := am.provider()

	var err error
	c, err := contracts.NewMain(token.Address, p.Client)
	if err != nil {
		return nil, err
	}

	opts, err := am.w.NewKeyedTransactorWithChainID(owner, common.Big0, am.chainId)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			am.w.ReleaseNonce(owner, opts.Nonce.Uint64())
		} else {
			am.w.BurnNonce(owner, opts.Nonce.Uint64())

		}
	}()

	tx, err := c.Approve(opts, spender, amount)
	if err != nil {
		return nil, err
	}
	return tx, nil
}
