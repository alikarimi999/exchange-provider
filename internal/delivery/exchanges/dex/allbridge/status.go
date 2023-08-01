package allbridge

import (
	"exchange-provider/internal/delivery/exchanges/dex/allbridge/types"
	"exchange-provider/internal/entity"
	"math"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

func (ex *exchange) UpdateStatus(eo entity.Order) error {
	var changed bool
	o := eo.(*types.Order)
	for i, s := range o.Steps {
		if s.SrcTxId == "" {
			continue
		} else {
			if i == 0 {
				if s.DstTxId == "" && (s.Status == entity.OCreated.String() ||
					s.Status == entity.OPending.String()) {
					lr := s.Routes[len(s.Routes)-1]
					out, err := ex.tl.getTokenInfo(lr.Out)
					if err != nil {
						return err
					}

					l, ok := ex.c.getRecievedLog(lr.Out.Network, o.Nonce)
					if !ok {
						return nil
					}
					s.AmountOut, _ = new(big.Float).Quo(new(big.Float).SetInt(l.Amount),
						big.NewFloat(math.Pow10(out.Decimals))).Float64()
					if len(o.Steps) == 2 {
						o.Steps[1].AmountIn = s.AmountOut
					}
					s.DstTxId = l.TxId
					s.Status = entity.OCompleted.String()
					changed = true
				}
			} else {
				if s.SrcTxId != "" && (s.Status == entity.OCreated.String() ||
					s.Status == entity.OPending.String()) {
					r := s.Routes[len(s.Routes)-1]
					ss := strings.Split(r.ExchangeNid, "-")
					id, err := strconv.Atoi(ss[len(ss)-1])
					if err != nil {
						return err
					}

					p, err := ex.pairs.Get(uint(id), r.In.String(), r.Out.String())
					if err != nil {
						return err
					}
					var out *entity.Token
					if p.T1.Id == r.Out {
						out = p.T1
					} else {
						out = p.T2
					}
					var (
						v       *big.Int
						pending bool
					)

					if out.Native {
						v, pending, err = ex.getNativeTransferAmount(s.SrcTxId, out)
					} else {
						v, pending, err = ex.getTokenTransferAmount(s.SrcTxId,
							out, common.HexToAddress(s.Receiver))
					}

					if err != nil {
						if err == errInvalidTx {
							s.Status = entity.OFailed.String()
							s.FailedDesc = err.Error()
							changed = true
							break
						}
						return err
					}

					if pending {
						continue
					}

					s.AmountOut, _ = new(big.Float).Quo(new(big.Float).SetInt(v),
						big.NewFloat(math.Pow10(int(out.Decimals)))).Float64()
					s.Status = entity.OCompleted.String()
					changed = true
				}
			}
		}
	}
	if changed {
		ex.repo.Update(o)
	}
	return nil
}
