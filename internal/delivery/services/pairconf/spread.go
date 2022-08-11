package pairconf

import (
	"fmt"
	"order_service/internal/entity"
	"order_service/pkg/errors"
	"order_service/pkg/logger"
	"strconv"
	"sync"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

const (
	dsr = "default_spread_rate"
)

type PairSpread struct {
	Pair   string `gorm:"primary_key"`
	Spread float64
}

type PairConfigs struct {
	v *viper.Viper

	sMux          *sync.Mutex
	spreadCache   map[string]float64
	defaultSpread float64
	db            *gorm.DB

	dMux           *sync.Mutex
	minDpositCache map[string]*PairDepositLimit

	l logger.Logger
}

func NewPairConfigs(db *gorm.DB, v *viper.Viper, l logger.Logger) (entity.PairConfigs, error) {
	s := &PairConfigs{
		v:             v,
		l:             l,
		sMux:          &sync.Mutex{},
		spreadCache:   make(map[string]float64),
		db:            db,
		defaultSpread: 0.00,

		dMux:           &sync.Mutex{},
		minDpositCache: make(map[string]*PairDepositLimit),
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

	r.v.Set(dsr, s)
	if err := r.v.WriteConfig(); err != nil {
		return errors.Wrap(errors.NewMesssage(err.Error()))
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

func (r *PairConfigs) ApplySpread(bc, qc *entity.Coin, vol string) (appliedVol, spreadVol, spreadRate string, err error) {
	r.sMux.Lock()
	defer r.sMux.Unlock()
	var rate float64
	rate, ok := r.spreadCache[fmt.Sprintf("%s/%s", bc.String(), qc.String())]
	if !ok {
		rate = r.defaultSpread
	}

	total, err := strconv.ParseFloat(vol, 64)
	if err != nil {
		return "", "", "", errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(err.Error()))
	}

	aVol := total * (1 - rate)
	sVol := total - aVol
	return strconv.FormatFloat(aVol, 'f', -1, 64), strconv.FormatFloat(sVol, 'f', -1, 64), strconv.FormatFloat(rate, 'f', -1, 64), nil

}

func (r *PairConfigs) GetAllPairsSpread() map[string]float64 {
	r.sMux.Lock()
	defer r.sMux.Unlock()
	return r.spreadCache
}

func (r *PairConfigs) retriveSpreads() error {

	if err := r.retriveDefaultSpread(); err != nil {
		return err
	}

	pairs := []PairSpread{}
	if err := r.db.Find(&pairs).Error; err != nil {
		return errors.Wrap(errors.NewMesssage(err.Error()))
	}
	for _, p := range pairs {
		if p.Spread <= 0 || p.Spread >= 1 {
			p.Spread = r.defaultSpread
		}

		if err := r.db.Save(&p).Error; err != nil {
			return errors.Wrap(errors.NewMesssage(err.Error()))
		}
		r.spreadCache[p.Pair] = p.Spread
	}
	return nil
}

func (r *PairConfigs) retriveDefaultSpread() error {
	sr := r.v.GetFloat64(dsr)
	if sr <= 0 || sr >= 1 {
		r.v.Set(dsr, 0.001)
		if err := r.v.WriteConfig(); err != nil {
			return errors.Wrap(errors.NewMesssage(err.Error()))
		}

		r.defaultSpread = 0.001
	} else {
		r.defaultSpread = sr
	}
	return nil
}
