package dto

type Account struct {
	Conf interface{} `json:"configs"`
}

type GetAllExchangesResponse struct {
	Exchanges map[string]*Account `json:"exchanges,omitempty"`
	Msgs      []string            `json:"messages"`
}
