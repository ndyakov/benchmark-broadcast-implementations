package broadcastr

type AsyncSubscriber struct {
	id        int
	inbox     chan interface{}
	bufferCap int
	closed    bool
}

func newAsyncSubscriber(bufferCap, id int) *AsyncSubscriber {
	s := new(AsyncSubscriber)
	s.id = id
	s.bufferCap = bufferCap
	s.inbox = make(chan interface{}, bufferCap)
	s.closed = false
	return s
}

func (s *AsyncSubscriber) Inbox() <-chan interface{} {
	return s.inbox
}

func (s *AsyncSubscriber) BufferCap() int {
	return s.bufferCap
}

func (s *AsyncSubscriber) Closed() bool {
	return s.closed
}

func (s *AsyncSubscriber) Close() {
	s.closed = true
	close(s.inbox)
}

func (s *AsyncSubscriber) Id() int {
	return s.id
}

func (s *AsyncSubscriber) Unsubscribe() {
	s.Close()
}

func (s *AsyncSubscriber) Read() (interface{}, bool) {
	select {
	case v, ok := <-s.inbox:
		return v, ok
	default:
		return nil, false
	}
}

func (s *AsyncSubscriber) ReadSync() (interface{}, bool) {
	v, ok := <-s.inbox
	return v, ok
}

func (s *AsyncSubscriber) ReadInt() (int, bool) {
	v, ok := s.Read()
	if !ok {
		return 0, false
	}

	value, ok := v.(int)
	return value, ok
}

func (s *AsyncSubscriber) ReadString() (string, bool) {
	v, ok := s.Read()
	if !ok {
		return "", false
	}

	value, ok := v.(string)
	return value, ok
}

func (s *AsyncSubscriber) flushBuffer() {
	for {
		select {
		case _, ok := <-s.inbox:
			if !ok {
				return
			}
			continue
		default:
			return
		}
	}
}
