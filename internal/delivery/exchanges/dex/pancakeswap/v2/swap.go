package panckakeswapv2

import (
	"exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/internal/delivery/exchanges/dex/utils"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/utils/numbers"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

func (ex *Panckakeswapv2) TrackSwap(o *entity.Order, p *types.Pair, i int) {
	agent := ex.id + "TrackSwap"
	doneCh := make(chan struct{})
	tf := &utils.TtFeed{
		P:        ex.provider(),
		TxHash:   common.HexToHash(o.Swaps[i].ExId),
		Receiver: &ex.router,
		NeedTx:   true,
		DoneCh:   doneCh,
	}

	go ex.tt.Track(tf)

	<-doneCh

	switch tf.Status {
	case utils.TxSuccess:
		vol, err := ex.parseSwapLogs(o, tf.Tx, tf.Receipt)
		if err != nil {
			o.Swaps[i].Status = entity.ExOrderFailed
			o.Swaps[i].FailedDesc = err.Error()
		}

		var decimals int
		if o.Routes[i].Output.CoinId == p.T1.Symbol {
			decimals = p.T1.Decimals
		} else {
			decimals = p.T2.Decimals
		}
		amount := numbers.BigIntToFloatString(vol, decimals)
		fee := utils.TxFee(tf.Tx.GasPrice(), tf.Receipt.GasUsed)

		o.Swaps[i].OutAmount = amount
		o.Swaps[i].Status = entity.ExOrderSucceed
		o.Swaps[i].Fee = fee
		o.Swaps[i].FeeCurrency = ex.nt

		ex.l.Debug(agent, fmt.Sprintf("order: `%d`, tx: `%s`, confirm: `%d/%d`",
			o.Id, tf.TxHash, tf.Confirmed, tf.Confirms))

	case utils.TxFailed:
		o.Swaps[i].Status = entity.ExOrderFailed
		o.Swaps[i].FailedDesc = tf.Faildesc
	}

}
