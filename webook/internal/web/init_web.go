// web包主要进行路由的集中注册，以及存放业务逻辑handler的具体代码
// --note--// 另一种选择：分散注册，把路由的注册分散在每个 Handler 那里
// --note--// 例如 User 相关的注册就与 UserHandler 的定义放在一起
package web

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes() *gin.Engine {
	server := gin.Default()

	// --note--// 可以对所有的路由进行分组，与 User 相关的路由就放在 UserRegister 中
	// --note--// 也可以把所有的路由注册全都一股脑放进 RegisterRoutes 中
	server = RegisterUserRoutes(server)
	return server
}

func RegisterUserRoutes(server *gin.Engine) *gin.Engine {
	u := NewUserHandler()

	// 路由分组
	//--note--// 这样每个同组的路由注册时，路径中不用重复写相同的前缀，如这里的 /user
	ug := server.Group("/user")
	// 注册
	ug.POST("/signup", u.SignUp)
	// 登陆
	ug.POST("/login", u.LogIn)
	// 查询
	ug.GET("/sprofile", u.Profile)
	// 修改
	ug.POST("/edit", u.Edit)

	return server
}
