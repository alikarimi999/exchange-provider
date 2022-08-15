package deposite

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"order_service/internal/entity"
	"sync"

	"order_service/pkg/errors"
	"order_service/pkg/logger"

	"github.com/spf13/viper"
)

type DepositServiceConfigs struct {
	Url string `json:"url"`
	STP string `json:"set_tx_id_path"`
	NDP string `json:"new_deposit_path"`
}
type depositeService struct {
	mux *sync.Mutex
	v   *viper.Viper
	l   logger.Logger

	c *DepositServiceConfigs
}

func NewDepositeService(v *viper.Viper, l logger.Logger) entity.DepositeService {

	return &depositeService{
		mux: &sync.Mutex{},
		v:   v,
		l:   l,
		c: &DepositServiceConfigs{
			Url: v.GetString("deposit_service.url"),
			STP: v.GetString("deposit_service.set_tx_id_path"),
			NDP: v.GetString("deposit_service.new_deposit_path"),
		},
	}
}

func (d *depositeService) ChangeConfigs(cfg interface{}) error {
	d.mux.Lock()
	defer d.mux.Unlock()
	c, ok := cfg.(*DepositServiceConfigs)
	if ok {
		if c.Url == "" || c.STP == "" || c.NDP == "" {
			return errors.Wrap(errors.NewMesssage("url or path is empty"))
		}
		if c.Url == d.c.Url && c.STP == d.c.STP && c.NDP == d.c.NDP {
			return errors.Wrap(errors.NewMesssage("not changed"))
		}

		d.v.Set("deposit_service", c)
		if err := d.v.WriteConfig(); err != nil {
			return err
		}

		d.c = &DepositServiceConfigs{
			Url: c.Url,
			STP: c.STP,
			NDP: c.NDP,
		}
		return nil
	}
	return errors.Wrap(errors.NewMesssage("invalid config"))
}

func (d *depositeService) Configs() interface{} {
	return &DepositServiceConfigs{
		Url: d.c.Url,
		STP: d.c.STP,
		NDP: d.c.NDP,
	}
}

func (d *depositeService) New(userId, orderId int64, coin *entity.Coin, exchange string) (*entity.Deposit, error) {
	const op = errors.Op("DepositeService.New")

	c := &CreateDopsiteRequest{
		UserId:   userId,
		OrderId:  orderId,
		CoinId:   coin.CoinId,
		ChainId:  coin.ChainId,
		Exchange: exchange,
	}

	req, err := d.newDepositRequest(c.reader())
	if err != nil {
		d.l.Error(string(op), err.Error())
		return nil, errors.Wrap(errors.ErrInternal)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		d.l.Error(string(op), err.Error())
		return nil, errors.Wrap(errors.ErrInternal)
	}

	if resp.StatusCode != http.StatusOK {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			d.l.Error(string(op), fmt.Sprintf("status code: %d\nerror ( %s )", resp.StatusCode, err.Error()))
			return nil, errors.Wrap(errors.ErrInternal)
		}
		d.l.Error(string(op), fmt.Sprintf("status code: %d\nbody ( %s )", resp.StatusCode, string(b)))
		return nil, errors.Wrap(errors.ErrInternal)
	}

	defer resp.Body.Close()

	bod, err := io.ReadAll(resp.Body)
	if err != nil {
		d.l.Error(string(op), err.Error())
		return nil, errors.Wrap(errors.ErrInternal)
	}

	cResp := CreateDepositeResp{}
	if err = json.Unmarshal(bod, &cResp); err != nil {
		d.l.Error(string(op), err.Error())
		return nil, errors.Wrap(errors.ErrInternal)
	}

	return cResp.MapToEntity(), nil

}

func (d *depositeService) SetTxId(userId, orderId, depositeId int64, txId string) error {
	const op = errors.Op("DepositeService.SetTxId")

	r := &SetTxIdRequest{
		UserId:     userId,
		OrderId:    orderId,
		DepositeId: depositeId,
		TxId:       txId,
	}

	req, err := d.setTxIdRequest(r.reader())
	if err != nil {
		d.l.Error(string(op), err.Error())
		return errors.Wrap(errors.ErrInternal)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		d.l.Error(string(op), err.Error())
		return errors.Wrap(errors.ErrInternal)
	}

	if resp.StatusCode != http.StatusOK {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			d.l.Error(string(op), fmt.Sprintf("status code: %d\nerror ( %s )", resp.StatusCode, err.Error()))
			return errors.Wrap(errors.ErrInternal)
		}
		d.l.Error(string(op), fmt.Sprintf("status code: %d\nbody ( %s )", resp.StatusCode, string(b)))
		return errors.Wrap(errors.ErrInternal)
	}

	return nil
}
