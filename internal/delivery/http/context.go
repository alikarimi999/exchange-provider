package http

import (
	"net/http"
	"time"

	h "exchange-provider/internal/adapter/http"
	"exchange-provider/pkg/errors"

	"github.com/gin-gonic/gin"
)

type ginContext struct {
	ctx   *gin.Context
	admin bool
}

func (ec *ginContext) Param(p string) string {
	return ec.ctx.Param(p)
}
func (ec *ginContext) Bind(i interface{}) error {
	if err := ec.ctx.Bind(i); err != nil {
		return errors.Wrap(errors.Wrap(errors.ErrBadRequest,
			errors.NewMesssage(err.Error())), err)
	}
	return nil
}

func (ec *ginContext) JSON(obj interface{}, err error) {
	var Status int
	if err != nil {
		eCode := errors.ErrorCode(err)
		var Error string

		switch eCode {
		case errors.ErrBadRequest:
			Status = http.StatusBadRequest
		case errors.ErrNotFound:
			Status = http.StatusNotFound
		case errors.ErrForbidden:
			Status = http.StatusForbidden
		default:
			Status = http.StatusInternalServerError
		}

		if ec.admin {
			Error = err.Error()
		} else {
			Error = errors.ErrorMsg(err)

		}
		obj = struct {
			Timestamp string
			Status    int
			Error     string
			Path      string
		}{
			Timestamp: time.Now().UTC().String(),
			Status:    Status,
			Error:     Error,
			Path:      ec.Request().URL.Path,
		}
	} else {
		Status = http.StatusOK
	}
	ec.ctx.JSON(Status, obj)
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
func newContext(ctx *gin.Context, admin bool) h.Context {
	return &ginContext{ctx: ctx, admin: admin}
}
