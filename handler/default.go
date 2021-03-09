package handler

import (
	"github.com/gorilla/websocket"
	"go-chat/dto"
)

type _default struct {
	client map[*websocket.Conn]bool
	broadcast chan dto.Message
}

func NewDefault(client map[*websocket.Conn]bool, broadcast chan dto.Message) _default {
	return _default{client: client, broadcast: broadcast}
}