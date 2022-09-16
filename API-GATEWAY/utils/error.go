package utils

import "errors"

var (
	ErrNotFound = errors.New("not found")

	ErrForbidden    = errors.New("forbidden")
	ErrCreateToken  = errors.New("error While Creating Token")
	ErrTokenExpired = errors.New("token is Expired")
	ErrTokenInvalid = errors.New("invalid Token")
	ErrJwtRequired  = errors.New("token Required")
	ErrInvalidToken = errors.New("invalid token")

	ErrBadRequest       = errors.New("bad Request")
	ErrItemOutOfStock   = errors.New("item out of stock")
	ErrItemAlreadyExist = errors.New("item already exist")

	ErrInternalServerError = errors.New("internal Server Error")
)

type Causer interface {
	Cause() error
}

type ErrorWithCode struct {
	code  int
	cause error
}

func (err *ErrorWithCode) Error() string {
	return err.Cause().Error()
}

func (err *ErrorWithCode) Code() int {
	return err.code
}

func (err *ErrorWithCode) Cause() error {
	return err.cause
}
