package main

import (
	"github.com/gorilla/websocket"
	"go-chat/dto"
	route "go-chat/router"
	"go-chat/util"
	"log"
	"net/http"
)

func main() {
	var clients = make(map[*websocket.Conn]bool)
	var broadcast = make(chan dto.Message)

	router := route.NewRouter(clients, broadcast)

	go util.HandleMessages(broadcast, clients)

	log.Fatal("http server started in: 8000", http.ListenAndServe(":8000", router))
}