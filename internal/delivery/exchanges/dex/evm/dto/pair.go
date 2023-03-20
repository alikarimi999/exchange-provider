package dto

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
)

type AddPairsRequest struct {
	Pairs []*Pair `json:"pairs"`
}

func (req AddPairsRequest) Validate() error {
	for _, p := range req.Pairs {
		if p.T1.Symbol == "" || p.T1.Decimals == 0 || p.T2.Symbol == "" || p.T2.Decimals == 0 {
			return errors.New("Invalid pair token")
		}
	}
	return nil
}

type Pair struct {
	T1 *entity.Token `json:"t1"`
	T2 *entity.Token `json:"t2"`
}

func (p *Pair) String() string {
	return fmt.Sprintf("%s/%s", p.T1, p.T2)
}
