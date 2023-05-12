package http

import (
	"exchange-provider/internal/entity"
	"net/http"
)

type Context interface {
	Param(string) string
	Bind(interface{}) error
	JSON(obj interface{}, err error)
	SetApi(*entity.APIToken)
	GetApi() (api *entity.APIToken)
	Request() *http.Request
	GetHeader(string) string
}
