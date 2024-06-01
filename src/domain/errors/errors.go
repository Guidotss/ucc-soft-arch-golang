package errors

import (
	"fmt"
)

type Error struct {
	Code           string
	Message        string
	HTTPStatusCode int
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func NewError(code, message string, httpStatusCode int) *Error {
	return &Error{
		Code:           code,
		Message:        message,
		HTTPStatusCode: httpStatusCode,
	}
}
