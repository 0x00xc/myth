/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/12/4 17:26
 */
package queue

type broadcaster struct {
	input chan []byte
	reg   chan chan<- []byte
	unreg chan chan<- []byte

	outputs map[chan<- []byte]bool
}

// The Broadcaster interface describes the main entry points to
// broadcasters.
type Broadcaster interface {
	// Register a new channel to receive broadcasts
	Register(chan<- []byte)
	// Unregister a channel so that it no longer receives broadcasts.
	Unregister(chan<- []byte)
	// Shut this broadcaster down.
	Close() error
	// Submit a new object to all subscribers
	Submit([]byte)
}

func (b *broadcaster) broadcast(data []byte) {
	for ch := range b.outputs {
		select {
		case ch <- data:
			//default: //写入失败直接丢弃
			//	continue
		}
	}
}

func (b *broadcaster) run() {
	for {
		select {
		case m := <-b.input:
			b.broadcast(m)
		case ch, ok := <-b.reg:
			if ok {
				b.outputs[ch] = true
			} else {
				return
			}
		case ch := <-b.unreg:
			delete(b.outputs, ch)
		}
	}
}

// NewBroadcaster creates a new broadcaster with the given input
// channel buffer length.
func NewBroadcaster(buflen int) Broadcaster {
	b := &broadcaster{
		input:   make(chan []byte, buflen),
		reg:     make(chan chan<- []byte),
		unreg:   make(chan chan<- []byte),
		outputs: make(map[chan<- []byte]bool),
	}

	go b.run()
	return b
}

func (b *broadcaster) Register(newch chan<- []byte) {
	b.reg <- newch
}

func (b *broadcaster) Unregister(newch chan<- []byte) {
	b.unreg <- newch
}

func (b *broadcaster) Close() error {
	close(b.reg)
	return nil
}

func (b *broadcaster) Submit(m []byte) {
	if b != nil {
		b.input <- m
	}
}
