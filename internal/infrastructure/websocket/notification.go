package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type notificationWebSocket struct {
}

func NewNotificationWebSocket() *notificationWebSocket {
	return &notificationWebSocket{}
}

func (soc notificationWebSocket) RegisterNotificationWebSocket(e *echo.Echo) {
	e.GET("/ws", soc.ConnectNotificationWebSocket)

	// Implementation for registering notification WebSocket goes here
}

func (soc notificationWebSocket) ConnectNotificationWebSocket(c echo.Context) error {
	w := c.Response().Writer
	r := c.Request()
	wHeader := w.Header()
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(w, r, wHeader)
	if err != nil {
		log.Println(err)
		return err
	}

	defer conn.Close()

	for {
		// Read
		msgType, msg, err := conn.ReadMessage()
		fmt.Printf("type: %v\n", msgType)

		if err != nil {
			log.Println("read:", err.Error())
			break
		}
		fmt.Printf("recv: %s\n", msg)

		// Write
		err = conn.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
		if err != nil {
			log.Println("write:", err)
			break
		}
	}

	// websocket.Handler(func(ws *websocket.Conn) {
	// 	defer ws.Close()
	// 	for {

	// 		//
	// 		// Write
	// 		err := websocket.Message.Send(ws, "Hello, Client!")
	// 		if err != nil {
	// 			c.Logger().Error(err)
	// 		}

	// 		// receive close message from client

	// 		// Read
	// 		msg := ""
	// 		err = websocket.Message.Receive(ws, &msg)
	// 		if err != nil {
	// 			c.Logger().Error(err)
	// 		}
	// 		fmt.Printf("%s\n", msg)
	// 	}
	// }).ServeHTTP(c.Response(), c.Request())
	return nil
}
