package fee

import (
	"order_service/internal/entity"
	"strconv"

	"order_service/pkg/errors"

	"github.com/go-redis/redis/v9"
)

type feeService struct {
	url string
	r   *redis.Client
}

func NewFeeService() entity.FeeService {
	return &feeService{}
}

func (f *feeService) ApplyFee(userId int64, total string) (remainder, fee string, err error) {
	const op = errors.Op("FeeService.ApplyFee")

	rate, _ := f.feeRate(userId)
	t, err := strconv.ParseFloat(total, 64)
	if err != nil {
		return "", "", errors.Wrap(op, err, errors.ErrInternal)
	}
	ff := t * rate
	re := t - ff

	return strconv.FormatFloat(re, 'f', 6, 64), strconv.FormatFloat(ff, 'f', 6, 64), nil
}

func (f *feeService) feeRate(userId int64) (rate float64, err error) {
	return 0.001, nil
}
