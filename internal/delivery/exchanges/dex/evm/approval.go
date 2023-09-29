package evm

import (
	"exchange-provider/internal/delivery/exchanges/dex/evm/contracts/erc20"
	"exchange-provider/internal/entity"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
)

func (d *exchange) GetToken(tId entity.TokenId) *entity.Token {
	return d.tl.get(tId)
}

func (d *exchange) Allowance(t *entity.Token, owner string) (*entity.Allowance, error) {
	a := &entity.Allowance{
		Token:   t,
		Owner:   owner,
		Spender: d.cfg.Contract,
	}
	if t.Native {
		a.Amount = math.MaxBig256
		return a, nil
	}

	c, err := erc20.NewContracts(common.HexToAddress(t.ContractAddress), d.provider())
	if err != nil {
		return nil, err
	}

	amount, err := c.Allowance(nil, common.HexToAddress(owner), d.cfg.contractAddress)
	if err != nil {
		return nil, err
	}
	a.Amount = amount
	return a, nil
}

func (d *exchange) approveTx(in *entity.Token) ([]byte, error) {
	return d.erc20.Pack("approve", d.cfg.contractAddress, math.MaxBig256)
}

func (d *exchange) needApproval(in *entity.Token, owner common.Address, minAmount float64) (bool, error) {
	if in.Native {
		return false, nil
	}

	c, err := erc20.NewContracts(common.HexToAddress(in.ContractAddress), d.provider())
	if err != nil {
		return false, err
	}

	amount, err := c.Allowance(nil, owner, d.cfg.contractAddress)
	if err != nil {
		return false, err
	}

	mAmount, _ := big.NewFloat(0).Mul(big.NewFloat(minAmount),
		big.NewFloat(0).SetInt(math.BigPow(10, int64(in.Decimals)))).Int(nil)
	return amount.Cmp(mAmount) == -1, nil
}
