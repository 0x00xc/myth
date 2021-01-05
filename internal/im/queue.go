/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2021/1/5 10:34
 */
package im

type Message struct {
	To   string
	Data []byte
}

type Queue interface {
	Put(data Message) error
	Fetch() <-chan Message
}

type mQueue struct {
	c chan Message
}

func (m *mQueue) Put(data Message) error {
	m.c <- data
	return nil
}

func (m *mQueue) Fetch() <-chan Message {
	return m.c
}

func NewQueue(size int) Queue {
	return &mQueue{c: make(chan Message, size)}
}
