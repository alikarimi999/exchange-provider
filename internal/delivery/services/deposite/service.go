package deposite

import (
	"encoding/json"
	"fmt"
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
		CoinId:   coin.CoinId,
		ChainId:  coin.ChainId,
		Exchange: exchange,
	}

	req, _ := http.NewRequest(http.MethodPost, joinUrl(d.url, path), c.reader())

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			return nil, errors.Wrap(errors.ErrNotFound, op)
		}
		return nil, errors.Wrap(errors.New(fmt.Sprintf("%d:%s", resp.StatusCode, err)), op, errors.ErrInternal)
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

func (d *depositeService) GetSupportedCoins() (map[string][]*entity.Depositcoin, error) {
	const op = errors.Op("DepositeService.Supported")
	const path = "/admin/coins/get_all"

	r := &GetSupportedCoinsRequest{[]string{"*"}}
	req, _ := http.NewRequest(http.MethodGet, joinUrl(d.url, path), r.reader())

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, errors.Wrap(errors.New(fmt.Sprintf("%d:%s", resp.StatusCode, err)), op, errors.ErrInternal)
	}

	defer resp.Body.Close()

	bod, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, op, errors.ErrInternal)
	}

	cResp := GetSupportedCoinsResponse{
		Exchanges: make(map[string]*AllCoins),
	}

	if err = json.Unmarshal(bod, &cResp); err != nil {
		return nil, errors.Wrap(err, op, errors.ErrInternal)
	}

	return cResp.Parse(), nil

}
