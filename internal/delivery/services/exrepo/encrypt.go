package exrepo

import (
	"encoding/json"
	"exchange-provider/internal/delivery/exchanges/dex/evm"
	"exchange-provider/internal/delivery/exchanges/dex/multichain"
	"exchange-provider/internal/delivery/exchanges/kucoin"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/utils"
	"fmt"
)

type Exchange struct {
	Id      string `bson:"id"`
	Name    string `bson:"name"`
	Configs string `bson:"configs"`
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
	case "uniswapv3", "uniswapv2", "panckakeswapv2":
		conf := ex.Configs().(*evm.Config)
		jb["hex_key"] = conf.HexKey
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
