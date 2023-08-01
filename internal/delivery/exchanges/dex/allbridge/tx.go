package allbridge

import (
	"exchange-provider/internal/delivery/exchanges/dex/allbridge/types"
	"exchange-provider/internal/delivery/exchanges/dex/evm/contracts"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/bind"
	"exchange-provider/pkg/errors"
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func errSTF() error {
	return errors.New("execution reverted: TransferHelper:safeTransferFrom")
}

func (ex *exchange) SetTxId(eo entity.Order, txId string) error {
	o := eo.(*types.Order)
	if o.Steps[0].SrcTxId == "" {
		o.Steps[0].SrcTxId = txId
	} else if len(o.Steps) == 2 && o.Steps[1].SrcTxId == "" {
		if strings.EqualFold(txId, o.Steps[0].SrcTxId) {
			return errors.Wrap(errors.ErrForbidden,
				errors.NewMesssage("this txId has used for previous step"))
		}
		o.Steps[1].SrcTxId = txId
	} else {
		return errors.Wrap(errors.ErrForbidden,
			errors.NewMesssage("txId for this order has setted before"))
	}
	return ex.repo.Update(o)
}

func (ex *exchange) CreateTx(eo entity.Order, step int) (entity.Tx, error) {
	o := eo.(*types.Order)
	si := step - 1
	if si == 0 {
		return ex.createTx(o, si)
	} else if si == 1 {
		if len(o.Steps) == 2 {
			if err := ex.UpdateStatus(o); err != nil {
				return nil, err
			}

			return ex.createTx(o, 1)

		} else {
			return nil, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("step is out of range"))
		}
	} else {
		return nil, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("step is out of range"))
	}
}

func (ex *exchange) createTx(o *types.Order, stepIndex int) (entity.Tx, error) {
	s := o.Steps[stepIndex]
	bs := [][]byte{}
	var (
		in0                 *entity.Token
		bridgeFee           *big.Int
		amountIn, feeAmount float64
	)

	if stepIndex == 0 {
		amountIn = s.Routes[0].AmountIn
		feeAmount = o.FeeAmount + o.ExchangeFeeAmount
	} else if stepIndex == 1 {
		amountIn = s.AmountIn
	}

	if s.GeneratedTime == 0 && s.TxData == nil && s.TxValue == nil {
		for routeIndex := 0; routeIndex < len(s.Routes); routeIndex++ {
			ri := s.Routes[routeIndex]
			ni := ex.ns[ri.In.Network]

			mainContract := common.HexToAddress(ex.cfg.Networks.network(ri.In.Network).MainContract)
			exId, _ := strconv.Atoi(strings.Split(ri.ExchangeNid, "-")[1])

			p, err := ex.pairs.Get(uint(exId), ri.In.String(), ri.Out.String())
			if err != nil {
				return nil, err
			}

			var in, out *entity.Token
			if p.T1.String() == ri.In.String() {
				in = p.T1
				out = p.T2
			} else {
				in = p.T2
				out = p.T1
			}

			if routeIndex == 0 {
				in0 = in.Snapshot()
			}
			if routeIndex == 0 {
				if ri.NeedApprove && !ri.Approved {
					need, err := ni.NeedApproval(in, o.Sender, amountIn)
					if err != nil {
						return nil, err
					}
					if need {
						return ni.ApproveTx(in, o.Sender, stepIndex+1)
					}
					ri.Approved = true
					ex.repo.Update(o)
				}
			}
			if ex.isInternal(ri) {
				var afterSwap bool
				if len(s.Routes) > 1 && routeIndex > 0 && !ex.isInternal(s.Routes[routeIndex-1]) {
					afterSwap = true
				}
				input, err := ex.getBridgeInput(o, ri, ni, in, out, afterSwap, amountIn, feeAmount)
				if err != nil {
					return nil, err
				}

				bridgeFee = input.BridgeFee
				b, err := ni.BridgeData(input)
				if err != nil {
					return nil, err
				}
				bs = append(bs, b)
			} else {
				exi, err := ex.exs.GetByNID(ri.ExchangeNid)
				if err != nil {
					return nil, err
				}

				var tokenOwner, sender, receiver common.Address
				if stepIndex == 0 {
					tokenOwner = common.HexToAddress(o.Sender)
					sender = common.HexToAddress(o.Sender)
					receiver = mainContract
				} else {
					tokenOwner = common.HexToAddress(ex.cfg.Networks.network(ri.Out.Network).MainContract)
					receiver = common.HexToAddress(o.Receiver)
					sender = receiver
				}

				prvKey := ex.cfg.Networks.network(ri.In.Network).prvKey
				b, err := exi.(entity.EVMDex).CreateSwapBytes(ri.In, ri.Out, tokenOwner, sender,
					receiver, mainContract, amountIn, feeAmount, prvKey)
				if err != nil {
					return nil, err
				}
				bs = append(bs, b)
			}

		}
	} else {
		bs = s.TxData
	}

	mainContract := common.HexToAddress(ex.cfg.Networks.network(s.Routes[0].In.Network).MainContract)
	prvKey := ex.cfg.Networks.network(s.Routes[0].In.Network).prvKey
	c, err := contracts.NewContracts(mainContract, ex.cfg.Networks.network(s.Routes[0].In.Network).client)
	if err != nil {
		return nil, err
	}

	opts, err := bind.NewKeyedTransactorWithChainID(prvKey,
		big.NewInt(ex.cfg.Networks.network(s.Routes[0].In.Network).chainId))
	if err != nil {
		return nil, err
	}

	if s.TxValue == nil {
		if in0.Native {
			opts.Value, _ = new(big.Float).Mul(big.NewFloat(s.Routes[0].AmountIn),
				big.NewFloat(math.Pow10(int(in0.Decimals)))).Int(nil)
		}

		if bridgeFee != nil {
			if opts.Value == nil {
				opts.Value = bridgeFee
			} else {
				opts.Value = new(big.Int).Add(opts.Value, bridgeFee)
			}
		}
	} else {
		opts.Value = new(big.Int).SetBytes(s.TxValue)
	}

	var sender string
	if stepIndex == 0 {
		sender = o.Sender
	} else {
		sender = o.Receiver
	}

	opts.NoSend = true
	opts.From = common.HexToAddress(sender)
	opts.Sign = false

	t, err := c.Multicall(opts, bs)
	if err != nil {
		if strings.Contains(err.Error(), "this txData already processed") {
			return nil, errors.Wrap(errors.ErrForbidden,
				errors.NewMesssage("you have already completed this transaction successfully"))
		}
		if err.Error() == errSTF().Error() {
			return nil, errors.Wrap(errors.ErrBadRequest,
				errors.NewMesssage("insufficient funds or the previous step was unsuccessful"))
		}
		if strings.Contains(err.Error(), "insufficient funds") {
			return nil, errors.Wrap(errors.ErrBadRequest,
				errors.NewMesssage(err.Error()))
		}
		return nil, err
	}
	s.TxData = bs
	var val []byte
	if opts.Value == nil {
		val = common.Big0.Bytes()
	} else {
		val = opts.Value.Bytes()
	}
	s.TxValue = val
	s.AmountIn = amountIn
	if s.GeneratedTime == 0 {
		s.GeneratedTime = time.Now().UnixMilli()
	}
	if err := ex.repo.Update(o); err != nil {
		return nil, err
	}
	if opts.Value == nil {
		opts.Value = common.Big0
	}
	d := &entity.Developer{
		Function: "multicall(bytes[] calldata data) external payable returns (bytes[] memory results);",
		Contract: mainContract.Hex(),
		Value:    opts.Value.String(),
	}

	for _, b := range bs {
		d.Parameters = append(d.Parameters, hexutil.Encode(b))
	}
	if err := ex.repo.Update(o); err != nil {
		return nil, err
	}
	return &entity.EvmTx{Tx: t, CurrentStep: uint(stepIndex + 1), Sender: sender, Developer: d}, nil
}
