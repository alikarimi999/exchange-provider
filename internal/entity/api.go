package entity

type APIToken struct {
	Id      string `bson:"_id"`
	BusName string
	BusId   uint
	Level   uint
	Ips     []string
	Write   bool
	CheckIp bool
}

type ApiService interface {
	ApiPrefix() string
	AddApiToken(*APIToken) error
	Get(id string) (*APIToken, error)
	GetByBusId(id uint) ([]*APIToken, error)
	Update(api *APIToken) error
	Remove(id string) error
	MaxIps() uint
}
