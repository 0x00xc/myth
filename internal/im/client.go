/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/12/30 17:54
 */
package im

import (
	"encoding/json"
	"golang.org/x/net/websocket"
	"time"
)

type client struct {
	id        string
	conn      *websocket.Conn
	exit      chan bool
	heartbeat chan bool
	timeout   time.Duration
	//broadcast chan *ServerMessage
}

func newClient(id string, conn *websocket.Conn, timeout time.Duration) *client {
	return &client{
		id:        id,
		conn:      conn,
		exit:      make(chan bool),
		heartbeat: make(chan bool),
		timeout:   timeout,
	}
}

func (c *client) serve() {
	if c.timeout == 0 {
		c.timeout = time.Minute * 5
	}
	go c.accept()
	for {
		select {
		case <-time.After(c.timeout): //timeout
			c.stop()
		case <-c.heartbeat:
		//case <-c.messages:
		case <-c.exit:
			return
		}
	}
}

func (c *client) stop() {
	c.conn.Close()
}

func (c *client) accept() {
	defer close(c.exit)

	for {
		var data []byte
		err := websocket.Message.Receive(c.conn, &data)
		if err != nil {
			break
		}
		var msg = new(ClientMessage)
		err = json.Unmarshal(data, msg)
		if err != nil {
			continue
		}
		if msg.Code == ClientPing {
			c.heartbeat <- true
		} else {
			//c.messages <- msg
		}
	}
}

func (c *client) write(msg []byte) error {
	return websocket.Message.Send(c.conn, msg)
}
