package httperrors

import (
	"errors"
	"net/http"
)

type HttpError struct {
	Code    int
	Message string
}

func (e *HttpError) Error() string {
	return e.Message
}

func NewHttpError(code int, message string) *HttpError {
	return &HttpError{code, message}
}
func Unauthorized() *HttpError {
	return NewHttpError(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
}
func InvalidCredentials() *HttpError {
	return NewHttpError(http.StatusUnauthorized, "Invalid credentials")
}
func NotFound() *HttpError {
	return NewHttpError(http.StatusNotFound, http.StatusText(http.StatusNotFound))
}
func InternalServerError() *HttpError {
	return NewHttpError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
}

var ErrNotFound = errors.New
