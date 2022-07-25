package dto

import "fmt"

type Deposit struct {
	Id         int64  `json:"id"`
	UserID     int64  `json:"user_id"`
	OrderId    int64  `json:"order_id"`
	DepositeId int64  `json:"deposit_id"`
	TxId       string `json:"tx_id"`
	Currency   string `json:"currency"`
	Chain      string `json:"chain"`
	Volume     string `json:"volume"`
	Status     string `json:"status"`
	Exchange   string `json:"exchange"`
}

func (d *Deposit) String() string {
	return fmt.Sprintf("%+v", *d)
}
