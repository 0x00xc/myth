/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2021/1/5 13:26
 */
package im

import "golang.org/x/net/websocket"

func (m *Manager) Handler() websocket.Handler {
	return func(conn *websocket.Conn) { m.Join(conn) }
}
