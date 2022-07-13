package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.New()
	r.GET("/ws", func() gin.HandlerFunc {
		return gin.HandlerFunc()
	}())
	r.Run(":8080")
}
