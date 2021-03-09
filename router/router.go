package router

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"go-chat/dto"
	handle "go-chat/handler"
)

func NewRouter(clients map[*websocket.Conn]bool, broadcast chan dto.Message) *mux.Router {
	router := mux.NewRouter()
	
	handler := handle.NewDefault(clients, broadcast)
	router.HandleFunc("/message", handler.HandleConnection)

	return router
}
