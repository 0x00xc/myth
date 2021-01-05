/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/12/30 17:55
 */
package im

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
	"sync"
	"time"
)

type Manager struct {
	locker  *sync.RWMutex
	clients map[string]*client
	queue   Queue
	exit    chan bool
	limit   chan bool
	opt     *Options

	Auth           func(conn *websocket.Conn) (string, error)
	OnConnected    func(id string, messenger Messenger)
	OnDisConnected func(id string, messenger Messenger)
}

func NewManager(opt *Options, queue Queue) *Manager {
	if opt == nil {
		opt = new(Options)
	}
	if queue == nil {
		queue = NewQueue(128)
	}
	opt.mDefault()
	return &Manager{
		locker:  new(sync.RWMutex),
		clients: make(map[string]*client),
		queue:   queue,
		exit:    make(chan bool),
		opt:     opt,
		limit:   make(chan bool, opt.MessageLimit),
		Auth: func(conn *websocket.Conn) (string, error) {
			return conn.RemoteAddr().String(), nil
		},
		OnConnected: func(id string, messenger Messenger) {

		},
		OnDisConnected: func(id string, messenger Messenger) {

		},
	}
}

func (m *Manager) Send(to string, msg []byte) {
	m.queue.Put(Message{To: to, Data: msg})
}

func (m *Manager) Stop() {
	close(m.exit)
}

func (m *Manager) IsOnline(id string) bool {
	m.locker.RLock()
	defer m.locker.RUnlock()
	return m.clients[id] != nil
}

func (m *Manager) Serve() {
	m.serve()
}

func (m *Manager) Join(conn *websocket.Conn) error {
	id, err := m.Auth(conn)
	if err != nil {
		return err
	}
	m.OnConnected(id, m)
	defer m.OnDisConnected(id, m)
	c := newClient(id, conn, m.opt.Heartbeat)
	m.join(c)
	defer m.remove(id)
	c.serve()
	return nil
}

func (m *Manager) serve() {
	for {
		select {
		case msg := <-m.queue.Fetch():
			m.doMessage(msg)
		case <-m.exit:
			return
		}
	}
}

func (m *Manager) doMessage(msg Message) {
	if m.opt.MessageWaitTimeout > 0 {
		select {
		case m.limit <- true:
		case <-time.After(m.opt.MessageWaitTimeout):
			//TODO m.queue.Put(msg)
			return
		case <-m.exit:
			//TODO m.queue.Put(msg)
			return
		}
	} else {
		select {
		case m.limit <- true:
		case <-m.exit:
			//TODO m.queue.Put(msg)
			return
		}
	}
	go m.send(msg.To, msg.Data)
}

func (m *Manager) send(to string, msg []byte) {
	defer func() { <-m.limit }()
	m.locker.RLock()
	defer m.locker.RUnlock()
	c := m.clients[to]
	if c != nil {
		err := c.write(msg)
		if err != nil {
			logrus.Errorln("send failed:", err)
			//TODO m.queue.Put(msg)
		}
	}
	//client not found
	//TODO m.queue.Put(msg)
}

func (m *Manager) size() int {
	m.locker.RLock()
	defer m.locker.RUnlock()
	return len(m.clients)
}

func (m *Manager) join(c *client) {
	m.locker.Lock()
	if m.clients[c.id] != nil {
		m.clients[c.id].stop()
	}
	m.clients[c.id] = c
	m.locker.Unlock()
}

func (m *Manager) remove(id string) {
	m.locker.Lock()
	delete(m.clients, id)
	m.locker.Unlock()
}
