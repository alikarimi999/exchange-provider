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

func (d *depositeService) New(userId, orderId int64, coin entity.Coin) (*entity.Deposite, error) {
	const op = errors.Op("DepositeService.New")
	c := &CreateDopsiteRequest{
		UserId:   userId,
		OrderId:  orderId,
		Currency: coin.Symbol,
		Chain:    string(coin.Chain),
	}

	req, _ := http.NewRequest(http.MethodPost, d.url, c.reader())

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
