package pairconf

import (
	"fmt"
	"order_service/internal/entity"
	"order_service/pkg/errors"
	"strconv"
	"sync"

	"gorm.io/gorm"
)

type PairSpread struct {
	Pair   string `gorm:"primary_key"`
	Spread float64
}

type PairConfigs struct {
	sMux          *sync.Mutex
	spreadCache   map[string]float64
	defaultSpread float64
	db            *gorm.DB

	dMux             *sync.Mutex
	minDpositCache   map[string]*PairDepositLimit
	defaultMinDposit float64
}

func NewPairConfigs(db *gorm.DB) (entity.PairConfigs, error) {
	s := &PairConfigs{
		sMux:          &sync.Mutex{},
		spreadCache:   make(map[string]float64),
		db:            db,
		defaultSpread: 0.00,

		dMux:             &sync.Mutex{},
		minDpositCache:   make(map[string]*PairDepositLimit),
		defaultMinDposit: 0.00,
	}

	if err := s.retriveSpreads(); err != nil {
		return nil, err
	}

	if err := s.retriveMinDeposits(); err != nil {
		return nil, err
	}

	return s, nil
}

func (r *PairConfigs) GetDefaultSpread() string {
	r.sMux.Lock()
	defer r.sMux.Unlock()
	return strconv.FormatFloat(r.defaultSpread, 'f', -1, 64)
}

func (r *PairConfigs) ChangeDefaultSpread(s float64) error {
	if s >= 1 {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("spread rate must be less than 1"))
	}

	r.sMux.Lock()
	defer r.sMux.Unlock()
	r.defaultSpread = s
	return nil
}

func (r *PairConfigs) GetPairSpread(bc, qc *entity.Coin) string {
	r.sMux.Lock()
	defer r.sMux.Unlock()
	if s, ok := r.spreadCache[fmt.Sprintf("%s/%s", bc.String(), qc.String())]; ok {
		return strconv.FormatFloat(s, 'f', -1, 64)
	}
	return strconv.FormatFloat(r.defaultSpread, 'f', -1, 64)
}

func (r *PairConfigs) ChangePairSpread(bc, qc *entity.Coin, s float64) error {
	r.sMux.Lock()
	defer r.sMux.Unlock()
	if err := r.db.Save(&PairSpread{Pair: fmt.Sprintf("%s/%s", bc.String(), qc.String()), Spread: s}).Error; err != nil {
		return err
	}
	r.spreadCache[fmt.Sprintf("%s/%s", bc.String(), qc.String())] = s
	return nil
}

func (r *PairConfigs) ApplySpread(bc, qc *entity.Coin, size string) (remainder, rate string, err error) {
	r.sMux.Lock()
	defer r.sMux.Unlock()
	if rate, ok := r.spreadCache[fmt.Sprintf("%s/%s", bc.String(), qc.String())]; ok {
		t, err := strconv.ParseFloat(size, 64)
		if err != nil {
			return "", "", err
		}

		ff := t * rate
		re := t - ff

		return strconv.FormatFloat(re, 'f', -1, 64), strconv.FormatFloat(rate, 'f', -1, 64), nil
	}
	r.spreadCache[fmt.Sprintf("%s/%s", bc.String(), qc.String())] = 0.01

	return r.ApplySpread(bc, qc, size)
}

func (r *PairConfigs) GetAllPairsSpread() map[string]float64 {
	r.sMux.Lock()
	defer r.sMux.Unlock()
	return r.spreadCache
}

func (r *PairConfigs) retriveSpreads() error {
	pairs := []PairSpread{}
	if err := r.db.Find(&pairs).Error; err != nil {
		return err
	}
	for _, p := range pairs {
		r.spreadCache[p.Pair] = p.Spread
	}
	return nil
}
