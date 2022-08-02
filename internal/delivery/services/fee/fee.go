package fee

import (
	"order_service/internal/entity"
	"strconv"
	"sync"

	"order_service/pkg/errors"

	"github.com/go-redis/redis/v9"
)

type feeService struct {
	url string
	r   *redis.Client

	mux *sync.Mutex
	fee float64 //atomic
}

func NewFeeService() entity.FeeService {
	return &feeService{
		mux: &sync.Mutex{},
		fee: 0.001,
	}
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

func (f *feeService) feeRate(userId int64) (rate float64) {
	f.mux.Lock()
	defer f.mux.Unlock()
	return f.fee
}

func (f *feeService) GetFee() string {
	f.mux.Lock()
	defer f.mux.Unlock()
	return strconv.FormatFloat(f.fee, 'f', 6, 64)
}

func (f *feeService) ChangeFee(fee float64) {
	f.mux.Lock()
	defer f.mux.Unlock()
	f.fee = fee
}
