package dex

import (
	"exchange-provider/internal/delivery/exchanges/dex/uniswap/v3/contracts"
	"exchange-provider/pkg/errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	ts "exchange-provider/internal/delivery/exchanges/dex/types"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type approveManager struct {
	u *dex

	pending  *amQueue
	approved *amQueue
}

func newApproveManager(u *dex) *approveManager {
	return &approveManager{
		u:        u,
		pending:  newAMQueue(),
		approved: newAMQueue(),
	}
}

func (m *approveManager) exists(t ts.Token, owner, spender common.Address) bool {
	if m.pending.exists(t, owner, spender) || m.approved.exists(t, owner, spender) {
		return true
	}
	return false
}

func (am *approveManager) notifyApproved(t ts.Token, owner, spender common.Address) bool {
	if am.approved.exists(t, owner, spender) {
		return true
	}

	if am.pending.exists(t, owner, spender) {
		time.Sleep(time.Second * 1)
	} else {
		return false
	}

	return am.notifyApproved(t, owner, spender)
}

func (m *approveManager) add(t ts.Token, owner, spender common.Address, approved bool) {
	if approved {
		m.approved.add(t, owner, spender)
		return
	}
	m.pending.add(t, owner, spender)
}

func (m *approveManager) remove(t ts.Token, owner, spender common.Address, fromPending bool) {
	if fromPending {
		m.pending.remove(t, owner, spender)
		return
	}
	m.approved.remove(t, owner, spender)
}

func (am *approveManager) infinitApproves(t ts.Token, spender common.Address, owners ...common.Address) []error {
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

func (am *approveManager) infinitApprove(t ts.Token, owner, spender common.Address) error {
	agent := am.u.agent("approveManager.infinitApprove")
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
		am.u.l.Debug(agent, fmt.Sprintf("`%s` approving tx:`%s`", prefix, tx.Hash().String()))
		doneCh := make(chan struct{})
		tf := &ttFeed{
			txHash:   tx.Hash(),
			receiver: &t.Address,
			needTx:   false,
			doneCh:   doneCh,
		}

		go am.u.tt.track(tf)
		<-doneCh
		if tf.status == txSuccess {
			am.u.l.Debug(agent, fmt.Sprintf("`%s` approve done", prefix))
			return nil
		}

		if tf.faildesc != "" {
			return errors.Wrap(errors.Op(agent), fmt.Sprintf("`%s` tx `%s` failed `%s`", prefix, tx.Hash().String(), tf.faildesc))
		}
		return errors.Wrap(errors.Op(agent), fmt.Sprintf("`%s` approve tx failed `%s`", prefix, tx.Hash().String()))
	}

	am.u.l.Debug(agent, fmt.Sprintf("`%s` approved before, amount: `%s`", prefix, amount))
	return nil
}

func (am *approveManager) allowance(token ts.Token, owner, spender common.Address) (*big.Int, error) {
	p := am.u.provider()
	c, err := contracts.NewMain(token.Address, p.Client)
	if err != nil {
		return nil, err
	}

	return c.Allowance(nil, owner, spender)
}

func (am *approveManager) approve(token ts.Token, owner, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	p := am.u.provider()

	var err error
	c, err := contracts.NewMain(token.Address, p.Client)
	if err != nil {
		return nil, err
	}

	opts, err := am.u.wallet.NewKeyedTransactorWithChainID(owner, common.Big0, int64(am.u.cfg.ChianId))
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			am.u.wallet.ReleaseNonce(owner, opts.Nonce.Uint64())
		} else {
			am.u.wallet.BurnNonce(owner, opts.Nonce.Uint64())

		}
	}()

	tx, err := c.Approve(opts, spender, amount)
	if err != nil {
		return nil, err
	}
	return tx, nil
}
