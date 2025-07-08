package errors

type ErrorCode string

var (
	ErrCodeNotFound      ErrorCode = "NOT_FOUND"
	ErrCodeBadRequest    ErrorCode = "BAD_REQUEST"
	ErrCodeInternalError ErrorCode = "INTERNAL_ERROR"
	ErrCodeUnauthorized  ErrorCode = "UNAUTHORIZED"
	ErrCodeForbidden     ErrorCode = "FORBIDDEN"
)
