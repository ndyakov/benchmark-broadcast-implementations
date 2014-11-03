package broadcastr

type AsyncBroadcaster struct {
	closed      bool
	subscribers []*AsyncSubscriber
	input       chan interface{}
	lockc chan struct{}
}

func NewAsyncBroadcaster() *AsyncBroadcaster {
	b := new(AsyncBroadcaster)
	b.subscribers = make([]*AsyncSubscriber, 0)
	b.input = make(chan interface{})
	b.lockc = make(chan struct{}, 1)
	b.closed = false
	go b.listen()
	return b
}
func (b *AsyncBroadcaster) listen() {
	for v := range b.input {
		go b.broadcast(v)
	}
}

func (b *AsyncBroadcaster) broadcast(v interface{}) {
	for _, s := range b.subscribers {
		if !s.Closed() {
			go func(inbox chan<- interface{}, v interface{}) {
				inbox <- v
			}(s.inbox, v)
		}
	}
}

func (b *AsyncBroadcaster) Send(v interface{}) (err error) {
	b.lock()
	defer b.unlock()
	if b.closed {
		return newClosedError()
	}
	b.input <- v
	return nil
}

func (b *AsyncBroadcaster) Subscribe() Subscriber {
	b.lock()
	defer b.unlock()
	s := newAsyncSubscriber(0, len(b.subscribers))
	b.subscribers = append(b.subscribers, s)
	return s
}

func (b *AsyncBroadcaster) SubscribeWithBuffer(bufferSize int) *AsyncSubscriber {
	s := newAsyncSubscriber(bufferSize, len(b.subscribers))
	b.subscribers = append(b.subscribers, s)
	return s
}

func (b *AsyncBroadcaster) FlushAllBuffers() {
	b.lock()
	defer b.unlock()
	for _, s := range b.subscribers {
		go s.flushBuffer()
	}
}

func (b *AsyncBroadcaster) Close() {
	b.lock()
	defer b.unlock()
	b.closed = true
	close(b.input)
	for _, s := range b.subscribers {
		s.Close()
	}
}

func (b *AsyncBroadcaster) lock() {
	b.lockc <- struct{}{}
}

func (b *AsyncBroadcaster) unlock() {
	<-b.lockc
}
