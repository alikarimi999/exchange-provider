package exrepo

import (
	"encoding/json"
	"exchange-provider/internal/delivery/exchanges/cex/kucoin"
	"exchange-provider/internal/delivery/exchanges/cex/swapspace"
	"exchange-provider/internal/delivery/exchanges/dex/evm"
	"exchange-provider/internal/delivery/exchanges/dex/multichain"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/utils"
	"fmt"
)

type Exchange struct {
	Id      uint   `bson:"_id"`
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
		jb["hexKey"] = conf.HexKey
		jb["network"] = conf.Network

	case "multichain":
		conf := ex.Configs().(*multichain.Config)
		jb["mnemonic"] = conf.Mnemonic

	case "kucoin":
		conf := ex.Configs().(*kucoin.Configs)
		jb["read.apiKey"] = conf.ReadApi.ApiKey
		jb["read.apiSecret"] = conf.ReadApi.ApiSecret
		jb["read.apiPassphrase"] = conf.ReadApi.ApiPassphrase

		jb["write.apiKey"] = conf.WriteApi.ApiKey
		jb["write.apiSecret"] = conf.WriteApi.ApiSecret
		jb["write.apiPassphrase"] = conf.WriteApi.ApiPassphrase

	case "swapspace":
		conf := ex.Configs().(*swapspace.Config)
		jb["apiKey"] = conf.ApiKey

	default:
		return nil, errors.Wrap(op, errors.ErrBadRequest, fmt.Errorf("'%d' unknown exchange Id", e.Id))
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
