package broadcastr

type SyncSubscriber struct {
	id        int
	inbox     chan interface{}
	bufferCap int
	closed    bool
}

func newSyncSubscriber(bufferCap, id int) *SyncSubscriber {
	s := new(SyncSubscriber)
	s.id = id
	s.bufferCap = bufferCap
	s.inbox = make(chan interface{}, bufferCap)
	s.closed = false
	return s
}

func (s *SyncSubscriber) Inbox() <-chan interface{} {
	return s.inbox
}

func (s *SyncSubscriber) BufferCap() int {
	return s.bufferCap
}

func (s *SyncSubscriber) Closed() bool {
	return s.closed
}

func (s *SyncSubscriber) Close() {
	s.closed = true
	close(s.inbox)
}

func (s *SyncSubscriber) Id() int {
	return s.id
}

func (s *SyncSubscriber) Unsubscribe() {
	s.Close()
}

func (s *SyncSubscriber) Read() (interface{}, bool) {
	select {
	case v, ok := <-s.inbox:
		return v, ok
	default:
		return nil, false
	}
}

func (s *SyncSubscriber) ReadSync() (interface{}, bool) {
	v, ok := <-s.inbox
	return v, ok
}

func (s *SyncSubscriber) ReadInt() (int, bool) {
	v, ok := s.Read()
	if !ok {
		return 0, false
	}

	value, ok := v.(int)
	return value, ok
}

func (s *SyncSubscriber) ReadString() (string, bool) {
	v, ok := s.Read()
	if !ok {
		return "", false
	}

	value, ok := v.(string)
	return value, ok
}

func (s *SyncSubscriber) flushBuffer() {
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
