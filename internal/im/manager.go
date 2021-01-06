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
}

func NewManager(opt *Options) *Manager {
	if opt == nil {
		opt = new(Options)
	}
	opt.mDefault()
	return &Manager{
		locker:  new(sync.RWMutex),
		clients: make(map[string]*client),
		queue:   opt.Queue,
		exit:    make(chan bool),
		opt:     opt,
		limit:   make(chan bool, opt.MessageLimit),
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
	var id string
	var err error
	for i := 0; i < m.opt.AuthRetry; i++ {
		id, err = m.opt.Auth(conn)
		if err == nil {
			break
		}
	}
	if err != nil {
		return err
	}
	m.opt.OnClientConnected(id, m)
	defer m.opt.OnClientDisConnected(id, m)
	c := newClient(id, conn, m.opt.Heartbeat, m.opt.IsClientPing, m.opt.OnClientMessage)
	m.join(c)
	defer m.remove(id)
	c.serve(m)
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
	} else if m.opt.MessageWaitTimeout < 0 {
		select {
		case m.limit <- true:
		default:
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
