/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/12/30 17:54
 */
package im

import (
	"golang.org/x/net/websocket"
	"time"
)

type client struct {
	id        string
	conn      *websocket.Conn
	exit      chan bool
	heartbeat chan bool
	timeout   time.Duration
	ping      func([]byte) bool
	onMessage func(string, []byte, Messenger)
}

func newClient(id string, conn *websocket.Conn, timeout time.Duration, ping func([]byte) bool, onMessage func(string, []byte, Messenger)) *client {
	return &client{
		id:        id,
		conn:      conn,
		exit:      make(chan bool),
		heartbeat: make(chan bool),
		timeout:   timeout,
		ping:      ping,
		onMessage: onMessage,
	}
}

func (c *client) serve(m Messenger) {
	if c.timeout == 0 {
		c.timeout = time.Minute * 5
	}
	go c.accept(m)
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

func (c *client) accept(m Messenger) {
	defer close(c.exit)

	for {
		var data []byte
		err := websocket.Message.Receive(c.conn, &data)
		if err != nil {
			break
		}
		if c.ping(data) {
			c.heartbeat <- true
		} else {
			c.onMessage(c.id, data, m)
		}
	}
}

func (c *client) write(msg []byte) error {
	return websocket.Message.Send(c.conn, msg)
}
