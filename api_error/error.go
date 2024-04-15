package api_error

import (
	"github.com/labstack/echo/v4"
)

// New function is a facade to handle error messages and status codes
func New(status int, message string) *echo.HTTPError {
	return &echo.HTTPError{
		Code:     status,
		Message:  message,
		Internal: nil,
	}
}
