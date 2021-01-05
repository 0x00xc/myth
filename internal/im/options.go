/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2021/1/5 10:44
 */
package im

import "time"

type Options struct {
	Heartbeat          time.Duration //心跳间隔（应当略大于客户端发送间隔）
	MessageLimit       int           //
	MessageWaitTimeout time.Duration //
	MaxConnection      int
}

func (o *Options) mDefault() {
	if o.Heartbeat == 0 {
		o.Heartbeat = time.Minute * 5
	}
	if o.MessageLimit <= 0 {
		o.MessageLimit = 100
	}
	if o.MaxConnection <= 0 {
		o.MaxConnection = 9999
	}
}
