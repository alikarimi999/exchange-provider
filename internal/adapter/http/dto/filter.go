package dto

import (
	"exchange-provider/internal/entity"
)

type Filter struct {
	Param    string        `json:"param"`
	Operator string        `json:"operator"`
	Values   []interface{} `json:"values"`
}

func (f *Filter) ToEntity() *entity.Filter {
	return &entity.Filter{
		Param:    f.Param,
		Operator: entity.ParseFilterOperator(f.Operator),
		Values:   f.Values,
	}
}
