package server

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func FindNewGame() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Do you want to start new game? [y/n]")
	scanner.Scan()

	switch scanner.Text() {
	case "y":
		connect()
	case "n":
		log.Fatal(":(")
	}
}

func StartServer() {
	hub := newHub()
	go hub.run()

	r := gin.New()
	r.GET("/ws", func(hub *Hub) gin.HandlerFunc {
		return gin.HandlerFunc(func(c *gin.Context) {
			serveWs(hub, c.Writer, c.Request)
		})
	}(hub))
	r.Run(":300")
}

func connect() {
	conn, _, _ := websocket.DefaultDialer.Dial("ws://127.0.0.1:3000/ws", nil)

	go func(c *websocket.Conn) {
		// TODO:
	}(conn)
}
