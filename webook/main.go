package main

import (
	"xws/webook/internal/web"
)

func main() {
	server := web.RegisterRoutes()

	// TODO：curl测试signup逻辑、添加注释，测试中间件的User函数

	server.Run(":8080")
}
