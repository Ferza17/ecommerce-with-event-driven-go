package utils

import "errors"

var (
	ErrNotFound            = errors.New("not found")
	ErrForbidden           = errors.New("forbidden")
	ErrQueryRead           = errors.New("error while querying storage")
	ErrQueryTxBegin        = errors.New("error while preparing repository transaction")
	ErrQueryTxInsert       = errors.New("error while inserting record")
	ErrQueryTxRollback     = errors.New("error while cancelling record")
	ErrQueryTxUpdate       = errors.New("error while updating record")
	ErrQueryTxCommit       = errors.New("error while committing record changes")
	ErrInvalidStatus       = errors.New("invalid status")
	ErrQueryNoResult       = errors.New("data Not Found")
	ErrCreateToken         = errors.New("error While Creating Token")
	ErrTokenExpired        = errors.New("token is Expired")
	ErrTokenInvalid        = errors.New("invalid Token")
	ErrBuyStock            = errors.New("can't buy more than stock")
	ErrBadRequest          = errors.New("bad Request")
	ErrJwtRequired         = errors.New("token Required")
	ErrUnableRegisterCTX   = errors.New("unable To Register CTX")
	ErrInternalServerError = errors.New("internal Server Error")
	ErrInvalidToken        = errors.New("invalid token")
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
