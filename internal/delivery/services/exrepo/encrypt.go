package exrepo

import (
	"encoding/json"
	"order_service/internal/app"
	"order_service/internal/delivery/exchanges/kucoin"
	uniswapv3 "order_service/internal/delivery/exchanges/uniswap/v3"
	"order_service/pkg/errors"
	"order_service/pkg/utils"
)

type Exchange struct {
	Id      string
	Name    string
	Configs string
	Status  string
}
type KucoinExchange struct {
	Id            string `gorm:"primary_key"`
	ApiKey        string
	ApiSecret     string
	ApiPassphrase string
	Status        string
}

func (r *ExchangeRepo) encryptConfigs(ex *app.Exchange) (*Exchange, error) {

	pub := r.prv.PublicKey

	e := &Exchange{
		Id:     ex.AccountId(),
		Name:   ex.Name(),
		Status: ex.CurrentStatus,
	}

	jb := make(jsonb)

	switch e.Name {
	case "uniswapv3":
		conf := ex.Configs().(*uniswapv3.Config)

		jb["mnemonic"] = conf.Mnemonic

	case "kucoin":
		conf := ex.Configs().(*kucoin.Configs)
		jb["api_key"] = conf.ApiKey
		jb["api_secret"] = conf.ApiSecret
		jb["api_passphrase"] = conf.ApiPassphrase

	default:
		return nil, errors.Wrap(errors.ErrBadRequest)
	}

	b, err := json.Marshal(jb)
	if err != nil {
		return nil, err
	}

	enc, err := utils.RSA_OAEP_Encrypt(string(b), pub)
	if err != nil {
		return nil, err
	}
	e.Configs = enc

	return e, nil
}
