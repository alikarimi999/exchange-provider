package app

import (
	"order_service/internal/entity"
	"time"
)

type ChangeExchangeStatus struct {
	Exchange string

	Removed        []*entity.PairsErr
	PreviousStatus string
	CurrentStatus  string

	// if the exchange is stopped, this is the time it was stopped
	LastChange time.Time
}
