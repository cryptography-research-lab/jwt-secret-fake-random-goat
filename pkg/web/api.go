package web

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/cryptography-research-lab/jwt-secret-fake-random-goat/static"
	"strconv"
)

func Run(port uint) {

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:       "/",
		Browse:     true,
		HTML5:      true,
		Filesystem: static.GetStaticFS(),
	}))
	//e.Use(responseHeaderServerTime())

	// Routes
	e.GET("/", homePage)
	e.GET("/register", register)
	e.GET("/login", login)
	e.GET("/user_list", userList)
	e.GET("/server_info", serverInfo)
	e.GET("/flag", flag)

	// Start server
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(int(port))))

}

// 主页
func homePage(c echo.Context) error {
	return c.Redirect(301, "/index.html")
}