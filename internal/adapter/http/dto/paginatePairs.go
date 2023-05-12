package dto

type PaginatedPairsResp struct {
	PaginatedResponse
	Pairs interface{} `json:"pairs"`
}
