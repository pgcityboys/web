package wsserver

import (
	"fmt"
	"web/rabbit"
)

type Hub struct {
	// Registered clients.
	clients map[*Client]string

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	// Assign client id to internal db
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]string),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = "nac"
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			handleInMessage(message)
		case chat := <-rabbit.ChatChannel:
			for client := range h.clients {
				select {
				case client.send <- []byte(fmt.Sprintf("%s: %s", chat.UserId, chat.Content)):
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
