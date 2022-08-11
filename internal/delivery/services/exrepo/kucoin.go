package exrepo

import (
	"order_service/internal/app"
	"order_service/internal/delivery/exchanges/kucoin"
	"order_service/pkg/errors"
	"order_service/pkg/utils"
)

type KucoinExchange struct {
	Id            string `gorm:"primary_key"`
	ApiKey        string
	ApiSecret     string
	ApiPassphrase string
	Status        string
}

func (a *ExchangeRepo) encryptKucoinConfigs(ex *app.Exchange) (*KucoinExchange, error) {
	pub := a.prv.PublicKey
	cfg := ex.Configs().(*kucoin.Configs)

	key, err := utils.RSA_OAEP_Encrypt(cfg.ApiKey, pub)
	if err != nil {
		return nil, errors.Wrap(errors.NewMesssage(err.Error()))
	}

	secret, err := utils.RSA_OAEP_Encrypt(cfg.ApiSecret, pub)
	if err != nil {
		return nil, errors.Wrap(errors.NewMesssage(err.Error()))
	}

	passphrase, err := utils.RSA_OAEP_Encrypt(cfg.ApiPassphrase, pub)
	if err != nil {
		return nil, errors.Wrap(errors.NewMesssage(err.Error()))
	}

	return &KucoinExchange{
		Id:            ex.AccountId(),
		ApiKey:        key,
		ApiSecret:     secret,
		ApiPassphrase: passphrase,
		Status:        ex.CurrentStatus,
	}, nil

}

func (a *ExchangeRepo) decryptKucoinConfigs(cfg *KucoinExchange) (*kucoin.Configs, error) {
	priv := a.prv
	key, err := utils.RSA_OAEP_Decrypt(cfg.ApiKey, *priv)
	if err != nil {
		return nil, errors.Wrap(errors.NewMesssage(err.Error()))
	}

	secret, err := utils.RSA_OAEP_Decrypt(cfg.ApiSecret, *priv)
	if err != nil {
		return nil, errors.Wrap(errors.NewMesssage(err.Error()))
	}

	passphrase, err := utils.RSA_OAEP_Decrypt(cfg.ApiPassphrase, *priv)
	if err != nil {
		return nil, errors.Wrap(errors.NewMesssage(err.Error()))
	}

	return &kucoin.Configs{
		ApiKey:        key,
		ApiSecret:     secret,
		ApiPassphrase: passphrase,
	}, nil
}
