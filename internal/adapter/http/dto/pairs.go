package dto

import (
	"exchange-provider/internal/entity"
	"fmt"
)

type UserPair struct {
	T1               string  `json:"t1"`
	T2               string  `json:"t2"`
	Price1           string  `json:"price1,omitempty"`
	Price2           string  `json:"price2,omitempty"`
	FeeRate          string  `json:"feeRate,omitempty"`
	TransferFee1     string  `json:"transferFee1,omitempty"`
	TransferFee2     string  `json:"transferFee2,omitempty"`
	MinDepositToken1 float64 `json:"minDepositToken1,omitempty"`
	MinDepositToken2 float64 `json:"minDepositToken2,omitempty"`
	Msg              string  `json:"message,omitempty"`
}

func EntityToPairUser(p *entity.Pair) *UserPair {
	pair := &UserPair{
		T1:      p.T1.String(),
		T2:      p.T2.String(),
		Price1:  p.Price1,
		Price2:  p.Price2,
		FeeRate: p.FeeRate,
	}
	return pair
}

type GetPairsToUserResponse struct {
	Pairs []*UserPair `json:"pairs"`
}

type Pair struct {
	T1 *entity.Token
	T2 *entity.Token
}

func (p *Pair) String() string {
	return fmt.Sprintf("%s/%s", p.T1.String(), p.T2.String())
}

// type GetPairsToUserRequest struct {
// 	Pairs []*UserPair `json:"pairs"`
// }

// func (r *GetPairsToUserRequest) Parse() ([]*Pair, error) {
// 	pairs := []*Pair{}
// 	for _, p := range r.Pairs {
// 		bc, err := ParseToken(p.T1)
// 		if err != nil {
// 			return nil, err
// 		}
// 		qc, err := ParseToken(p.T2)
// 		if err != nil {
// 			return nil, err
// 		}
// 		pairs = append(pairs, &Pair{
// 			T1: bc,
// 			T2: qc,
// 		})
// 	}
// 	return pairs, nil
// }
