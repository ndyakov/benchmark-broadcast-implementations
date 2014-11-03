package broadcastr

import "sync"

type SyncBroadcaster struct {
	closed      bool
	subscribers []*SyncSubscriber
	input       chan interface{}
	lockm       *sync.Mutex
	//lockc chan struct{}
}

func NewSyncBroadcaster() *SyncBroadcaster {
	b := new(SyncBroadcaster)
	b.subscribers = make([]*SyncSubscriber, 0)
	b.input = make(chan interface{})
	//b.lockc = make(chan struct{}, 1)
	b.lockm = new(sync.Mutex)
	b.closed = false
	go b.listen()
	return b
}
func (b *SyncBroadcaster) listen() {
	for v := range b.input {
		b.broadcast(v)
	}
}

func (b *SyncBroadcaster) broadcast(v interface{}) {
	for _, s := range b.subscribers {
		if !s.Closed() {
			s.inbox <- v
		}
	}
}

func (b *SyncBroadcaster) Send(v interface{}) (err error) {
	b.lock()
	defer b.unlock()
	if b.closed {
		return newClosedError()
	}
	b.input <- v
	return nil
}

func (b *SyncBroadcaster) Subscribe() Subscriber {
	b.lock()
	defer b.unlock()
	s := newSyncSubscriber(0, len(b.subscribers))
	b.subscribers = append(b.subscribers, s)
	return s
}

func (b *SyncBroadcaster) SubscribeWithBuffer(bufferSize int) *SyncSubscriber {
	s := newSyncSubscriber(bufferSize, len(b.subscribers))
	b.subscribers = append(b.subscribers, s)
	return s
}

func (b *SyncBroadcaster) FlushAllBuffers() {
	b.lock()
	defer b.unlock()
	for _, s := range b.subscribers {
		s.flushBuffer()
	}
}

func (b *SyncBroadcaster) Close() {
	b.lock()
	defer b.unlock()
	b.closed = true
	close(b.input)
	for _, s := range b.subscribers {
		s.Close()
	}
}

func (b *SyncBroadcaster) lock() {
	b.lockm.Lock()
	//b.lockc <- struct{}{}
}

func (b *SyncBroadcaster) unlock() {
	b.lockm.Unlock()
	//<-b.lockc
}
