package infrastructure

import "github.com/labstack/echo/v4"

func httpErrorHandler(err error, c echo.Context) {
	status := getErrorStatus(err)
	body := map[string]any{
		"status":  status,
		"error": err.Error(),
	}

	if err := c.JSON(status, body); err != nil {
		c.Logger().Error(err)
	}
}

func getErrorStatus(err error) int {
	if e, ok := err.(*echo.HTTPError); ok {
		return e.Code
	}

	return 200
}
