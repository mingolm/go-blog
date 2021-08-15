package errutil

import (
	"errors"
	"fmt"
)

type Error struct {
	err     error
	message string
	base    bool
}

func (e *Error) Unwrap() error {
	return e.err
}

func (e *Error) Error() string {
	if e.message == "" {
		return e.err.Error()
	}
	return fmt.Sprintf("%s: %s", e.err, e.message)
}

func (e *Error) Msg(message string) *Error {
	if e.base {
		return &Error{
			err:     e.err,
			message: message,
			base:    true,
		}
	}
	return &Error{
		err:     e,
		message: message,
		base:    true,
	}
}

func NewBase(err error) *Error {
	return &Error{
		err:  err,
		base: false,
	}
}

var (
	ErrAlreadyExists      = NewBase(errors.New("already exists"))
	ErrNotFound           = NewBase(errors.New("not found"))
	ErrInternal           = NewBase(errors.New("internal error"))
	ErrInvalidArguments   = NewBase(errors.New("invalid arguments"))
	ErrFailedPrecondition = NewBase(errors.New("failed precondition"))
	ErrResourceExhausted  = NewBase(errors.New("resource exhausted"))
	ErrPermissionDenied   = NewBase(errors.New("permission denied"))
	ErrUnauthenticated    = NewBase(errors.New("unauthenticated"))
	ErrUnimplemented      = NewBase(errors.New("unimplemented"))
	ErrMethodNotAllowed   = NewBase(errors.New("method not allowed"))
	ErrPageNotFound       = NewBase(errors.New("page not found"))
)
