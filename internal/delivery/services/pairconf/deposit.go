package pairconf

import (
	"fmt"
	"exchange-provider/internal/entity"
)

type PairDepositLimit struct {
	Pair  string `gorm:"primary_key"`
	MinBc float64
	MinQc float64
}

func (c *PairConfigs) PairMinDeposit(bc, qc *entity.Coin) (minBc, minQc float64) {
	c.dMux.Lock()
	defer c.dMux.Unlock()

	if v, ok := c.minDpositCache[fmt.Sprintf("%s/%s", bc.String(), qc.String())]; ok {
		return v.MinBc, v.MinQc
	}
	return 0, 0
}

func (c *PairConfigs) ChangeMinDeposit(bc, qc *entity.Coin, minBc, minQc float64) error {
	c.dMux.Lock()
	defer c.dMux.Unlock()

	m := &PairDepositLimit{
		Pair:  fmt.Sprintf("%s/%s", bc.String(), qc.String()),
		MinBc: minBc,
		MinQc: minQc,
	}

	if err := c.db.Save(m).Error; err != nil {
		return err
	}

	c.minDpositCache[m.Pair] = m
	return nil
}

func (r *PairConfigs) AllMinDeposit() []*entity.PairMinDeposit {
	r.dMux.Lock()
	defer r.dMux.Unlock()

	var m []*entity.PairMinDeposit
	for _, v := range r.minDpositCache {
		m = append(m, &entity.PairMinDeposit{
			Pair:         v.Pair,
			MinBaseCoin:  v.MinBc,
			MinQouteCoin: v.MinQc,
		})
	}
	return m
}

func (r *PairConfigs) retriveMinDeposits() error {
	var pairs []*PairDepositLimit
	if err := r.db.Find(&pairs).Error; err != nil {
		return err
	}
	for _, v := range pairs {
		r.minDpositCache[v.Pair] = v
	}
	return nil
}
