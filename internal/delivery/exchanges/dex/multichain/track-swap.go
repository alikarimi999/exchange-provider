package multichain

import (
	"encoding/json"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/try"
	"fmt"
	"io"
	"net/http"
	"time"
)

type info struct {
	Msg  string `json:"msg"`
	Info struct {
		Txid     string `json:"txid"`
		Swaptx   string `json:"swaptx"`
		Status   int    `json:"status"`
		Swapinfo struct {
			RouterSwapInfo struct {
				Token   string `json:"token"`
				TokenID string `json:"tokenID"`
			} `json:"routerSwapInfo"`
		} `json:"swapinfo"`
		Formatswapvalue string `json:"formatswapvalue"`
		FormatFee       string `json:"formatfee"`
	} `json:"info"`
}

func (ex *Multichain) trackSwap(o *entity.Order, index int) {
	err := try.Do(1000, func(attempt uint64) (retry bool, err error) {
		i, err := ex.getInfo(o.Swaps[index].ExId)
		if err != nil {
			time.Sleep(time.Second * 30)
			return true, err
		}
		switch i.Info.Status {
		case ExceedLimit:
			o.Swaps[index].Status = entity.ExOrderFailed
			o.Swaps[index].FailedDesc = "LessThenMinAmount"
			return false, nil

		case Confirming, Swapping:
			time.Sleep(time.Second * 30)
			return true, fmt.Errorf("Confirming")

		case Success:
			o.Swaps[index].Status = entity.ExOrderSucceed
			o.Swaps[index].OutAmount = i.Info.Formatswapvalue
			o.Swaps[index].Fee = i.Info.FormatFee
			o.Swaps[index].FeeCurrency = o.Routes[index].In.String()
			o.Swaps[index].MetaData["SwapTx"] = i.Info.Swaptx
			return false, nil
		}
		return false, fmt.Errorf("unknown error")
	})

	if err != nil {
		o.Swaps[index].Status = entity.ExOrderFailed
		o.Swaps[index].FailedDesc = err.Error()
	}

}

func (ex *Multichain) getInfo(txId string) (*info, error) {
	res, err := http.Get(ex.apiUrl + txId)
	if err != nil {
		return nil, err
	}
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	i := &info{}
	err = json.Unmarshal(b, i)
	return i, err
}
