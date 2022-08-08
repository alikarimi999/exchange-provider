package http

import "net/http"

type Context interface {
	Param(string) string
	Bind(interface{}) error
	JSON(int, interface{})
	GetKey(key string) (interface{}, bool)
	SetKey(key string, value interface{})
	Request() *http.Request
	GetHeader(string) string
}
