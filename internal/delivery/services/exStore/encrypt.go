package store

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
)

type Exchange struct {
	Id      uint          `bson:"_id"`
	Name    string        `bson:"name"`
	Enable  bool          `bson:"enable"`
	Type    entity.ExType `bson:"type"`
	Configs string        `bson:"configs"`
}

func (r *exchangeRepo) encryptConfigs(ex entity.Exchange, cfg interface{}) (*Exchange, error) {
	pub := r.prv.PublicKey

	e := &Exchange{
		Id:     ex.Id(),
		Type:   ex.Type(),
		Enable: ex.IsEnable(),
		Name:   ex.Name(),
	}

	b, err := bson.Marshal(cfg)
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
