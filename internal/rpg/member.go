/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/12/4 17:14
 */
package rpg

import (
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/net/websocket"
	"myth/internal/dao"
	"time"
)

const (
	KP = 1
	PC = 2
)

type member struct {
	session  string
	roleType int
	info     *dao.RPGRole
	conn     *websocket.Conn
	room     *room
	c        chan *message
	hb       chan bool
}

func newMember(conn *websocket.Conn, role *dao.RPGRole) *member {
	m := &member{
		session:  uuid.New().String(),
		roleType: PC,
		info:     role,
		conn:     conn,
		c:        make(chan *message),
	}
	return m
}

func (m *member) join(r *room) {
	msg := &message{
		Action:    actionJoin,
		RoomID:    r.id,
		From:      m.info.ID,
		Content:   fmt.Sprintf("%s joined", m.info.Name),
		Timestamp: time.Now().Unix(),
	}
	r.submit(msg)
	r.register(m.c)
}

func (m *member) leave() {
	if m.room != nil {
		m.room.unregister(m.c)
		msg := &message{
			SessionID: m.session,
			MessageID: uuid.New().String(),
			Action:    actionLeave,
			RoomID:    m.room.id,
			From:      m.info.ID,
			Content:   fmt.Sprintf("%s leaved", m.info.Name),
			Timestamp: time.Now().Unix(),
		}
		m.room.submit(msg)
	}
}

func (m *member) serve() {
	go m.heartbeat()
	for {
		select {
		case data := <-m.c:
			if data.SessionID == m.session {
				continue
			}
			if data.To == 0 {
				data.To = m.info.ID
			}
			m.send(data)
		case <-m.hb:
		case <-time.After(time.Minute * 5):
			m.leave()
			m.Close()
			return
		}
	}

}

func (m *member) heartbeat() {
	for {
		var data []byte
		err := websocket.Message.Receive(m.conn, &data)
		if err != nil {
			break
		}
		m.hb <- true
	}
}

func (m *member) send(data *message) error {
	return websocket.JSON.Send(m.conn, data)
}

func (m *member) Close() error {
	m.leave()
	err := m.conn.Close()
	close(m.c)
	close(m.hb)
	return err
}
