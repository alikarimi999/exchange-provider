package dto

import "order_service/internal/app"

type Account struct {
	Status string      `json:"status"`
	Conf   interface{} `json:"configs"`
}

type GetAllExchangesResponse struct {
	Exchanges map[string]*Account `json:"exchanges,omitempty"`
	Msgs      []string            `json:"messages"`
}

type ChangeExchangeStatusResponse struct {
	Exchange string `json:"exchange"`

	PreviousStatus string `json:"previous_status"`
	CurrentStatus  string `json:"current_status"`

	LastChange   string      `json:"last_change"`
	RemovedPairs []*PairsErr `json:"removed_pairs,omitempty"`
}

func (r *ChangeExchangeStatusResponse) FromEntity(e *app.ChangeExchangeStatus) {
	r.Exchange = e.Exchange
	r.PreviousStatus = e.PreviousStatus
	r.CurrentStatus = e.CurrentStatus
	r.LastChange = e.LastChange.Format("2006-01-02 15:04:05")
	r.RemovedPairs = make([]*PairsErr, 0, len(e.Removed))
	for _, p := range e.Removed {
		r.RemovedPairs = append(r.RemovedPairs, &PairsErr{
			Pair: p.Pair,
			Err:  p.Err.Error(),
		})
	}
}
