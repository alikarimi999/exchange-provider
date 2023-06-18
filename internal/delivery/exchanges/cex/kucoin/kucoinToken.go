package kucoin

import (
	"encoding/json"
	"exchange-provider/internal/entity"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/Kucoin/kucoin-go-sdk"
)

type kucoinToken struct {
	WithdrawDisabledTip   string `json:"withdrawDisabledTip"`
	ContractAddress       string `json:"contractAddress"`
	IsDepositEnabled      string `json:"isDepositEnabled"`
	WithdrawMinSize       string `json:"withdrawMinSize"`
	UserAddressName       string `json:"userAddressName"`
	TxURL                 string `json:"txUrl"`
	PreWithdrawTipEnabled string `json:"preWithdrawTipEnabled"`
	Currency              string `json:"currency"`
	DepositMinSize        string `json:"depositMinSize"`
	PreConfirmationCount  string `json:"preConfirmationCount"`
	PreDepositTip         string `json:"preDepositTip"`
	PreWithdrawTip        string `json:"preWithdrawTip"`
	WithdrawMinFee        string `json:"withdrawMinFee"`
	ChainName             string `json:"chainName"`
	PreDepositTipEnabled  string `json:"preDepositTipEnabled"`
	Chain                 string `json:"chain"`
	IsChainEnabled        string `json:"isChainEnabled"`
	WalletPrecision       string `json:"walletPrecision"`
	ChainFullName         string `json:"chainFullName"`
	DepositDisabledTip    string `json:"depositDisabledTip"`
	WithdrawFeeRate       string `json:"withdrawFeeRate"`
	ConfirmationCount     string `json:"confirmationCount"`
	IsWithdrawEnabled     string `json:"isWithdrawEnabled"`
	Status                string `json:"status"`
}

type tokens struct {
	Code string         `json:"code"`
	Msg  string         `json:"msg"`
	Data []*kucoinToken `json:"data"`
}

func (k *exchange) downloadTokens() (*tokens, error) {
	req, err := http.NewRequest(http.MethodGet, k.cfg.CoinListUrl, nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	ts := &tokens{}
	if err := json.Unmarshal(b, ts); err != nil {
		return nil, err
	}
	if ts.Code != "200" {
		return nil, fmt.Errorf(ts.Msg)
	}
	return ts, nil
}

func (k *exchange) isDipositAndWithdrawEnable(t *Token) (bool, bool, error) {
	bt, err := k.si.getToken(t.Currency, t.ChainName)
	if err != nil {
		return false, false, err
	}

	if bt.IsChainEnabled != "true" {
		return false, false, nil
	}

	t.ConfirmBlocks, _ = strconv.ParseInt(bt.ConfirmationCount, 10, 64)
	t.MinWithdrawalSize, _ = strconv.ParseFloat(bt.WithdrawMinSize, 64)
	t.MinWithdrawalFee, _ = strconv.ParseFloat(bt.WithdrawMinFee, 64)
	return bt.IsDepositEnabled == "true", bt.IsWithdrawEnabled == "true", nil
}

func (k *exchange) setTokenInfos(et *entity.Token) error {
	t := et.ET.(*Token)
	bt, err := k.si.getToken(t.Currency, t.ChainName)
	if err != nil {
		return err
	}

	if bt.IsChainEnabled != "true" || bt.IsDepositEnabled != "true" ||
		bt.IsWithdrawEnabled != "true" {
		return fmt.Errorf("token %s-%s is not enabled", t.Currency, t.ChainName)
	}

	t.Chain = bt.Chain
	if t.WithdrawalPrecision == 0 {
		res, err := k.readApi.WithdrawalQuotas(t.Currency, t.Chain)
		if err != nil {
			return err
		}
		wq := &kucoin.WithdrawalQuotasModel{}
		if err := res.ReadData(wq); err != nil {
			return err
		}

		t.WithdrawalPrecision = int(wq.Precision)
	}

	t.ConfirmBlocks, _ = strconv.ParseInt(bt.PreConfirmationCount, 10, 64)
	t.MinWithdrawalSize, _ = strconv.ParseFloat(bt.WithdrawMinSize, 64)
	t.MinWithdrawalFee, _ = strconv.ParseFloat(bt.WithdrawMinFee, 64)
	wp, _ := strconv.ParseInt(bt.WalletPrecision, 10, 64)
	et.Decimals = uint64(wp)
	et.ContractAddress = bt.ContractAddress
	return nil
}
