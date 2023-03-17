package dto

type Account struct {
	Conf interface{} `json:"configs"`
}

type GetAllExchangesResponse struct {
	Exchanges map[uint]*Account `json:"exchanges,omitempty"`
	Msgs      []string          `json:"messages"`
}
