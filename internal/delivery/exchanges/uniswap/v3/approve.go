package uniswapv3

import (
	"fmt"
	"math/big"
	"exchange-provider/internal/delivery/exchanges/uniswap/v3/contracts"
	"exchange-provider/pkg/errors"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type approveManager struct {
	u *UniSwapV3

	pending  *amQueue
	approved *amQueue
}

func newApproveManager(u *UniSwapV3) *approveManager {
	return &approveManager{
		u:        u,
		pending:  newAMQueue(),
		approved: newAMQueue(),
	}
}

func (m *approveManager) exists(t token, owner, spender common.Address) bool {
	if m.pending.exists(t, owner, spender) || m.approved.exists(t, owner, spender) {
		return true
	}
	return false
}

func (am *approveManager) notifyApproved(t token, owner, spender common.Address) bool {
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

func (m *approveManager) add(t token, owner, spender common.Address, approved bool) {
	if approved {
		m.approved.add(t, owner, spender)
		return
	}
	m.pending.add(t, owner, spender)
}

func (m *approveManager) remove(t token, owner, spender common.Address, fromPending bool) {
	if fromPending {
		m.pending.remove(t, owner, spender)
		return
	}
	m.approved.remove(t, owner, spender)
}

func (am *approveManager) infinitApproves(t token, spender common.Address, owners ...common.Address) []error {
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

func (am *approveManager) infinitApprove(t token, owner, spender common.Address) error {
	agent := am.u.agent("approveManager.infinitApprove")
	prefix := fmt.Sprintf("%s-%s", t.Symbol, owner)

	amount, err := am.allowance(t, owner, spender)
	if err != nil {
		return err
	}

	max := abi.MaxUint256
	if t.Symbol == "UNI" {
		max = max96
	}

	if amount.Cmp(max) == -1 {
		tx, err := am.approve(t, owner, spender, abi.MaxUint256)
		if err != nil {
			return errors.Wrap(errors.Op(agent), err)
		}
		am.u.l.Debug(agent, fmt.Sprintf("`%s` approving", prefix))
		doneCh := make(chan struct{})
		tf := &ttFeed{
			txHash:   tx.Hash(),
			receiver: &t.Address,
			needTx:   false,
			doneCh:   doneCh,
		}

		am.u.tt.push(tf)
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

func (am *approveManager) allowance(token token, owner, spender common.Address) (*big.Int, error) {
	c, err := contracts.NewMain(token.Address, am.u.provider.Client)
	if err != nil {
		return nil, err
	}

	return c.Allowance(nil, owner, spender)
}

func (am *approveManager) approve(token token, owner, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	c, err := contracts.NewMain(token.Address, am.u.provider.Client)
	if err != nil {
		return nil, err
	}

	opts, err := am.u.newKeyedTransactorWithChainID(owner, common.Big0)
	if err != nil {
		return nil, err
	}

	tx, err := c.Approve(opts, spender, amount)
	if err != nil {
		am.u.wallet.ReleaseNonce(owner, opts.Nonce.Uint64())
		return nil, err
	}
	am.u.wallet.BurnNonce(owner, tx.Nonce())
	return tx, nil
}
