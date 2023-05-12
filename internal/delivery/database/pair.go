package database

import (
	"exchange-provider/internal/entity"
	"fmt"
	"strings"
)

func pairFromString(id string) (*entity.TokenId, *entity.TokenId, error) {
	id = strings.ToUpper(id)
	ss := strings.Split(strings.ToUpper(id), "/")
	if len(ss) == 1 {
		in, err := string2TokenId(ss[0])
		if err != nil {
			return nil, nil, err
		}
		return in, nil, nil
	} else if len(ss) == 2 {
		in, err := string2TokenId(ss[0])
		if err != nil {
			return nil, nil, err
		}
		out, err := string2TokenId(ss[1])
		if err != nil {
			return nil, nil, err
		}
		return in, out, nil
	}
	return nil, nil, invalidErr(fmt.Sprintf("pairId %s", id), id)
}

func string2TokenId(id string) (*entity.TokenId, error) {
	if id == "" {
		return nil, nil
	}
	ts := strings.Split(id, "-")
	if len(ts) != 3 {
		return nil, invalidErr(fmt.Sprintf("token %s", id), id)
	}
	return &entity.TokenId{
		Symbol:   ts[0],
		Standard: ts[1],
		Network:  ts[2],
	}, nil
}
