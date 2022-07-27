package deposite

import (
	"bytes"
	"encoding/json"
	"io"
	"order_service/internal/entity"
)

type Coin struct {
	CoinId  string `json:"coin_id"`
	ChainId string `json:"chain_id"`
}

type CreateDopsiteRequest struct {
	UserId   int64  `json:"user_id"`
	OrderId  int64  `json:"order_id"`
	CoinId   string `json:"coin_id"`
	ChainId  string `json:"chain_id"`
	Exchange string `json:"exchange"`
}

// return io.Reader for the request body
func (r *CreateDopsiteRequest) reader() io.Reader {
	b, _ := json.Marshal(r)
	return bytes.NewReader(b)
}

type CreateDepositeResp struct {
	Id       int64  `json:"id"`
	UserId   int64  `json:"user_id"`
	OrderId  int64  `json:"order_id"`
	Exchange string `json:"exchange"`
	Address  string `json:"address"`
}

func (c *CreateDepositeResp) MapToEntity() *entity.Deposit {
	return &entity.Deposit{
		Id:         c.Id,
		UserId:     c.UserId,
		OrderId:    c.OrderId,
		Exchange:   c.Exchange,
		Fullfilled: false,
		Address:    c.Address,
	}
}

type SupportedRequest struct {
	Exchange string  `json:"exchange"`
	Coins    []*Coin `json:"coins"`
}

// return io.Reader for the request body
func (r *SupportedRequest) reader() io.Reader {
	b, _ := json.Marshal(r)
	return bytes.NewReader(b)
}

type supportCoin struct {
	*Coin
	Supported bool `json:"supported"`
}

func (c *supportCoin) MapToEntity() *entity.Coin {
	return &entity.Coin{
		Id: c.CoinId,
		Chain: &entity.Chain{
			Id: c.ChainId,
		},
	}
}

type SupportedRespons struct {
	Coins []*supportCoin `json:"coins"`
}
