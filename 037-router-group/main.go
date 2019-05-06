package main

import (
	"github.com/gin-gonic/gin"
)

func loginEndpointV1(c *gin.Context) {
	c.String(200, "/v1/login/")
}

func submitEndpointV1(c *gin.Context) {
	c.String(200, "/v1/submit/")
}

func readEndpointV1(c *gin.Context) {
	c.String(200, "/v1/read/")
}

func loginEndpointV2(c *gin.Context) {
	c.String(200, "/v2/login/")
}

func submitEndpointV2(c *gin.Context) {
	c.String(200, "/v2/submit/")
}

func readEndpointV2(c *gin.Context) {
	c.String(200, "/v2/read/")
}

func main() {
	router := gin.Default()

	// 简单的路由组: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", loginEndpointV1)
		v1.POST("/submit", submitEndpointV1)
		v1.POST("/read", readEndpointV1)
	}

	// 简单的路由组: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", loginEndpointV2)
		v2.POST("/submit", submitEndpointV2)
		v2.POST("/read", readEndpointV2)
	}

	router.Run(":8090")
}
