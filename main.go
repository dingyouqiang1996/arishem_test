package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 拿到结构体自定义计算规则
func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H {
			"title": "Main website",
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H {
			"message": "pong",
		})
	})
	r.Run()
}