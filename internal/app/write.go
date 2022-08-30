package app

import (
	"order_service/internal/entity"
)

// Write writes the given data first to the persistent storage and then to the cache.
func (o *OrderUseCase) write(data interface{}) error {

	// Write to persistent storage.
	if err := o.writeToPersistentStorage(data); err != nil {
		return err
	}

	// Write to cache.
	if err := o.writeToCache(data); err != nil {
		return err
	}

	return nil

}

// WriteToPersistentStorage writes the given data to the persistent storage.
func (o *OrderUseCase) writeToPersistentStorage(data interface{}) error {
	switch d := data.(type) {
	case *entity.UserOrder:
		if d.Status == entity.OrderStatusOpen {
			return o.repo.Add(d)
		}
		return o.repo.Update(d)
	case *entity.Deposit:
		return o.repo.UpdateDeposit(d)
	default:
		return nil
	}
}

// WriteToCache writes the given data to the cache.
func (o *OrderUseCase) writeToCache(data interface{}) error {
	switch d := data.(type) {
	case *entity.UserOrder:
		if d.Status == entity.OrderStatusSucceed || d.Status == entity.OrderStatusFailed || d.Status == entity.OsDepositFailed {
			return nil
		}

		return o.cache.Add(d)
	case *entity.Withdrawal:
		if d.Status == entity.WithdrawalPending {
			return o.cache.AddPendingWithdrawal(d)
		}
		return nil
	case *entity.Deposit:
		return o.cache.UpdateDeposit(d)
	default:
		return nil
	}
}