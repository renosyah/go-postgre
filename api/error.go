package api

import (
	"fmt"
)

type (
	Error struct {
		Log        string `json:"log"`
		StatusCode int    `json:"status_code"`
		Message    string `json:"message"`
	}
)

func (e *Error) Error() string {
	return e.Message
}

func NewError(err error, message string, status int) *Error {
	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}

	return &Error{
		Log:        errMsg,
		Message:    message,
		StatusCode: status,
	}
}

func NewErrorWrap(err error, prefix, suffix, message string, status int) *Error {
	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}

	return &Error{
		Log:        fmt.Sprintf("%s/%s : %s", prefix, suffix, errMsg),
		Message:    message,
		StatusCode: status,
	}
}
