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

func (d *exchange) CreateTx(ord entity.Order) ([]entity.Tx, error) {
	o := ord.(*types.Order)
	p, err := d.pairs.Get(d.Id(), o.In.String(), o.Out.String())
	if err != nil {
		return nil, err
	}

	txs := []entity.Tx{}

	var in, out *entity.Token
	if p.T1.Id.String() == o.In.String() {
		in = p.T1
		out = p.T2
	} else {
		in = p.T2
		out = p.T1
	}

	toTx := d.cfg.contractAddress.String()
	if o.NeedApprove {
		etx := &entity.EvmTx{CurrentStep: 1, From: o.Sender.Hex(), To: toTx, Value: common.Big0}
		txData, err := d.approveTx(in)
		if err != nil {
			return nil, err
		}

		etx.IsApproveTx = true
		etx.Network = in.Id.Network
		etx.TxData = txData
		d := &entity.Developer{
			Function:   "approve(address spender, uint256 value) external returns (bool);",
			Contract:   in.ContractAddress,
			Parameters: []string{d.cfg.contractAddress.Hex(), em.MaxBig256.String()},
			Value:      etx.Value.String(),
		}
		etx.Developer = d
		txs = append(txs, etx)
	}

	txData, bs, val, err := d.createTx(in, out, o.Sender, o.Sender, o.Receiver,
		d.cfg.contractAddress, o.AmountIn, o.FeeAmount+o.ExchangeFeeAmount, d.cfg.prvKey)
	if err != nil {
		return nil, err
	}

	etx := &entity.EvmTx{Network: in.Id.Network, CurrentStep: 1, From: o.Sender.Hex(), To: toTx, Value: val}
	etx.TxData = txData
	dev := &entity.Developer{
		Function: "multicall(bytes[] calldata data) external payable returns (bytes[] memory results);",
		Contract: d.cfg.contractAddress.Hex(),
		Value:    val.String(),
	}

	for _, b := range bs {
		dev.Parameters = append(dev.Parameters, hexutil.Encode(b))
	}
	etx.Developer = dev
	txs = append(txs, etx)
	return txs, nil
}
