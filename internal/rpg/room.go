/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/12/4 17:13
 */
package rpg

type room struct {
	id    int
	kp    *member
	title string

	input   chan *message
	reg     chan chan<- *message
	unreg   chan chan<- *message
	outputs map[chan<- *message]bool
}

func newRoom(id int, title string, kp *member) *room {
	return &room{
		id:      id,
		title:   title,
		kp:      kp,
		input:   make(chan *message, 16),
		reg:     make(chan chan<- *message),
		unreg:   make(chan chan<- *message),
		outputs: make(map[chan<- *message]bool),
	}
}

func (r *room) count() int {
	return len(r.reg)
}

//func NewBroadcaster(buflen int) Broadcaster {
//	b := &broadcaster{
//		input:   make(chan []byte, buflen),
//		reg:     make(chan chan<- []byte),
//		unreg:   make(chan chan<- []byte),
//		outputs: make(map[chan<- []byte]bool),
//	}
//
//	go b.run()
//	return b
//}

func (r *room) submit(m *message) {
	if r != nil {
		r.input <- m
	}
}

func (r *room) run() {
	for {
		select {
		case m := <-r.input:
			r.broadcast(m)
		case ch, ok := <-r.reg:
			if ok {
				r.outputs[ch] = true
			} else {
				return
			}
		case ch := <-r.unreg:
			delete(r.outputs, ch)
		}
	}
}

func (r *room) broadcast(data *message) {
	for ch := range r.outputs {
		select {
		case ch <- data:
			//default: //写入失败直接丢弃
			//	continue
		}
	}
}

func (r *room) register(newch chan<- *message) {
	r.reg <- newch
}

func (r *room) unregister(newch chan<- *message) {
	r.unreg <- newch
}

func (r *room) Close() error {
	close(r.reg)
	return nil
}
