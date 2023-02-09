package evm

import (
	"exchange-provider/internal/delivery/exchanges/dex/evm/dto"
	"exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
)

var errPairNotSupport = fmt.Errorf("pair not supported")

func (d *EvmDex) AddPairs(data interface{}) (*entity.AddPairsResult, error) {
	req := data.(*dto.AddPairsRequest)

	res := &entity.AddPairsResult{}

	for _, p := range req.Pairs {
		if d.pairsRepo.Exists(d.Id(), p.T1, p.T2) {
			res.Existed = append(res.Existed, p.String())
			continue
		}

		ps := make([]*entity.Pair, len(req.Pairs))
		for i, p := range req.Pairs {
			ps[i] = &entity.Pair{T1: &entity.PairToken{Token: p.T1}, T2: &entity.PairToken{Token: p.T2}}
		}

		ps, err := d.price(ps...)
		if err != nil {
			return nil, err
		}
		d.pairsRepo.Add(d, ps...)
	}
	return res, nil
}

func (d *EvmDex) Support(t1, t2 *entity.Token) bool {
	if t1.ChainId != d.TokenStandard || t2.ChainId != d.TokenStandard {
		return false
	}
	T1 := t1.Snapshot()
	T2 := t2.Snapshot()

	if T1.TokenId == d.NativeToken {
		T1.TokenId = d.WrappedNativeToken
	} else if T2.TokenId == d.NativeToken {
		T2.TokenId = d.WrappedNativeToken
	}

	return d.pairsRepo.Exists(d.Id(), T1, T2)
}

func (d *EvmDex) RemovePair(t1, t2 *entity.Token) error {
	if t1.ChainId != d.TokenStandard || t2.ChainId != d.TokenStandard {
		return errors.New("pair not found")
	}
	return d.pairsRepo.Remove(d.Id(), t1, t2)
}

func (d *EvmDex) retreivePairs() ([]*entity.Pair, error) {
	v := viper.New()
	v.SetConfigFile(d.PairsFile)
	if err := v.ReadInConfig(); err != nil {
		// create config file if not exists
		if err := v.WriteConfigAs(d.PairsFile); err != nil {
			return nil, err
		}
	}
	pairs := v.GetStringMap("Pairs")
	ps := []*entity.Pair{}
	for _, pm := range pairs {
		pi := pm.(map[string]interface{})

		tA := interfaceToToken(pi["t1"])
		tB := interfaceToToken(pi["t2"])

		t1 := types.Token{}
		t2 := types.Token{}
		if tA.Address.Hash().Big().Cmp(tB.Address.Hash().Big()) == -1 {
			t1 = tA
			t2 = tB
		} else {
			t2 = tA
			t1 = tB
		}

		p := &entity.Pair{
			T1: t1.ToEntity(d.TokenStandard),
			T2: t2.ToEntity(d.TokenStandard),
		}

		ps = append(ps, p)
	}

	return ps, nil
}

func (d *EvmDex) removePairs(ids ...string) error {
	v := viper.New()
	v.SetConfigFile(d.PairsFile)
	if err := v.ReadInConfig(); err != nil {
		// create config file if not exists
		if err := v.WriteConfigAs(d.PairsFile); err != nil {
			return err
		}
	}
	for _, id := range ids {
		delete(v.Get("Pairs").(map[string]interface{}),
			strings.ToLower(id))
	}
	if err := v.WriteConfig(); err != nil {
		return err
	}
	return nil
}

func interfaceToToken(ti interface{}) types.Token {
	tm := ti.(map[string]interface{})
	return types.Token{
		Name:     tm["name"].(string),
		Symbol:   tm["symbol"].(string),
		Address:  common.HexToAddress(tm["address"].(string)),
		Decimals: int(tm["decimals"].(float64)),
		ChainId:  int64(tm["chainid"].(float64)),
	}

}
