package exrepo

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
)

type Exchange struct {
	Id      uint   `bson:"_id"`
	Name    string `bson:"name"`
	Configs string `bson:"configs"`
}

func (r *ExchangeRepo) encryptConfigs(ex entity.Exchange) (*Exchange, error) {
	// op := errors.Op("ExchangeRepo.encryptConfigs")
	pub := r.prv.PublicKey

	e := &Exchange{
		Id:   ex.Id(),
		Name: ex.Name(),
	}

	b, err := bson.Marshal(ex.Configs())
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
