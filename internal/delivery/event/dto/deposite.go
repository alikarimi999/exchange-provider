package dto

import "fmt"

type Deposite struct {
	Id         int64  `json:"id"`
	UserID     int64  `json:"userId"`
	OrderId    int64  `json:"orderId"`
	DepositeId int64  `json:"depositeId"`
	TxId       string `json:"txId"`
	Currency   string `json:"currency"`
	Chain      string `json:"chain"`
	Volume     string `json:"volume"`
	Status     string `json:"status"`
	Exchange   string `json:"exchange"`
}

func (d *Deposite) String() string {
	return fmt.Sprintf("%+v", *d)
}
