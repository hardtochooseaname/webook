package main

import (
	"xws/webook/internal/repository"
	"xws/webook/internal/repository/dao"
	"xws/webook/internal/service"
	"xws/webook/internal/web"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 1. 初始化数据库（思考为什么放在最前面？）
	db := initDB()

	// 2. 初始化服务器
	// 后面如果还有添加中间件等复杂操作可以提取出来成 initWebServer 函数
	server := gin.Default()

	// 3. 初始化具体业务
	u := initUser(db)

	// 4. 注册路由
	u.RegisterRoutes(server)

	// 5. 启动服务器
	server.Run(":8080")
}

func initUser(db *gorm.DB) *web.UserHandler {
	ud := dao.NewUserDAO(db)
	r := repository.NewRepository(ud)
	svc := service.NewUserService(r)
	u := web.NewUserHandler(svc)
	return u
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13306)/webook"))
	if err != nil {
		panic("failed to connect database")
	}
	if err := dao.InitTable(db); err != nil {
		panic("failed to init table")
	}
	return db
}


