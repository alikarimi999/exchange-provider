package dto

import (
	"encoding/json"
	"exchange-provider/internal/entity"
)

type Order struct{ *entity.Order }

func (o *Order) MarshalBinary() ([]byte, error) {
	return json.Marshal(o.Order)
}
