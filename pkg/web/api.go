package web

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"strconv"
)

func Run(port int) {

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "static",
		Browse: true,
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
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(port)))

}

// 主页
func homePage(c echo.Context) error {
	return c.Redirect(301, "static/index.html")
}
