package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/JSONP", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}
		// http://localhost:8080/JSONP?callback=x
		// x({"foo":"bar"})
		c.JSONP(http.StatusOK, data)
	})

	r.Run(":8080")
}
