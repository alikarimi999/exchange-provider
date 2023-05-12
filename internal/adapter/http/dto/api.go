package dto

type CreateApiReq struct {
	Key     string   `json:"apiKey"`
	BusName string   `json:"busName"`
	BusId   uint     `json:"busId"`
	Level   uint     `json:"level"`
	Ips     []string `json:"ips"`
	Write   bool     `json:"write"`
	CheckIp bool     `json:"checkIp"`
}
