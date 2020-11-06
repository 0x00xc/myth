/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/10/27 11:10
 */
package cache

import (
	"container/list"
	"errors"
	"sync"
	"time"
)

type LRU struct {
	locker sync.Mutex
	data   *list.List
	index  map[string]*list.Element
	max    int
}

type lruData struct {
	id   string
	t    int64
	data []byte
}

func NewLRUCache(max int) *LRU {
	if max < 1 {
		panic("invalid cache size")
	}
	c := &LRU{}
	c.data = list.New()
	c.index = make(map[string]*list.Element)
	c.max = max
	return c
}

func (c *LRU) Put(id string, v []byte, exp ...int64) error {
	c.locker.Lock()
	defer c.locker.Unlock()
	if c.data.Len() >= c.max {
		back := c.data.Back()
		c.data.Remove(back)
		delete(c.index, back.Value.(*lruData).id)
	}
	data := &lruData{id: id, data: v}
	if len(exp) > 0 {
		data.t = time.Now().Unix() + exp[0]
	}
	c.index[id] = c.data.PushFront(data)
	return nil
}

func (c *LRU) Get(id string) ([]byte, error) {
	c.locker.Lock()
	defer c.locker.Unlock()
	if ele, exist := c.index[id]; exist {
		value := ele.Value.(*lruData)
		if value.t == 0 || value.t > time.Now().Unix() {
			c.data.MoveToFront(ele)
			return value.data, nil
		} else {
			c.data.Remove(ele)
			delete(c.index, id)
		}
	}
	return nil, errors.New("not found")
}

func (c *LRU) Delete(id string) error {
	c.locker.Lock()
	defer c.locker.Unlock()
	if ele, exist := c.index[id]; exist {
		c.data.Remove(ele)
		delete(c.index, id)
	}
	return nil
}
