package errors

import (
	"bytes"
	"fmt"
)

type ErrCode string

const (
	ErrNotFound         = ErrCode("not found")
	ErrForbidden        = ErrCode("forbidden")
	ErrConflict         = ErrCode("conflict")
	ErrConnectionFailed = ErrCode("connection failed")
	ErrInternal         = ErrCode("internal")
	ErrBadRequest       = ErrCode("bad request")
	ErrUnknown          = ErrCode("unknown error")
)

type Op string // Operation name
type Msg interface {
	String() string
}

type Error struct {
	code    ErrCode
	message Msg

	op              Op
	err             error
	relationContext string
}

func Wrap(args ...interface{}) error {
	if len(args) == 0 {
		return nil
	}

	e := &Error{}
	for _, arg := range args {
		switch arg := arg.(type) {
		case ErrCode:
			e.code = arg
		case Msg:
			e.message = arg
		case Op:
			e.op = arg
		case error:
			e.err = arg
		case string:
			e.relationContext = arg
		default:
			continue
		}
	}
	return e
}

func ErrorCode(err error) ErrCode {
	if err == nil {
		return ""
	} else if e, ok := err.(*Error); ok && e.code != "" {
		return e.code
	} else if ok && e.err != nil {
		return ErrorCode(e.err)
	}

	return ErrUnknown
}

func ErrorMsg(err error) string {
	if err == nil {
		return ""
	} else if e, ok := err.(*Error); ok && e.message != nil {
		return e.message.String()
	} else if ok && e.err != nil {
		return ErrorMsg(e.err)
	}
	return "An internal error has occurred."
}

func (e *Error) Error() string {
	var buf bytes.Buffer

	// Print the current operation in our stack, if any.
	if e.op != "" {
		fmt.Fprintf(&buf, "%s: ", e.op)
	}

	// If wrapping an error, print its Error() message.
	// Otherwise print the error code & message.
	if e.err != nil {
		buf.WriteString(e.err.Error())
	} else {
		if e.code != "" {
			fmt.Fprintf(&buf, "<%s> ", e.code)
		}
		if e.message != nil {
			buf.WriteString(e.message.String())
		}
	}

	if e.relationContext != "" {
		buf.WriteString(fmt.Sprintf(" related to (%s) ", e.relationContext))
	}

	return buf.String()
}

func ErrorOp(e error) Op {
	if e == nil {
		return ""
	} else if e, ok := e.(*Error); ok && e.op != "" {
		return e.op
	} else if ok && e.err != nil {
		return ErrorOp(e.err)
	}
	return ""
}
