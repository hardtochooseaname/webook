package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}


// 登陆页面的处理逻辑：
// 1. 接受请求并校验请求（不同的业务不同的数据，校验要求不同，例如注册需要校验邮箱格式、两次密码是否一致等）
// 2. 调用业务逻辑处理请求
// 3. 根据业务处理结果返回响应
func (uh *UserHandler) SignUp(ctx *gin.Context) {
	// 定义私有结构体
	// 只希望结构体在函数内部被使用
	// json格式前后端数据传输常用格式
	type SignUpReq struct {
		Email string `json:"email"`
		Password string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}

	req := SignUpReq{}

	// Bind方法根据响应的 Content-Type（json或xml）
	// 将请求体的数据（按json或xml格式）解析并拷贝到 req 结构体中
	// 字段的解析按照提前设定好的结构体 schema（要求两处字段同名）
	if err := ctx.Bind(&req); err != nil {
		return
	}


	// TODO: 添加请求校验，正则表达式判断邮箱格式是否正确、两次密码是否一致、密码复杂度是否符合要求
	
	ctx.String(http.StatusOK, "Hello user %s! You have signed up successfully!", req.Email)
}

func (uh *UserHandler) LogIn(ctx *gin.Context) {

}

func (uh *UserHandler) Edit(ctx *gin.Context) {

}

func (uh *UserHandler) Profile(ctx *gin.Context) {

}

// 简单请求测试CORS
func (uh *UserHandler) ProfileGET(ctx *gin.Context) {
    user := struct {
        ID    int    `json:"id"`
        Email string `json:"email"`
        Name  string `json:"name"`
    }{
        ID:    1,
        Email: "test@example.com",
        Name:  "Go 学习者",
    }

    //  第三步：以 JSON 格式返回给前端
    ctx.JSON(http.StatusOK, gin.H{
        "status": "success",
        "data":   user,
    })
}

// 测非简单请求
func (uh *UserHandler) ProfilePOST(ctx *gin.Context) {
    // 定义你的请求结构
    var req struct {
        Ping string `json:"ping"`
    }

    // 解析 JSON Body
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // 这里你可以根据 req.Ping 做逻辑，比如与数据库交互、校验权限等
    // 这里只是简单地把它原样返回
    ctx.JSON(http.StatusOK, gin.H{
        "status": "pong",
        "echo":   req.Ping,
    })
}