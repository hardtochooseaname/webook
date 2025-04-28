package web

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
    emailExp *regexp.Regexp
    passwordExp *regexp.Regexp
}

func NewUserHandler() *UserHandler {
    const (
        emailRegexPattern = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
        passwordRegexPattern = `^[A-Za-z\d@$!%*?&]{8,}$`
    )

    // 编译正则表达式
    emailRegex := regexp.MustCompile(emailRegexPattern)
    passwordRegex := regexp.MustCompile(passwordRegexPattern)



    return &UserHandler{
        emailExp: emailRegex,
        passwordExp: passwordRegex, 
    }
}


// 登陆页面的处理逻辑：
// 1. 接受请求并校验请求（不同的业务不同的数据，校验要求不同，例如注册需要校验邮箱格式、两次密码是否一致等）
// 2. 调用业务逻辑处理请求
// 3. 根据业务处理结果返回响应
func (uh *UserHandler) SignUp(ctx *gin.Context) {
    //-------------------------------------
    // 接收用户请求和数据
    //-------------------------------------
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


    //-------------------------------------
    // 用户数据校验
    //-------------------------------------

    // 1. 邮箱格式校验
    if !uh.emailExp.MatchString(req.Email) {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "邮箱格式不正确"})
        return
    }

    // 2. 两次密码是否一致
    if req.Password != req.ConfirmPassword {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "两次输入的密码不一致"})
        return
    }

    // 3. 密码格式校验
    if !uh.passwordExp.MatchString(req.Password) {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "密码格式不符合要求（至少8位，仅允许字母、数字和 @$!%*?&）",
        })
        return
    }

    //-------------------------------------
    // 业务逻辑处理 
    //-------------------------------------



    // 最终返回成功
    ctx.JSON(http.StatusOK, gin.H{
        "message": "Hello user " + req.Email + "! You have signed up successfully!",
    })
}
    

func (uh *UserHandler) LogIn(ctx *gin.Context) {

}

func (uh *UserHandler) Edit(ctx *gin.Context) {

}

func (uh *UserHandler) Profile(ctx *gin.Context) {

}

