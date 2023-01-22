package pairconf

import (
	"context"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/logger"
	"fmt"
	"strconv"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

	spreadCache   map[string]float64
	defaultSpread float64
	db            *mongo.Collection

	minDpositCache map[string]*PairDepositLimit

	l logger.Logger
}

func NewPairConfigs(db *mongo.Database, v *viper.Viper, l logger.Logger) (entity.PairConfigs, error) {
	s := &PairConfigs{
		v: v,
		l: l,

		spreadCache:   make(map[string]float64),
		db:            db.Collection("pairconfigs"),
		defaultSpread: 0.00,

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

	r.defaultSpread = s
	return nil
}

func (r *PairConfigs) GetPairSpread(bc, qc *entity.Token) string {
	if s, ok := r.spreadCache[fmt.Sprintf("%s/%s", bc.String(), qc.String())]; ok {
		return strconv.FormatFloat(s, 'f', -1, 64)
	}
	return strconv.FormatFloat(r.defaultSpread, 'f', -1, 64)
}

func (r *PairConfigs) ChangePairSpread(bc, qc *entity.Token, s float64) error {
	if s <= 0 || s >= 1 {
		return errors.New("spread rate must be > 0 and < 1")
	}
	p := &PairSpread{Pair: fmt.Sprintf("%s/%s", bc.String(), qc.String()), Spread: s}
	_, err := r.db.InsertOne(context.Background(), p)
	if err != nil {
		return err
	}
	r.spreadCache[fmt.Sprintf("%s/%s", bc.String(), qc.String())] = s
	return nil
}

func (r *PairConfigs) ApplySpread(in, out *entity.Token, vol string) (appliedVol, spreadVol, spreadRate string, err error) {
	var rate float64
	rate, ok := r.spreadCache[fmt.Sprintf("%s/%s", in.String(), out.String())]
	if !ok {
		rate, ok = r.spreadCache[fmt.Sprintf("%s/%s", out.String(), in.String())]
		if !ok {
			rate = r.defaultSpread
		}
	}

	total, err := strconv.ParseFloat(vol, 64)
	if err != nil {
		return "", "", "", errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(err.Error()))
	}

	aVol := total * (1 - rate)
	sVol := total - aVol
	return strconv.FormatFloat(aVol, 'f', -1, 64), strconv.FormatFloat(sVol, 'f', -1, 64),
		strconv.FormatFloat(rate, 'f', -1, 64), nil

}

func (r *PairConfigs) GetAllPairsSpread() map[string]float64 {
	return r.spreadCache
}

func (r *PairConfigs) retriveSpreads() error {
	if err := r.retriveDefaultSpread(); err != nil {
		return err
	}

	pairs := []PairSpread{}
	cur, err := r.db.Find(context.Background(), bson.D{})
	if err != nil {
		return err
	}
	if err = cur.All(context.Background(), &pairs); err != nil {
		return err
	}

	for _, p := range pairs {
		if p.Spread <= 0 || p.Spread >= 1 {
			p.Spread = r.defaultSpread
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
