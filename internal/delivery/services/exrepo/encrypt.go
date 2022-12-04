package exrepo

import (
	"encoding/json"
	"exchange-provider/internal/delivery/exchanges/dex"
	"exchange-provider/internal/delivery/exchanges/dex/multichain"
	"exchange-provider/internal/delivery/exchanges/kucoin"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/utils"
	"fmt"
)

type Exchange struct {
	Id      string
	Name    string
	Configs string
}
type KucoinExchange struct {
	Id            string `gorm:"primary_key"`
	ApiKey        string
	ApiSecret     string
	ApiPassphrase string
	Status        string
}

func (r *ExchangeRepo) encryptConfigs(ex entity.Exchange) (*Exchange, error) {
	op := errors.Op("ExchangeRepo.encryptConfigs")
	pub := r.prv.PublicKey

	e := &Exchange{
		Id:   ex.Id(),
		Name: ex.Name(),
	}

	jb := make(jsonb)

	switch e.Name {
	case "uniswapv3", "panckakeswapv2":
		conf := ex.Configs().(*dex.Config)
		jb["mnemonic"] = conf.Mnemonic
		jb["network"] = conf.Network

	case "multichain":
		conf := ex.Configs().(*multichain.Config)
		jb["mnemonic"] = conf.Mnemonic

	case "kucoin":
		conf := ex.Configs().(*kucoin.Configs)
		jb["api_key"] = conf.ApiKey
		jb["api_secret"] = conf.ApiSecret
		jb["api_passphrase"] = conf.ApiPassphrase

	default:
		return nil, errors.Wrap(op, errors.ErrBadRequest, fmt.Errorf("'%s' unknown exchange Id", e.Id))
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
