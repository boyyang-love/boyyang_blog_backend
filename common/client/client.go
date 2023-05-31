package client

import (
	"bytes"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024 * 10,
	WriteBufferSize: 1024 * 10,
}

type Client struct {
	conn *websocket.Conn
	id   int64
	hub  *Hub
	send chan []byte
}

func (c *Client) readMessage() {
	defer func() {
		c.hub.unregister <- c
	}()
	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			return
		}

		message := bytes.TrimSpace(bytes.Replace(msg, newline, space, -1))
		c.hub.broadcast <- message
	}
}

func (c *Client) sendMessage() {
	for {
		select {
		case message := <-c.send:
			if string(message) == "ping" {
				err := c.conn.WriteMessage(websocket.TextMessage, []byte("pong"))
				if err != nil {
					return
				}
			} else {
				err := c.conn.WriteMessage(websocket.TextMessage, message)
				if err != nil {
					return
				}
			}
		}
	}
}

func WsRun(h *Hub, w http.ResponseWriter, r *http.Request) {
	defer func() {
		fmt.Println("WsRun end")
	}()
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	id, _ := r.Context().Value("Id").(int64)
	client := &Client{
		conn: conn,
		id:   id,
		hub:  h,
		send: make(chan []byte, 1024*10),
	}

	go client.hub.HubRun()

	client.hub.register <- client

	go client.readMessage()
	go client.sendMessage()
}
