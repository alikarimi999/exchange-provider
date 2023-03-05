package dto

import (
	"exchange-provider/internal/entity"
	"strings"
)

func ParseId(id string, pref entity.ObjectPrefix) (*entity.ObjectId, error) {
	ss := strings.Split(id, entity.IdDelimiter)
	if len(ss) != 2 {
		return nil, errInvalidID
	}
	if ss[0] == string(pref) {
		return &entity.ObjectId{Prefix: pref, Id: ss[1]}, nil
	}
	return nil, errInvalidID
}
