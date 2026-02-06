package websocket

import (
	"context"
	"fmt"
	"log"
	"net/http"

	mAuth "clone-instagram-service/internal/domain/model/auth"
	mNoti "clone-instagram-service/internal/domain/model/notification"
	eRela "clone-instagram-service/internal/domain/model/relationship/entity"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type notificationWebSocket struct {
	notificationSubscriber mNoti.NotificationSubscriber
	authMiddleWare         mAuth.AuthMiddleWare
}

type notificationSocketConnection struct {
	notificationSubscriber mNoti.NotificationSubscriber
	authMiddleWare         mAuth.AuthMiddleWare
	userID                 int
	conn                   *websocket.Conn
}

func NewNotificationWebSocket(notificationSubscriber mNoti.NotificationSubscriber, authMiddleWare mAuth.AuthMiddleWare) *notificationWebSocket {
	return &notificationWebSocket{
		notificationSubscriber: notificationSubscriber,
		authMiddleWare:         authMiddleWare,
	}
}

func (soc notificationWebSocket) RegisterNotificationWebSocket(e *echo.Echo) {
	e.GET("/ws", soc.ConnectNotificationWebSocket)

	// Implementation for registering notification WebSocket goes here
}

func (soc notificationWebSocket) ConnectNotificationWebSocket(c echo.Context) error {

	w := c.Response().Writer
	r := c.Request()

	token := r.URL.Query().Get("accessToken")
	fmt.Println("token -> ", token)
	user, err := soc.authMiddleWare.GetUserInfoByAccessToken(c.Request().Context(), token)
	if err != nil {
		log.Println("Error => unauthorize => ", err.Error())
	}

	wHeader := w.Header()
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(w, r, wHeader)
	fmt.Println("connect ", conn.LocalAddr().String())
	fmt.Println("connect ", conn.RemoteAddr().String())
	if err != nil {
		log.Println(err)
		return err
	}

	socConn := notificationSocketConnection{
		userID:                 user.ID, // TODO: get userID from token
		conn:                   conn,
		notificationSubscriber: soc.notificationSubscriber,
	}

	socConn.liveConnection()

	fmt.Println("end")

	return nil
}

func (socConn notificationSocketConnection) followingNotiCallback(ctx context.Context, message eRela.FollowingTopicMessage) error {
	msg := fmt.Sprintf("User %d followed you", message.UserID)
	err := socConn.conn.WriteMessage(websocket.TextMessage, []byte(msg))
	return err
}

func (socConn notificationSocketConnection) liveConnection() {
	conCtx, cancelConCtx := context.WithCancel(context.Background())
	defer func() {
		fmt.Println("closing connection")
		socConn.conn.Close()
		cancelConCtx()
	}()

	// sub topic and push message to client
	// soccon must have noti sub
	go socConn.notificationSubscriber.SubscribeFollowingWithUserID(conCtx, socConn.userID, socConn.followingNotiCallback)

	for {
		// Read
		msgType, msg, err := socConn.conn.ReadMessage()
		fmt.Printf("type: %v\n", msgType)
		if msgType == websocket.CloseMessage || msgType == -1 {
			fmt.Println("Client disconnected")
			break
		}

		if err != nil {
			log.Println("read:", err.Error())
			break
		}
		fmt.Printf("recv: %s\n", msg)

		// Write
		err = socConn.conn.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
