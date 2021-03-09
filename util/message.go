package util

import (
	"github.com/gorilla/websocket"
	"go-chat/dto"
	"log"
)

func HandleMessages(broadcast chan dto.Message, clients map[*websocket.Conn]bool) {
	for {
		msg := <-broadcast

		for client := range clients {
			err := client.WriteJSON(msg)

			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}