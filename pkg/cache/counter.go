/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/10/30 13:39
 */
package cache

import (
	"sync"
	"time"
)

type counter struct {
	t     int64
	count int
}

func (c *counter) add(durationCheck int) int {
	now := time.Now().Unix()
	if int(now-c.t) > durationCheck {
		c.count = 1
		c.t = now
	} else {
		c.count++
	}
	return c.count
}

type Counter struct {
	locker   sync.Mutex
	m        map[string]*counter
	duration int
	limit    int
}

func NewCounter(duration int, limit int) *Counter {
	return &Counter{
		m:        make(map[string]*counter),
		duration: duration,
		limit:    limit,
	}
}

func (f *Counter) check(key string) int {
	f.locker.Lock()
	if f.m[key] == nil {
		f.m[key] = new(counter)
	}
	v := f.m[key].add(f.duration)
	f.locker.Unlock()
	return v
}

func (f *Counter) Flush() {
	f.locker.Lock()
	now := time.Now().Unix()
	for k, v := range f.m {
		if int(now-v.t) > f.duration*10 {
			delete(f.m, k)
		}
	}
	f.locker.Unlock()
}

func (f *Counter) Len() (n int) {
	f.locker.Lock()
	n = len(f.m)
	f.locker.Unlock()
	return
}

func (f *Counter) Banned(ip string) bool {
	return f.check(ip) > f.limit
}
