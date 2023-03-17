package swapspace

import (
	"exchange-provider/internal/entity"
	"strings"
)

func (ex *exchange) Tokens() []*entity.Token {
	ts := []*entity.Token{}
	ex.tokens.RLock()
	for _, t := range ex.tokens.list {
		ts = append(ts, t.toEntity())
	}
	ex.tokens.RUnlock()
	return ts
}

type token struct {
	Code       string
	Network    string
	HasExtraId bool `json:"hasExtraId"`
}

func (t *token) toEntity() *entity.Token {
	token := &entity.Token{
		TokenId:    strings.ToUpper(t.Code),
		ChainId:    strings.ToUpper(t.Network),
		HasExtraId: t.HasExtraId,
	}

	return token
}

func fromEntity(t *entity.Token) *token {
	token := &token{
		Code:       strings.ToLower(t.TokenId),
		Network:    strings.ToLower(t.ChainId),
		HasExtraId: t.HasExtraId,
	}

	return token
}
