package infrastructure

import "github.com/labstack/echo/v4"

type CustomContext struct {
	echo.Context
}

func (cc *CustomContext) CustomContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &CustomContext{Context: c}
		return next(cc)
	}
}
