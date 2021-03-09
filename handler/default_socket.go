package handler

import (
	"github.com/gorilla/websocket"
	"go-chat/dto"
	"log"
	"net/http"
)

func (d *_default) HandleConnection(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer ws.Close()

	d.client[ws] = true

	for {
		var msg dto.Message

		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(d.client, ws)
			break
		}

		d.broadcast <- msg
	}
}