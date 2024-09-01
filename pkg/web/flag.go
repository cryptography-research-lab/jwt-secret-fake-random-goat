package web

import (
	"fmt"
	"github.com/cryptography-research-lab/jwt-secret-fake-random-goat/pkg/config"
	"github.com/cryptography-research-lab/jwt-secret-fake-random-goat/pkg/dal"
	"github.com/cryptography-research-lab/jwt-secret-fake-random-goat/pkg/models"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func flag(c echo.Context) error {

	tokenString := c.QueryParam("token")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return config.JwtSecret, nil
	})
	if err != nil {
		return c.JSON(200, models.Response{
			Code: 403,
			Msg:  err.Error(),
		})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.JSON(200, models.Response{
			Code: 403,
			Msg:  "token error",
		})
	}

	if claims["username"] == dal.FlagUserName {
		return c.JSON(200, models.Response{
			Code: 200,
			Msg:  "ok",
			Data: "flag{YouAreWin}",
		})
	} else {
		return c.JSON(200, models.Response{
			Code: 403,
			Msg:  "who are you? only " + dal.FlagUserName + " can see the flag!",
		})
	}

}
