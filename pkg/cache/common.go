/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/10/26 15:46
 */
package cache

import (
	"encoding/binary"
	"encoding/json"
	"errors"
)

type CommonStorage struct {
	Marshal   func(interface{}) ([]byte, error)
	Unmarshal func([]byte, interface{}) error
	Storage   Storage
}

func (s *CommonStorage) SetMarshal(v func(interface{}) ([]byte, error)) {
	s.Marshal = v
}
func (s *CommonStorage) SetUnmarshal(v func([]byte, interface{}) error) {
	s.Unmarshal = v
}
func (s *CommonStorage) SetStorage(v Storage) {
	s.Storage = v
}

func (s *CommonStorage) Put(key string, v []byte, exp ...int64) error {
	return s.Storage.Put(key, v, exp...)
}

func (s *CommonStorage) Get(key string) ([]byte, error) {
	return s.Storage.Get(key)
}

func (s *CommonStorage) Delete(key string) error {
	return s.Storage.Delete(key)
}

func (s *CommonStorage) Set(key string, v interface{}, exp ...int64) error {
	b, err := s.Marshal(v)
	if err != nil {
		return err
	}
	return s.Put(key, b, exp...)
}

func (s *CommonStorage) Load(key string, v interface{}) error {
	b, err := s.Get(key)
	if err != nil {
		return err
	}
	return s.Unmarshal(b, v)
}

func NewJSONStorage(s Storage) *CommonStorage {
	return &CommonStorage{
		Marshal:   json.Marshal,
		Unmarshal: json.Unmarshal,
		Storage:   s,
	}
}

type MStorage struct {
	Storage Storage
}

func (s *MStorage) Put(key string, v []byte, exp ...int64) error {
	return s.Storage.Put(key, v, exp...)
}

func (s *MStorage) Get(key string) ([]byte, error) {
	return s.Storage.Get(key)
}

func (s *MStorage) GetString(key string) (string, error) {
	b, err := s.Get(key)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (s *MStorage) PutUint8(key string, i uint8, exp ...int64) error {
	return s.Put(key, []byte{i}, exp...)
}

func (s *MStorage) PutUint16(key string, i uint16, exp ...int64) error {
	var b = make([]byte, 2)
	binary.BigEndian.PutUint16(b, i)
	return s.Put(key, b, exp...)
}

func (s *MStorage) PutUint32(key string, i uint32, exp ...int64) error {
	var b = make([]byte, 4)
	binary.BigEndian.PutUint32(b, i)
	return s.Put(key, b, exp...)
}

func (s *MStorage) PutUint68(key string, i uint64, exp ...int64) error {
	var b = make([]byte, 8)
	binary.BigEndian.PutUint64(b, i)
	return s.Put(key, b, exp...)
}

func (s *MStorage) GetUint8(key string) (uint8, error) {
	b, err := s.Get(key)
	if err != nil {
		return 0, err
	}
	if len(b) != 1 {
		return 0, errors.New("invalid data")
	}
	return b[0], nil
}
func (s *MStorage) GetUint16(key string) (uint16, error) {
	b, err := s.Get(key)
	if err != nil {
		return 0, err
	}
	if len(b) != 2 {
		return 0, errors.New("invalid data")
	}
	return binary.BigEndian.Uint16(b), nil
}
func (s *MStorage) GetUint32(key string) (uint32, error) {
	b, err := s.Get(key)
	if err != nil {
		return 0, err
	}
	if len(b) != 4 {
		return 0, errors.New("invalid data")
	}
	return binary.BigEndian.Uint32(b), nil
}
func (s *MStorage) GetUint68(key string) (uint64, error) {
	b, err := s.Get(key)
	if err != nil {
		return 0, err
	}
	if len(b) != 8 {
		return 0, errors.New("invalid data")
	}
	return binary.BigEndian.Uint64(b), nil
}

func (s *MStorage) Delete(key string) error {
	return s.Storage.Delete(key)
}

func NewStorage(s Storage) *MStorage {
	return &MStorage{Storage: s}
}
