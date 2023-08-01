package evm

import (
	"exchange-provider/internal/delivery/exchanges/dex/evm/types"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	em "github.com/ethereum/go-ethereum/common/math"
)

func (d *exchange) SetTxId(eo entity.Order, txId string) error {
	o := eo.(*types.Order)
	if o.TxId != "" {
		return errors.Wrap(errors.ErrForbidden,
			errors.NewMesssage("txId for this order has setted before"))
	}
	o.TxId = txId
	return d.repo.Update(o)
}

func (d *exchange) CreateTx(ord entity.Order) (entity.Tx, error) {
	o := ord.(*types.Order)
	p, err := d.pairs.Get(d.Id(), o.In.String(), o.Out.String())
	if err != nil {
		return nil, err
	}

	etx := &entity.EvmTx{CurrentStep: 1, Sender: o.Sender.Hex()}
	var in, out *entity.Token
	if p.T1.Id.String() == o.In.String() {
		in = p.T1
		out = p.T2
	} else {
		in = p.T2
		out = p.T1
	}

	if o.NeedApprove && !o.Approved {
		need, err := d.needApproval(in, o.Sender, o.AmountIn)
		if err != nil {
			return nil, err
		}
		if need {
			tx, err := d.approveTx(in)
			etx.IsApproveTx = true
			etx.Tx = tx
			d := &entity.Developer{
				Function:   "approve(address spender, uint256 value) external returns (bool);",
				Contract:   in.ContractAddress,
				Parameters: []string{d.cfg.contractAddress.Hex(), em.MaxBig256.String()},
				Value:      common.Big0.String(),
			}
			etx.Developer = d
			return etx, err
		} else {
			o.Approved = true
			d.repo.Update(o)
		}
	}

	tx, bs, err := d.createTx(in, out, o.Sender, o.Sender, o.Receiver,
		d.cfg.contractAddress, o.AmountIn, o.FeeAmount+o.ExchangeFeeAmount, d.cfg.prvKey)
	if err != nil {
		return nil, err
	}

	etx.IsApproveTx = false
	etx.Tx = tx

	dev := &entity.Developer{
		Function: "multicall(bytes[] calldata data) external payable returns (bytes[] memory results);",
		Contract: d.cfg.contractAddress.Hex(),
		Value:    tx.Value().String(),
	}

	for _, b := range bs {
		dev.Parameters = append(dev.Parameters, hexutil.Encode(b))
	}
	etx.Developer = dev
	return etx, nil
}
