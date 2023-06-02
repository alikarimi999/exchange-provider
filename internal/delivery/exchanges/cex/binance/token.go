package binance

import (
	"exchange-provider/internal/entity"
	"strconv"
	"strings"
	"time"

	"github.com/adshao/go-binance/v2"
)

type Token struct {
	Coin    string `json:"coin"`
	Network string `json:"network,omitempty"  bson:"network,omitempty"`

	StableToken    string `json:"stableToken"`
	DepositAddress string `json:"depositAddress,omitempty" bson:"-"`
	DepositTag     string `json:"depositTag,omitempty" bson:"-"`

	BlockTime     time.Duration `json:"blockTime,omitempty" bson:"blockTime,omitempty"`
	ConfirmBlocks int64         `json:"confirmBlocks,omitempty" bson:"-"`
	UnLockConfirm int64         `json:"unlockConfirm,omitempty" bson:"-"`

	MinOrderSize float64 `json:"minOrderSize" bson:"-"`
	MaxOrderSize float64 `json:"maxOrderSize" bson:"-"`

	MinWithdrawalSize float64 `json:"minWithdrawalSize,omitempty" bson:"-"`
	MinWithdrawalFee  float64 `json:"minWithdrawalFee,omitempty" bson:"-"`

	WithdrawalPrecision int `json:"withdrawalPrecision,omitempty" bson:"-"`
	OrderPrecision      int `json:"orderPrecision" bson:"-"`
}

func (t *Token) Snapshot() entity.ExchangeToken {
	return &Token{
		Coin:    t.Coin,
		Network: t.Network,

		StableToken:    t.StableToken,
		DepositAddress: t.DepositAddress,
		DepositTag:     t.DepositTag,

		BlockTime:     t.BlockTime,
		ConfirmBlocks: t.ConfirmBlocks,
		UnLockConfirm: t.UnLockConfirm,

		MinOrderSize: t.MinOrderSize,
		MaxOrderSize: t.MaxOrderSize,

		MinWithdrawalSize: t.MinWithdrawalSize,
		MinWithdrawalFee:  t.MinWithdrawalFee,

		WithdrawalPrecision: t.WithdrawalPrecision,
		OrderPrecision:      t.OrderPrecision,
	}
}

func (t *Token) setInfos(n binance.Network) {
	minWithdrawalSize, _ := strconv.ParseFloat(n.WithdrawMin, 64)
	minWithdrawalFee, _ := strconv.ParseFloat(n.WithdrawFee, 64)

	ss := strings.Split(n.WithdrawIntegerMultiple, ".")
	if len(ss) != 2 {
		t.WithdrawalPrecision = 0
	} else {
		t.WithdrawalPrecision = len(ss[1])
	}

	t.ConfirmBlocks = int64(n.MinConfirm)
	t.UnLockConfirm = int64(n.UnLockConfirm)
	t.MinWithdrawalSize = minWithdrawalSize
	t.MinWithdrawalFee = minWithdrawalFee
}
