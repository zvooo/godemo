package server

import "github.com/gorilla/websocket"

type Client struct {
	Conn *websocket.Conn
	Name string
	Send chan []byte
}

func (c *Client) Read() {
	defer func() {
		HubServer.UnRegister <- c
		c.Conn.Close()
	}()

	for {
		c.Conn.PingHandler()
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			HubServer.UnRegister <- c
			c.Conn.Close()
		}
		HubServer.Broadcast <- message
	}
}
func (c *Client) Write() {
	defer func() {

		c.Conn.Close()
	}()
	for message := range c.Send {
		c.Conn.WriteMessage(websocket.BinaryMessage, message)
	}
}
