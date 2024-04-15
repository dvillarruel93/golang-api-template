package middleware

import (
	"empty-api-struct/api_error"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

const validToken = "valid-token"

var errInvalidToken = errors.New("invalid token")

func AuthMW() echo.MiddlewareFunc {
	return AuthenticationHandlerFunc
}

func AuthenticationHandlerFunc(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authToken := c.Request().Header.Get("auth-token")
		if authToken != validToken {
			return api_error.New(http.StatusUnauthorized, "invalid token")
			//return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
		}

		return next(c)

	}
}
