package http

import (
	"net/http"

	h "order_service/internal/adapter/http"

	"github.com/gin-gonic/gin"
)

type ginContext struct {
	ctx *gin.Context
}

func (ec *ginContext) Param(p string) string {
	return ec.ctx.Param(p)
}
func (ec *ginContext) Bind(i interface{}) error {
	return ec.ctx.Bind(i)
}
func (ec *ginContext) JSON(code int, obj interface{}) {
	ec.ctx.JSON(code, obj)
}
func (ec *ginContext) Request() *http.Request {
	return ec.ctx.Request
}
func (ec *ginContext) GetKey(key string) (value interface{}, exists bool) {
	return ec.ctx.Get(key)
}

func (ec *ginContext) SetKey(key string, value interface{}) {
	ec.ctx.Set(key, value)
}

func (ec *ginContext) GetHeader(key string) string {
	return ec.ctx.GetHeader(key)
}

// newContext function return a new ServerContext
func newContext(ctx *gin.Context) h.Context {
	return &ginContext{ctx: ctx}
}
