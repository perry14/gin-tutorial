package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/static", "static")
	router.LoadHTMLGlob("templates/*")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "JSONP TEST",
		})
	})

	router.GET("/JSONP", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}
		// http://localhost:8080/JSONP?callback=x
		// x({"foo":"bar"})
		c.JSONP(http.StatusOK, data)
	})

	router.Run(":8080")
}
