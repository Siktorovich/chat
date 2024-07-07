package chat

import (
	"chat/internal/socket"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var users = make(map[socket.ConnectUser]int64)

func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)

	defer func() {
		if err := conn.Close(); err != nil {
			log.Println("Websocket could not be closed", err.Error())
		}
	}()

	log.Println("Client connected:", conn.RemoteAddr().String())

	socketCLient := socket.NewConnectUser(conn, conn.RemoteAddr().String())

	users[*socketCLient] = 0

	log.Println("Number client connected ...", len(users))

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Ws disconnect waiting", err.Error())
			delete(users, *socketCLient)
			log.Println("Number of client still connected ...", len(users))
			return
		}

		for client := range users {
			if err = client.Websocket.WriteMessage(messageType, message); err != nil {
				log.Println("Could not send Message to ", client.ClientIP, err.Error())
			}
		}
	}
}
