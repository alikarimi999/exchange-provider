package app

import (
	"exchange-provider/internal/entity"
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
	case *entity.CexOrder:
		if d.Status == entity.ONew {
			return o.repo.Add(d)
		}
		return o.repo.Update(d)
	case *entity.EvmOrder:
		if d.Status == entity.ONew {
			return o.repo.Add(d)
		}
		return o.repo.Update(d)
	default:
		return nil
	}
}

// WriteToCache writes the given data to the cache.
func (o *OrderUseCase) writeToCache(data interface{}) error {
	switch d := data.(type) {
	case *entity.CexOrder:
		if d.Status == entity.Oucceeded || d.Status == entity.OFailed {
			return nil
		}
		return o.repo.Add(d)
	case *entity.EvmOrder:
		if d.Status == entity.Oucceeded || d.Status == entity.OFailed {
			return nil
		}
		return o.repo.Add(d)
	default:
		return nil
	}
}
