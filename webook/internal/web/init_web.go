// web包主要进行路由的集中注册，以及存放业务逻辑handler的具体代码
// --note--// 另一种选择：分散注册，把路由的注册分散在每个 Handler 那里
// --note--// 例如 User 相关的注册就与 UserHandler 的定义放在一起
package web

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes() *gin.Engine {
	server := gin.Default()

	// 测试中间件逻辑
	//   中间件其实就是下面这样注册的函数
	//   用户请求到达路由之前，先全部经过一次中间件，所以中间件适合用来处理一些所有业务都关心的东西
	//   中间件执行顺序与注册顺序保持一致
	//   中间件其实属于AOP，面向切面编程
	// 中间件与路由 handler
	//   两者的函数其实都是同一个签名，是 type HandlerFunc func(*Context)
	//   区别简单说其实就是中间件作用于所有的路由，而handler只作用于一个对应的路由
	// 中间件注册位置
	//   中间件的注册需要在路由注册之前
	//   注册在一个server := gin.Default()上的中间件只会作用于这个server上的路由
	//   Gin 会把中间件和路由一起做链式处理，如果你先注册路由，后注册中间件，中间件不会被插入到已有的路由链中
	server.Use(func(ctx *gin.Context) {
		fmt.Println("第一个middleware测试")
	})

	server.Use(func(ctx *gin.Context) {
		fmt.Println("第22222个middleware测试")
	})

	// CORS中间件测试
	server.Use(cors.New(cors.Config{
		// 下面两行配置服务器允许跨域请求所在的源
		//   第一行直接列出允许的源
		//   第二行通过函数来判断 origin 源是否是允许的源
		// 修改下面两行配置即可复现跨域问题
		AllowOrigins: []string{"http://localhost:3000", "http://192.168.20.243:3000"},
		AllowOriginFunc: func(origin string) bool {
			return strings.HasPrefix(origin, "http://192.168.20.")
		},

		AllowMethods:  []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		// 预检请求的 access-control-request-headers 字段中说明了跨域请求的需要用到的请求头部字段，需要在这里全部allow
		AllowHeaders:  []string{"Origin", "Content-Type", "Authorization", "x-debug"},
		
		// 可以安全暴露给 API 的头部
		ExposeHeaders: []string{"Content-Length"},
		
		// 这个就是 cookie 之类的凭证信息
		AllowCredentials: true,
		
		// 这个指定的是preflight请求有效的时常（当前设置时12个小时之内进行过预检，再发起跨域请求时就不用再预检了）
		MaxAge:           12 * time.Hour,
	}))

	// --note--// 可以对所有的路由进行分组，与 User 相关的路由就放在 UserRegister 中
	// --note--// 也可以把所有的路由注册全都一股脑放进 RegisterRoutes 中
	server = RegisterUserRoutes(server)
	return server
}

func RegisterUserRoutes(server *gin.Engine) *gin.Engine {
	u := &UserHandler{}

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
	// 供 CORS 测试的路由
	ug.GET("/profile", u.ProfileGET)
	ug.POST("/profile", u.ProfilePOST)


	return server
}
