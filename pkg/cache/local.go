/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/10/26 17:47
 */
package cache

import (
	"errors"
	"sync"
	"time"
)

type localData struct {
	expireAt int64
	data     []byte
}

type LocalStorage struct {
	mutex *sync.RWMutex
	data  map[string]*localData
}

func NewLocalStorage() *LocalStorage {
	return &LocalStorage{
		mutex: new(sync.RWMutex),
		data:  make(map[string]*localData),
	}
}

func (s *LocalStorage) Put(key string, b []byte, exp ...int64) error {
	data := &localData{data: b}
	if len(exp) > 0 {
		data.expireAt = time.Now().Unix() + exp[0]
	}
	s.mutex.Lock()
	s.data[key] = data
	s.mutex.Unlock()
	return nil
}

func (s *LocalStorage) Get(key string) ([]byte, error) {
	s.mutex.RLock()
	data := s.data[key]
	s.mutex.RUnlock()
	if data == nil {
		return nil, errors.New("not found")
	}
	if data.expireAt == 0 || data.expireAt > time.Now().Unix() {
		return data.data, nil
	}
	s.Delete(key)
	return nil, errors.New("not found")
}

func (s *LocalStorage) Delete(key string) error {
	s.mutex.Lock()
	delete(s.data, key)
	s.mutex.Unlock()
	return nil
}

func (s *LocalStorage) Clean() {
	s.mutex.Lock()
	for k, v := range s.data {
		if v.expireAt != 0 && v.expireAt <= time.Now().Unix() {
			delete(s.data, k)
		}
	}
	s.mutex.Unlock()
}

func (s *LocalStorage) Cleaning(dur time.Duration) {
	for {
		time.Sleep(dur)
		s.Clean()
	}
}
