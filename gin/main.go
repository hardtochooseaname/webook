package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func main() {
  // 服务器引擎，服务器的路由、请求响应等东西都在这个引擎上面处理
  r := gin.Default() 

  // 注册路由（GET, POST, DELETE, PUT方法的路由分别有对应method来注册）
  // 1. 静态路由
  r.GET("/hello", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "Hello, Go!",
    })
  })

  // 2. 参数路由
  r.GET("/hello/:name", func(c *gin.Context) {
    c.String(http.StatusOK, "Hello, %s", c.Param("name"))
  })

  // 3. 通配符路由
  r.GET("/page/*.html", func(c *gin.Context) {
    c.String(http.StatusOK, "This is page: %s.html", c.Param(".html"))
  })

  // 查询参数 "/whether?city=chengdu"
  r.GET("/whether", func(c *gin.Context) {
    c.String(http.StatusOK, "Looking up whether in %s", c.Query("city")) 
  })


  r.Run(":8088") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}