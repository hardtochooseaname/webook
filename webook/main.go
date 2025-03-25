package main

import (
	"xws/webook/internal/web"
)

func main() {
	server := web.RegisterRoutes()
	server.Run(":8080")
}
