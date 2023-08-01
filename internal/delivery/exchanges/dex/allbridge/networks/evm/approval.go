package evm

import (
	"exchange-provider/internal/delivery/exchanges/dex/evm/contracts/erc20"
	"exchange-provider/internal/entity"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"math/big"

	em "github.com/ethereum/go-ethereum/common/math"

	"github.com/ethereum/go-ethereum/common"
)

func (n *net) NeedApproval(in *entity.Token, owner string, minAmount float64) (bool, error) {
	return n.needApproval(in, common.HexToAddress(owner), minAmount)
}

func (n *net) needApproval(in *entity.Token, owner common.Address, minAmount float64) (bool, error) {
	if in.Native {
		return false, nil
	}

	c, err := erc20.NewContracts(common.HexToAddress(in.ContractAddress), n.provider)
	if err != nil {
		return false, err
	}
	amount, err := c.Allowance(nil, owner, n.mainContract)
	if err != nil {
		return false, err
	}

	mAmount, _ := big.NewFloat(0).Mul(big.NewFloat(minAmount),
		big.NewFloat(0).SetInt(em.BigPow(10, int64(in.Decimals)))).Int(nil)
	return amount.Cmp(mAmount) == -1, nil
}

func (n *net) ApproveTx(in *entity.Token, ownder string, step int) (entity.Tx, error) {

	c, err := erc20.NewContracts(common.HexToAddress(in.ContractAddress), n.provider)
	if err != nil {
		return nil, err
	}

	opts, err := bind.NewKeyedTransactorWithChainID(n.prvKey, big.NewInt(n.chainId))
	if err != nil {
		return nil, err
	}
	opts.NoSend = true

	tx, err := c.Approve(opts, n.mainContract, em.MaxBig256)
	if err != nil {
		return nil, err
	}

	d := &entity.Developer{
		Function:   "approve(address spender, uint256 value) external returns (bool);",
		Contract:   in.ContractAddress,
		Parameters: []string{n.mainContract.Hex(), em.MaxBig256.String()},
		Value:      common.Big0.String(),
	}

	return &entity.EvmTx{
		Tx:          tx,
		IsApproveTx: true,
		Sender:      ownder,
		CurrentStep: uint(step),
		Developer:   d,
	}, nil
}
