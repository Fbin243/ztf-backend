package errors

import (
	"errors"
	"fmt"
)

type Error struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}

func NewError(code ErrorCode, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("Code: %s, Message: %s", e.Code, e.Message)
}

var (
	ErrorNotFound       = NewError(ErrCodeNotFound, "resource not found")
	ErrorBadRequest     = NewError(ErrCodeBadRequest, "bad request")
	ErrorInternal       = NewError(ErrCodeInternalError, "internal server error")
	ErrorUnauthorized   = NewError(ErrCodeUnauthorized, "unauthorized access")
	ErrorNoRowsAffected = errors.New("no rows affected")
)
