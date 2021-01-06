/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2021/1/5 10:44
 */
package im

import (
	"golang.org/x/net/websocket"
	"strings"
	"time"
)

type Options struct {
	MessageLimit         int                                               //消息队列容量限制
	MessageWaitTimeout   time.Duration                                     //消息队列等待超时时间（设置为0永不超时,-1立即超时）
	MaxConnection        int                                               //最大连接数
	Auth                 func(*websocket.Conn) (string, error)             //连接认证，认证成功返回用户唯一标识，认证失败返回错误，将断开连接
	AuthRetry            int                                               //认证允许重试次数
	OnClientConnected    func(id string, messenger Messenger)              //建立连接触发（认证成功后触发，认证成功前先触发Auth方法，没必要再加一个回调）
	OnClientDisConnected func(id string, messenger Messenger)              //连接断开触发
	OnClientMessage      func(id string, data []byte, messenger Messenger) //收到客户端消息触发
	IsClientPing         func(data []byte) bool                            //是否是客户端心跳消息
	Heartbeat            time.Duration                                     //心跳间隔（应当略大于客户端发送间隔）
	Queue                Queue                                             //消息队列
}

func NewOptions() *Options {
	o := new(Options)
	o.mDefault()
	return o
}

func (o *Options) mDefault() {
	if o.Heartbeat == 0 {
		o.Heartbeat = time.Minute * 6
	}
	if o.MessageLimit <= 0 {
		o.MessageLimit = 100
	}
	if o.MaxConnection <= 0 {
		o.MaxConnection = 9999
	}
	if o.Auth == nil {
		o.Auth = func(conn *websocket.Conn) (string, error) { return conn.RemoteAddr().String(), nil }
	}
	if o.AuthRetry <= 0 {
		o.AuthRetry = 1
	}
	if o.OnClientConnected == nil {
		o.OnClientConnected = func(id string, messenger Messenger) {}
	}
	if o.OnClientDisConnected == nil {
		o.OnClientDisConnected = func(id string, messenger Messenger) {}
	}
	if o.IsClientPing == nil {
		o.IsClientPing = func(b []byte) bool { return strings.TrimSpace(string(b)) == "ping" }
	}
	if o.OnClientMessage == nil {
		o.OnClientMessage = func(id string, data []byte, m Messenger) {}
	}
	if o.Queue == nil {
		o.Queue = NewQueue(128)
	}
}
