package server

import "fmt"

type Hub struct {
	Register   chan *Client
	UnRegister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan []byte
}

var HubServer = NewHub()

func NewHub() *Hub {
	h := &Hub{
		Register:   make(chan *Client),
		UnRegister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan []byte),
	}
	return h
}

func (h *Hub) Start() {
	for {
		select {
		case conn := <-h.Register:
			h.Clients[conn] = true
			msg := fmt.Sprintf("system:welcome %s\n", conn.Name)
			conn.Send <- []byte(msg)
		case conn := <-h.UnRegister:
			if _, ok := h.Clients[conn]; ok {
				close(conn.Send)
				delete(h.Clients, conn)
			}
		case Message := <-h.Broadcast:
			for conn := range h.Clients {
				select {
				case conn.Send <- Message:
				default:
					close(conn.Send)
					delete(h.Clients, conn)
				}
			}
		}

	}

}
