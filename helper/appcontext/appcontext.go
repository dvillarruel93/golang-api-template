package appcontext

import (
	"context"
	"github.com/labstack/echo/v4"
)

func EchoContextToContext(c echo.Context, additionalVals ...string) context.Context {
	ctx := c.Request().Context()

	var eCtxVals []string
	eCtxVals = append(eCtxVals, additionalVals...)
	for _, key := range eCtxVals {
		value := c.Get(key)
		if value != nil {
			ctx = context.WithValue(ctx, key, value)
		}
	}

	return ctx
}
