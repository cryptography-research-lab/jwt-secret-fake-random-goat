package web

import (
	"github.com/cryptography-research-lab/jwt-secret-fake-random-goat/pkg/config"
	"github.com/cryptography-research-lab/jwt-secret-fake-random-goat/pkg/dal"
	"github.com/cryptography-research-lab/jwt-secret-fake-random-goat/pkg/models"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"time"
)

// 登录，登录成功之后会签发jwt token
func login(c echo.Context) error {

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

	findPasswd, exists := dal.FindPasswd(username)
	if !exists || findPasswd != passwd {
		return c.JSON(200, models.Response{
			Code: 403,
			Msg:  "username or passwd is wrong",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"nbf":      time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	tokenString, err := token.SignedString(config.JwtSecret)
	if err != nil {
		return c.JSON(200, models.Response{
			Code: 500,
			Msg:  err.Error(),
		})
	}

	return c.JSON(200, models.Response{
		Code: 200,
		Msg:  "login success",
		Data: tokenString,
	})
}
