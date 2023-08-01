package entity

type ExchangeStore interface {
	Get(id uint) (Exchange, error)
	GetByNID(name string) (Exchange, error)
	Exists(id uint) bool
	AddExchange(ex Exchange) error
	GetAll() []Exchange
	GetAllMap() map[string]Exchange
	EnableDisable(exId uint, enable bool) error
	EnableDisableAll(enable bool) error
	Remove(id uint) error
	RemoveAll() error
}
