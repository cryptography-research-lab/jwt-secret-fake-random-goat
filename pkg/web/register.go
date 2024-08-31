package web

import (
	"github.com/cryptography-research-lab/jwt-secret-fake-random-goat/pkg/dal"
	"github.com/cryptography-research-lab/jwt-secret-fake-random-goat/pkg/models"
	"github.com/labstack/echo/v4"
)

// 注册用户
func register(c echo.Context) error {

	username := c.QueryParam("username")
	if username == "" {
		return c.JSON(200, models.Response{
			Code: 400,
			Msg:  "username is empty",
		})
	}

	passwd := c.QueryParam("passwd")
	if passwd == "" {
		return c.JSON(200, models.Response{
			Code: 400,
			Msg:  "passwd is empty",
		})
	}

	err := dal.Register(username, passwd)
	if err != nil {
		return c.JSON(200, models.Response{
			Code: 500,
			Msg:  err.Error(),
		})
	}

	return c.JSON(200, models.Response{
		Code: 200,
		Msg:  "register success",
	})

}
