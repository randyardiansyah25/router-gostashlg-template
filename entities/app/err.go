package app

import (
	"errors"
	"fmt"
)

var (
	ErrNoRecord                = errors.New("no records found")
	ErrInternalServiceError    = errors.New("internal service error")
	ErrDuplicateEntry          = errors.New("duplicate entry")
	ErrInvalidInput            = errors.New("invalid input")
	ErrSystem                  = errors.New("system error")
	ErrPaymentNotAvailable     = errors.New("payment method not available")
	ErrAlreadyPaid             = errors.New("already paid")
	ErrInsufficientStock       = errors.New("insufficient stock")
	ErrInvalidProductVariantId = errors.New("invalid product variant id")
	ErrNotAllowed              = errors.New("request not allowed")
)

func NewCustomError(args ...interface{}) error {
	er := CustomError{}
	er.SetError(args...)
	return &er
}

func NewCustomErrorf(format string, args ...interface{}) error {
	er := CustomError{}
	er.SetErrorf(format, args...)
	return &er
}

type CustomError struct {
	Err error
}

func (oe *CustomError) Error() string {
	return oe.Err.Error()
}

func (oe *CustomError) SetErrorf(format string, args ...interface{}) {
	oe.SetError(fmt.Sprintf(format, args...))
}

func (oe *CustomError) SetError(args ...interface{}) {
	oe.Err = errors.New(fmt.Sprint(args...))
}
