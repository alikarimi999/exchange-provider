package http

import (
	"net/http"
	"exchange-provider/pkg/errors"
)

func handlerErr(ctx Context, err error) {
	var msg string
	if errors.ErrorMsg(err) != errors.InternalErrMsg {
		msg = errors.ErrorMsg(err)
	} else {
		msg = err.Error()
	}
	switch errors.ErrorCode(err) {
	case errors.ErrNotFound:
		ctx.JSON(http.StatusNotFound, msg)
		return
	case errors.ErrBadRequest:
		ctx.JSON(http.StatusBadRequest, msg)
		return
	case errors.ErrForbidden:
		ctx.JSON(http.StatusForbidden, msg)
		return
	default:
		ctx.JSON(http.StatusInternalServerError, msg)
		return
	}
}
