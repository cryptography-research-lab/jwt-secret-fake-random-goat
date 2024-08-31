package web

import (
	"github.com/labstack/echo/v4"
	"strconv"
	"time"
)

func responseHeaderServerTime() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Add("x-server-time", strconv.FormatInt(time.Now().Unix(), 10))
			return nil
		}
	}
}
