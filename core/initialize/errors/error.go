// Package errors provides a way to return detailed information
// for an RPC request error. The error is normally JSON encoded.
package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/gommon/log"
)

// Error implements the error interface.
type Error struct {
	ID      int    `json:"id"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Title   string `json:"-"`
}

func (e *Error) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func (e *Error) StatusCode() int {
	return int(e.ID)
}

func (e *Error) ErrorCode() string {
	return e.Code
}

// New generates a custom error.
func New(code, title, detail string, id int) error {
	return &Error{
		ID:      id,
		Code:    code,
		Message: detail,
		Title:   title,
	}
}

// Parse tries to parse a JSON string into an error. If that
// fails, it will set the given string as the error detail.
func Parse(err string) *Error {
	e := new(Error)
	if marshaledErr := json.Unmarshal([]byte(err), e); marshaledErr != nil {
		e.Message = err
	}
	return e
}

// BadRequest generates a 400 error.
func BadRequest(a ...interface{}) error {
	return &Error{
		Code:    ErrCommonBadRequest,
		ID:      http.StatusBadRequest,
		Message: formatDetail(a...),
		Title:   http.StatusText(http.StatusBadRequest),
	}
}

// Unauthorized generates a 401 error.
func Unauthorized(a ...interface{}) error {
	return &Error{
		Code:    ErrCommonUnauthorized,
		ID:      http.StatusUnauthorized,
		Message: formatDetail(a...),
		Title:   http.StatusText(http.StatusUnauthorized),
	}
}

// Forbidden generates a 403 error.
func Forbidden(a ...interface{}) error {
	return &Error{
		Code:    ErrCommonForbidden,
		ID:      http.StatusForbidden,
		Message: formatDetail(a...),
		Title:   http.StatusText(http.StatusForbidden),
	}
}

// NotFound generates a 404 error.
func NotFound(a ...interface{}) error {
	return &Error{
		Code:    ErrCommonNotFound,
		ID:      http.StatusNotFound,
		Message: formatDetail(a...),
		Title:   http.StatusText(http.StatusNotFound),
	}
}

// MethodNotAllowed generates a 405 error.
func MethodNotAllowed(a ...interface{}) error {
	return &Error{
		Code:    ErrCommonMethodNotAllowed,
		ID:      http.StatusMethodNotAllowed,
		Message: formatDetail(a...),
		Title:   http.StatusText(http.StatusMethodNotAllowed),
	}
}

// Timeout generates a 408 error.
func Timeout(a ...interface{}) error {
	return &Error{
		Code:    ErrCommonTimeout,
		ID:      http.StatusRequestTimeout,
		Message: formatDetail(a...),
		Title:   http.StatusText(http.StatusRequestTimeout),
	}
}

// Conflict generates a 409 error.
func Conflict(a ...interface{}) error {
	return &Error{
		Code:    ErrCommonConflict,
		ID:      http.StatusConflict,
		Message: formatDetail(a...),
		Title:   http.StatusText(http.StatusConflict),
	}
}

// InternalServerError generates a 500 error.
func InternalServerError(a ...interface{}) error {
	return &Error{
		Code:    ErrCommonInternalServerError,
		ID:      http.StatusInternalServerError,
		Message: formatDetail(a...),
		Title:   http.StatusText(http.StatusInternalServerError),
	}
}

// ServiceError generates a custom error of system.
func ServiceError(errCode string) error {
	return &Error{
		Code: errCode,
		ID:   http.StatusOK,
	}
}

func ErrorCode(err error) string {
	var targetErr *Error
	if errors.As(err, &targetErr) {
		return targetErr.Code
	}
	return ""
}

func formatDetail(a ...interface{}) string {
	if len(a) == 0 {
		return ""
	}
	return fmt.Sprintf("%s", a...)
}

func ValidateError(err error) []error {
	var listErr []error
	for _, err := range err.(validator.ValidationErrors) {
		log.Info(err)
		listErr = append(listErr, &Error{
			Code:    err.Field() + "Invalid",
			ID:      http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return listErr
}
