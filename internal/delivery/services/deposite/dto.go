package deposite

import (
	"bytes"
	"encoding/json"
	"io"
	"order_service/internal/entity"
)

type CreateDopsiteRequest struct {
	UserId   int64  `json:"userId"`
	OrderId  int64  `json:"orderId"`
	Currency string `json:"currency"`
	Chain    string `json:"chain"`
	Exchange string `json:"exchange"`
}

// return io.Reader for the request body
func (r *CreateDopsiteRequest) reader() io.Reader {
	b, _ := json.Marshal(r)
	return bytes.NewReader(b)
}

type CreateDepositeResp struct {
	Id       int64  `json:"id"`
	UserId   int64  `json:"userId"`
	OrderId  int64  `json:"orderId"`
	Exchange string `json:"exchange"`
	Address  string `json:"address"`
}

func (c *CreateDepositeResp) MapToEntity() *entity.Deposite {
	return &entity.Deposite{
		Id:         c.Id,
		UserId:     c.UserId,
		OrderId:    c.OrderId,
		Exchange:   c.Exchange,
		Fullfilled: false,
		Address:    c.Address,
	}
}
