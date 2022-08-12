package app

import (
	"fmt"
	"order_service/internal/entity"
	"order_service/pkg/errors"
)

func (o *OrderUseCase) ChangeExchangeStatus(nid, status string, force bool) (*ChangeExchangeStatus, error) {
	const op = errors.Op("OrderUseCase.ChangeExchangeStatus")
	ex, err := o.exs.get(nid)
	if err != nil {
		return nil, errors.Wrap(op, err)
	}

	lt := ex.LastChangeTime
	switch status {
	case ExchangeStatusActive:

		switch ex.CurrentStatus {
		case ExchangeStatusActive:
			return &ChangeExchangeStatus{
				Exchange:       nid,
				CurrentStatus:  ExchangeStatusActive,
				PreviousStatus: ExchangeStatusActive,
				LastChange:     lt,
				Removed:        []*entity.PairsErr{},
			}, nil

		case ExchangeStatusDeactive:

			if err := o.exs.activate(nid); err != nil {
				return nil, errors.Wrap(op, err)
			}

			o.l.Info(string(op), fmt.Sprintf("exchange %s : status changed from deactive to active", nid))
			return &ChangeExchangeStatus{
				Exchange:       nid,
				CurrentStatus:  ExchangeStatusActive,
				PreviousStatus: ExchangeStatusDeactive,
				LastChange:     lt,
				Removed:        []*entity.PairsErr{},
			}, nil

		case ExchangeStatusDisable:
			res, err := ex.StartAgain()
			if err != nil {
				return nil, errors.Wrap(op, err)
			}

			if err := o.exs.activate(nid); err != nil {
				return nil, errors.Wrap(op, err)
			}

			o.l.Info(string(op), fmt.Sprintf("exchange %s : status changed from disabled to active", nid))
			return &ChangeExchangeStatus{
				Exchange:       nid,
				CurrentStatus:  ExchangeStatusActive,
				PreviousStatus: ExchangeStatusDisable,
				LastChange:     lt,
				Removed:        res.Removed,
			}, nil

		}

	case ExchangeStatusDeactive:
		switch ex.CurrentStatus {
		case ExchangeStatusActive:
			if err := o.exs.deactivate(nid); err != nil {
				return nil, errors.Wrap(op, err)
			}

			o.l.Info(string(op), fmt.Sprintf("exchange %s : status changed from active to deactive", nid))
			return &ChangeExchangeStatus{
				Exchange:       nid,
				CurrentStatus:  ExchangeStatusDeactive,
				PreviousStatus: ExchangeStatusActive,
				LastChange:     lt,
				Removed:        []*entity.PairsErr{},
			}, nil

		case ExchangeStatusDeactive:
			return &ChangeExchangeStatus{
				Exchange:       nid,
				CurrentStatus:  ExchangeStatusDeactive,
				PreviousStatus: ExchangeStatusDeactive,
				LastChange:     lt,
				Removed:        []*entity.PairsErr{},
			}, nil

		case ExchangeStatusDisable:
			res, err := ex.StartAgain()
			if err != nil {
				return nil, errors.Wrap(op, err)
			}

			if err := o.exs.deactivate(nid); err != nil {
				return nil, errors.Wrap(op, err)
			}

			o.l.Info(string(op), fmt.Sprintf("exchange %s : status changed from disabled to deactive", nid))
			return &ChangeExchangeStatus{
				Exchange:       nid,
				CurrentStatus:  ExchangeStatusDeactive,
				PreviousStatus: ExchangeStatusDisable,
				LastChange:     lt,
				Removed:        res.Removed,
			}, nil
		}

	case ExchangeStatusDisable:

		switch ex.CurrentStatus {
		case ExchangeStatusActive:
			if !force {
				// First, check whether there is a request being processed for this exchange.
				// if there is, we cannot disable it
				t, err := o.totalPendingOrders(ex.Exchange)
				if err != nil {
					return nil, errors.Wrap(op, err)
				}

				if t > 0 {
					o.l.Info(string(op), fmt.Sprintf("unable to disable exchange %s because there are %d pending orders", nid, t))
					return nil, errors.Wrap(errors.ErrBadRequest,
						errors.NewMesssage(fmt.Sprintf("exchange %s has %d pending orders, so you can't disable it, unless force it", nid, t)))
				}
			}

			if err := o.exs.disable(nid); err != nil {
				return nil, errors.Wrap(op, err)
			}

			o.l.Info(string(op), fmt.Sprintf("exchange %s : status changed from active to disabled", nid))
			return &ChangeExchangeStatus{
				Exchange:       nid,
				CurrentStatus:  ExchangeStatusDisable,
				PreviousStatus: ExchangeStatusActive,
				LastChange:     lt,
				Removed:        []*entity.PairsErr{},
			}, nil

		case ExchangeStatusDeactive:
			if !force {
				// First, check whether there is a request being processed for this exchange.
				// if there is, we cannot disable it
				t, err := o.totalPendingOrders(ex.Exchange)
				if err != nil {
					return nil, errors.Wrap(op, err)
				}

				if t > 0 {
					o.l.Info(string(op), fmt.Sprintf("unable to disable exchange %s because there are %d pending orders", nid, t))
					return nil, errors.Wrap(errors.ErrBadRequest,
						errors.NewMesssage(fmt.Sprintf("exchange %s has %d pending orders, so you can't disable it, unless force it", nid, t)))
				}
			}

			if err := o.exs.disable(nid); err != nil {
				return nil, errors.Wrap(op, err)
			}

			o.l.Info(string(op), fmt.Sprintf("exchange %s : status changed from deactive to disabled", nid))
			return &ChangeExchangeStatus{
				Exchange:       nid,
				CurrentStatus:  ExchangeStatusDisable,
				PreviousStatus: ExchangeStatusDeactive,
				LastChange:     lt,
				Removed:        []*entity.PairsErr{},
			}, nil

		case ExchangeStatusDisable:
			return &ChangeExchangeStatus{
				Exchange:       nid,
				CurrentStatus:  ExchangeStatusDisable,
				PreviousStatus: ExchangeStatusDisable,
				LastChange:     lt,
				Removed:        []*entity.PairsErr{},
			}, nil

		}
	}

	return nil, errors.Wrap(errors.ErrBadRequest, errors.New(fmt.Sprintf("incorrect status `%s`", status)))
}

func (o *OrderUseCase) RemoveExchange(nid string, force bool) error {
	op := "OrderUseCase.RemoveExchange"
	ex, err := o.exs.get(nid)
	if err != nil {
		return errors.Wrap(op, err)
	}

	if ex.CurrentStatus != ExchangeStatusDisable {
		if !force {
			// First, check whether there is a request being processed for this exchange.
			// if there is, we cannot disable it
			t, err := o.totalPendingOrders(ex.Exchange)
			if err != nil {
				return errors.Wrap(op, err)
			}

			if t > 0 {
				o.l.Info(string(op), fmt.Sprintf("unable to remove exchange %s because there are %d pending orders", nid, t))
				return errors.Wrap(errors.ErrBadRequest,
					errors.NewMesssage(fmt.Sprintf("exchange %s has %d pending orders, so you can't remove it, unless force it", nid, t)))
			}
		}
		ex.Stop()
	}

	if err := o.exs.remove(nid); err != nil {
		return errors.Wrap(op, err)
	}

	o.l.Info(string(op), fmt.Sprintf("exchange %s removed", nid))
	return nil
}
