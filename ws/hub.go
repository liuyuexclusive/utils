// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ws

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan *Broadcast

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte

	username string
}

type Broadcast struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan *Broadcast),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			targets := make(map[*Client]bool)
			if message.To == nil || len(message.To) == 0 {
				targets = h.clients
			} else {
				for a := range h.clients {
				loop:
					for _, b := range message.To {
						if a.username == b {
							targets[a] = true
							break loop
						}
					}
				}
			}
			for client := range targets {
				if bytes, err := json.Marshal(message); err == nil {
					select {
					case client.send <- bytes:
					default:
						close(client.send)
						delete(h.clients, client)
					}
				}
			}
		}
	}
}
