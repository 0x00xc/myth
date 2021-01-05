/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/12/30 18:01
 */
package im

const (
	ClientPing = 0
)

type ClientMessage struct {
	ID   string
	Code int
	Data string
}

type Messenger interface {
	Send(to string, msg []byte)
}
