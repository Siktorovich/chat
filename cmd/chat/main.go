package main

import (
	"log"
	"net/http"

	impl "chat/internal/app/chat"
)

func init() {
	http.HandleFunc("/", impl.IndexHandler)
	http.HandleFunc("/ws", impl.WebsocketHandler)
}

func main() {
	log.Fatal(http.ListenAndServe("logalhost:8000", nil))
}
