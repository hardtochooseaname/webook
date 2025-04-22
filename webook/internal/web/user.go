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
