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
	Id       int64  `json:"deposit_id"`
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

type AllCoins struct {
	Coins []*ExchangeCoin `json:"coins"`
	Msg   string          `json:"message,omitempty"`
}

func (r *AllCoins) ToDepositCoins() []*entity.Depositcoin {
	var coins []*entity.Depositcoin
	for _, v := range r.Coins {
		coins = append(coins, v.ToDepositCoin())
	}
	return coins
}

type ExchangeCoin struct {
	*Coin
	Address   string `json:"address"`
	SetChain  bool   `json:"set_chain"`
	BlockTime string `json:"block_time"`
	Confirms  int    `json:"confirms"`
	Precision int    `json:"precision"`
}

func (c *ExchangeCoin) ToDepositCoin() *entity.Depositcoin {
	return &entity.Depositcoin{
		CoinId:   c.CoinId,
		ChainId:  c.ChainId,
		SetChain: c.SetChain,
	}
}

type GetSupportedCoinsRequest struct {
	Exchanges []string `json:"exchanges"`
}

func (r *GetSupportedCoinsRequest) reader() io.Reader {
	b, _ := json.Marshal(r)
	return bytes.NewReader(b)
}

type GetSupportedCoinsResponse struct {
	Exchanges map[string]*AllCoins `json:"exchanges"`
}

func (r *GetSupportedCoinsResponse) Parse() map[string][]*entity.Depositcoin {
	m := make(map[string][]*entity.Depositcoin)
	for k, v := range r.Exchanges {
		m[k] = v.ToDepositCoins()
	}
	return m
}
