package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type formA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}

type formB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}

func SomeHandlerA(c *gin.Context) {
	objA := formA{}
	objB := formB{}
	// c.ShouldBind 使用了 c.Request.Body，不可重用。
	if errA := c.ShouldBind(&objA); errA == nil {
		c.String(http.StatusOK, `the body should be formA`)
		// 因为现在 c.Request.Body 是 EOF，所以这里会报错。
	} else if errB := c.ShouldBind(&objB); errB == nil {
		c.String(http.StatusOK, `the body should be formB`)
	}
}

func SomeHandlerB(c *gin.Context) {
	objA := formA{}
	objB := formB{}
	// 读取 c.Request.Body 并将结果存入上下文。
	if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
		c.String(http.StatusOK, `the body should be formA`)
		// 这时, 复用存储在上下文中的 body。
	} else if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB == nil {
		c.String(http.StatusOK, `the body should be formB JSON`)
		// 可以接受其他格式
	} else if errB2 := c.ShouldBindBodyWith(&objB, binding.XML); errB2 == nil {
		c.String(http.StatusOK, `the body should be formB XML`)
	}
}

func main() {
	r := gin.Default()

	r.GET("/ShouldBind", SomeHandlerA)

	r.GET("/ShouldBindBodyWith", SomeHandlerB)

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8090")
}
