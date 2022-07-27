package deposite

import (
	"encoding/json"
	"io"
	"net/http"
	"order_service/internal/entity"

	"order_service/pkg/errors"
)

type depositeService struct {
	url string
}

func NewDepositeService(url string) entity.DepositeService {
	return &depositeService{url: url}
}

func (d *depositeService) New(userId, orderId int64, coin *entity.Coin, exchange string) (*entity.Deposit, error) {
	const op = errors.Op("DepositeService.New")
	const path = "deposites"

	c := &CreateDopsiteRequest{
		UserId:   userId,
		OrderId:  orderId,
		CoinId:   coin.Id,
		ChainId:  coin.Chain.Id,
		Exchange: exchange,
	}

	req, _ := http.NewRequest(http.MethodPost, joinUrl(d.url, path), c.reader())

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, errors.Wrap(err, op, errors.ErrInternal)
	}

	defer resp.Body.Close()

	bod, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, op, errors.ErrInternal)
	}

	cResp := CreateDepositeResp{}

	if err = json.Unmarshal(bod, &cResp); err != nil {
		return nil, errors.Wrap(err, op, errors.ErrInternal)
	}

	return cResp.MapToEntity(), nil

}

func (d *depositeService) Supported(exchange string, coins ...*entity.Coin) ([]*entity.Coin, error) {
	const op = errors.Op("DepositeService.Supported")
	const path = "/admin/coins/supported"

	sr := &SupportedRequest{
		Exchange: exchange,
	}

	for _, coin := range coins {
		sr.Coins = append(sr.Coins, &Coin{
			CoinId:  coin.Id,
			ChainId: coin.Chain.Id,
		})
	}

	req, _ := http.NewRequest(http.MethodPost, joinUrl(d.url, path), sr.reader())

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, errors.Wrap(err, op, errors.ErrInternal)
	}

	defer resp.Body.Close()

	bod, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, op, errors.ErrInternal)
	}

	res := SupportedRespons{}

	if err = json.Unmarshal(bod, &res); err != nil {
		return nil, errors.Wrap(err, op, errors.ErrInternal)
	}

	var cs []*entity.Coin
	for _, c := range res.Coins {
		if c.Supported {
			cs = append(cs, c.MapToEntity())
		}
	}

	return cs, nil

}
