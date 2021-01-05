/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/12/4 17:15
 */
package rpg

const (
	typeText  = 0
	typeImage = 1
	typeLink  = 2
)

const (
	actionSystem     = 1
	actionOnline     = 10
	actionOffline    = 11
	actionJoin       = 20
	actionLeave      = 21
	actionNewMessage = 100
)

type message struct {
	SessionID string `json:"session_id"`
	MessageID string `json:"message_id"`
	Action    int    `json:"action"`
	RoomID    int    `json:"room_id"`
	From      int    `json:"from"`
	To        int    `json:"to"`
	Content   string `json:"content"`
	Type      int    `json:"type"`
	Timestamp int64  `json:"timestamp"`
}
