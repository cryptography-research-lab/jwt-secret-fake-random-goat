package web

import (
	"github.com/cryptography-research-lab/jwt-secret-fake-random-goat/pkg/dal"
	"github.com/cryptography-research-lab/jwt-secret-fake-random-goat/pkg/models"
	"github.com/labstack/echo/v4"
)

// 查看平台上当前注册的所有用户
func userList(c echo.Context) error {
	usernames := dal.ListUsername()
	return c.JSON(200, models.Response{
		Code: 200,
		Msg:  "ok",
		Data: usernames,
	})
}
