package dto

type Exchange struct {
	Type string      `json:"type"`
	Conf interface{} `json:"configs"`
}

type GetAllExchangesResponse struct {
	Exchanges []Exchange `json:"exchanges"`
}
