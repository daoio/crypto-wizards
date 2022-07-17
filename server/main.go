package main

import "github.com/gin-gonic/gin"

func main() {
	hub := newHub()
	go hub.run()
	r := gin.New()
	r.GET("/ws", func(hub *Hub) gin.HandlerFunc {
		return gin.HandlerFunc(func(c *gin.Context) {
			serveWs(hub, c.Writer, c.Request)
		})
	}(hub))
	r.Run(":8080")
}
