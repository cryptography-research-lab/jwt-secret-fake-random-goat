package web

import (
	"github.com/cryptography-research-lab/jwt-secret-fake-random-goat/pkg/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"runtime"
	"time"
)

// 服务启动时间
var startTime = time.Now()

type ServerInfo struct {

	// 服务器的启动时间
	BootstrapTime time.Time `json:"bootstrap_time"`

	// 执行程序或者编译时使用的golang的sdk的版本
	// 从Go 1.20版本开始，Go的随机数生成器在启动时会自动使用当前时间的纳秒值作为种子进行初始化
	// 而在此之前，math/rand的默认的seed是1
	GolangSdkVersion string `json:"golang_sdk_version"`
}

// 查看服务启动信息
func serverInfo(c echo.Context) error {
	return c.JSON(http.StatusOK, &models.Response{
		Code: 200,
		Msg:  "ok",
		Data: &ServerInfo{
			BootstrapTime:    startTime,
			GolangSdkVersion: runtime.Version(),
		},
	})
}
