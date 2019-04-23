package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	// 全局中间件
	// Logger 中间件将日志写入 gin.DefaultWriter，即使你将 GIN_MODE 设置为 release。
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery 中间件会 recover 任何 panic。如果有 panic 的话，会写入 500。
	r.Use(gin.Recovery())

	// 你可以为每个路由添加任意数量的中间件。
	// r.GET("/benchmark", MyBenchLogger(), benchEndpoint)

	// 认证路由组
	// authorized := r.Group("/", AuthRequired())
	// 和使用以下两行代码的效果完全一样:
	authorized := r.Group("/")
	// 路由组中间件! 在此例中，我们在 "authorized" 路由组中使用自定义创建的
	// AuthRequired() 中间件
	authorized.Use(AuthRequired())
	{
		authorized.GET("/login", func(c *gin.Context) {
			c.String(http.StatusOK, "login")
		})
		authorized.GET("/submit", func(c *gin.Context) {
			c.String(http.StatusOK, "submit")
		})
		authorized.GET("/read", func(c *gin.Context) {
			c.String(http.StatusOK, "read")
		})

		// 嵌套路由组
		testing := authorized.Group("testing")
		testing.GET("/analytics", func(c *gin.Context) {
			c.String(http.StatusOK, "testing")
		})
	}

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8090")
}
