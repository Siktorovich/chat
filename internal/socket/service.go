package socket

import "github.com/gorilla/websocket"

type ConnectUser struct {
	Websocket *websocket.Conn
	ClientIP  string
}

func NewConnectUser(conn *websocket.Conn, clientIP string) *ConnectUser {
	return &ConnectUser{
		Websocket: conn,
		ClientIP:  clientIP,
	}
}
