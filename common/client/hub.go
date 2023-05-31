package client

import (
	"fmt"
)

type Hub struct {
	broadcast  chan []byte
	message    chan []byte
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		message:    make(chan []byte),
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) HubRun() {
	for {
		select {
		case client := <-h.register:
			h.broadcast <- []byte(fmt.Sprintf("user(id=%v) has connected.\n", client.id))
			h.clients[client] = true
		case unregister := <-h.unregister:
			delete(h.clients, unregister)
		case boardcastMessage := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- boardcastMessage:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
