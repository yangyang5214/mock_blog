package main

import (
	"github.com/gin-gonic/gin"
	h "mock_blog/src/handle"
	s "mock_blog/src/service"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Use(h.Cors())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/mock", func(c *gin.Context) {
		body := s.MockModel{}
		c.BindJSON(&body)
		s.Mock{}.InsertMock(body)
		c.JSON(http.StatusOK, nil)
	})

	r.Run(":9999") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
