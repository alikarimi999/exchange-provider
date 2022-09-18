package fee

import (
	"exchange-provider/internal/entity"
	"strconv"
	"sync"

	"exchange-provider/pkg/errors"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

const (
	dfr = "default_fee_rate"
)

type UserFee struct {
	UserId int64 `gorm:"primary_key"`
	Fee    float64
}

type feeService struct {
	db *gorm.DB

	mux        *sync.Mutex
	defaultFee float64
	fees       map[int64]float64

	v *viper.Viper
}

func NewFeeService(db *gorm.DB, v *viper.Viper) (entity.FeeService, error) {

	f := &feeService{
		db:         db,
		v:          v,
		mux:        &sync.Mutex{},
		fees:       make(map[int64]float64),
		defaultFee: 0.001,
	}

	if err := f.retrievDefaultFee(); err != nil {
		return nil, err
	}

	if err := f.getFees(); err != nil {
		return nil, err
	}

	return f, nil
}

func (f *feeService) GetDefaultFee() string {
	f.mux.Lock()
	defer f.mux.Unlock()
	return strconv.FormatFloat(f.defaultFee, 'f', 6, 64)
}

func (f *feeService) ChangeDefaultFee(fee float64) error {
	if fee < 0 || fee > 1 {
		return errors.Wrap(errors.NewMesssage("fee rate must be between 0 and 1"))
	}

	f.v.Set(dfr, fee)
	if err := f.v.WriteConfig(); err != nil {
		return errors.Wrap(errors.NewMesssage(err.Error()))
	}

	f.mux.Lock()
	f.defaultFee = fee
	f.mux.Unlock()
	return nil
}

func (f *feeService) ApplyFee(userId int64, total string) (remainder, fee string, err error) {
	const op = errors.Op("FeeService.ApplyFee")

	rate := f.feeRate(userId)
	t, err := strconv.ParseFloat(total, 64)
	if err != nil {
		return "", "", errors.Wrap(op, err, errors.ErrInternal)
	}
	ff := t * rate
	re := t - ff

	return strconv.FormatFloat(re, 'f', -1, 64), strconv.FormatFloat(ff, 'f', 6, 64), nil
}

func (f *feeService) GetUserFee(userId int64) string {
	f.mux.Lock()
	defer f.mux.Unlock()
	fee := f.fees[userId]
	if fee == 0 {
		fee = f.defaultFee
	}

	return strconv.FormatFloat(fee, 'f', 6, 64)
}

func (f *feeService) ChangeUserFee(userId int64, fee float64) error {
	f.mux.Lock()
	defer f.mux.Unlock()

	if err := f.db.Save(&UserFee{UserId: userId, Fee: fee}).Error; err != nil {
		return err
	}

	f.fees[userId] = fee

	return nil
}

func (f *feeService) GetAllUsersFees() map[int64]string {
	f.mux.Lock()
	defer f.mux.Unlock()
	res := make(map[int64]string)
	for u, f := range f.fees {
		res[u] = strconv.FormatFloat(f, 'f', 6, 64)
	}
	return res
}

func (f *feeService) getFees() error {
	fees := []*UserFee{}
	if err := f.db.Find(&fees).Error; err != nil {
		return errors.Wrap(errors.NewMesssage(err.Error()))
	}

	for _, fee := range fees {
		f.fees[fee.UserId] = fee.Fee
	}
	return nil

}

func (f *feeService) feeRate(userId int64) (rate float64) {
	f.mux.Lock()
	defer f.mux.Unlock()
	fee := f.fees[userId]
	if fee == 0 {
		fee = f.defaultFee
	}
	return fee
}

func (f *feeService) retrievDefaultFee() error {
	fee := f.v.GetFloat64(dfr)

	if fee > 0 && fee < 1 {
		f.defaultFee = fee
		return nil
	}

	f.defaultFee = 0.001
	f.v.Set(dfr, f.defaultFee)
	if err := f.v.WriteConfig(); err != nil {
		return errors.Wrap(errors.NewMesssage(err.Error()))
	}
	return nil
}
