package app

import (
	"exchange-provider/internal/entity"
)

// Write writes the given data first to the persistent storage and then to the cache.
func (o *OrderUseCase) write(data interface{}) error {
	return o.writeToPersistentStorage(data)
}

// WriteToPersistentStorage writes the given data to the persistent storage.
func (o *OrderUseCase) writeToPersistentStorage(data interface{}) error {
	switch d := data.(type) {
	case *entity.CexOrder:
		if d.Status == entity.ONew {
			return o.repo.Add(d)
		}
		return o.repo.Update(d)
	case *entity.DexOrder:
		if d.Status == entity.ONew {
			d.Status = entity.OPending
			return o.repo.Add(d)
		}
		return o.repo.Update(d)
	default:
		return nil
	}
}
