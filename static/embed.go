package static

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed *
var content embed.FS

func GetStaticFS() http.FileSystem {
	// 返回嵌入的文件系统，排除 .go 文件
	fsys, err := fs.Sub(content, ".")
	if err != nil {
		panic(err)
	}
	return http.FS(fsys)
}
