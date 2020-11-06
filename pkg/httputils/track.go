/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/10/23 16:40
 */
package httputils

import (
	"fmt"
	"net/url"
	"time"
)

type Track struct {
	Timestamp     time.Time
	Method        string
	Address       string
	StatusCode    int
	Status        string
	ContentLength int64
	Error         error
	Latency       time.Duration
	u             *url.URL
	e             error
}

func (t *Track) URL() (*url.URL, error) {
	if t.u == nil && t.e == nil {
		t.u, t.e = url.Parse(t.Address)
	}
	return t.u, t.e
}

func (t *Track) Addr() string {
	u, err := t.URL()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%s://%s/%s", u.Scheme, u.Host, u.Path)
}

func (t *Track) Host() string {
	u, err := t.URL()
	if err != nil {
		return ""
	}
	return u.Host
}

func (t *Track) Path() string {
	u, err := t.URL()
	if err != nil {
		return ""
	}
	return u.Path
}

type Tracker interface {
	Trace(track *Track)
}

type wrapper struct {
	t Tracker
	q chan bool
}

func newWrapper(t Tracker, n int) Tracker {
	return &wrapper{t: t, q: make(chan bool, n)}
}

func (w *wrapper) Trace(data *Track) {
	select {
	case w.q <- true:
		go w.do(data)
	default:
	}
}

func (w *wrapper) do(data *Track) {
	w.t.Trace(data)
	<-w.q
}
