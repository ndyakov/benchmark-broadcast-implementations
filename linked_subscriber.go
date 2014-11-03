package broadcastr

type LinkedSubscriber struct {
	current chan *payload
	closed  bool
}

func newLinkedSubscriber(c chan *payload) *LinkedSubscriber {
	s := new(LinkedSubscriber)
	s.current = c
	s.closed = false
	return s
}

func (s *LinkedSubscriber) Closed() bool {
	return s.closed
}

func (s *LinkedSubscriber) Close() {
	s.closed = true
}

func (s *LinkedSubscriber) Unsubscribe() {
	s.Close()
}

func (s *LinkedSubscriber) Read() (interface{}, bool) {
	select {
	case p, ok := <-s.current:
		if !ok {
			return nil, false
		}
		value := p.value
		s.current <- p
		s.current = p.link
		return value, true
	default:
		return nil, false
	}
}

func (s *LinkedSubscriber) ReadSync() (interface{}, bool) {
	if s.closed {
		return nil, false
	}

	p, ok := <-s.current

	if !ok {
		return nil, false
	}

	value := p.value
	s.current <- p
	s.current = p.link
	return value, true
}

func (s *LinkedSubscriber) ReadInt() (int, bool) {
	v, ok := s.Read()
	if !ok {
		return 0, false
	}

	value, ok := v.(int)
	return value, ok
}

func (s *LinkedSubscriber) ReadString() (string, bool) {
	v, ok := s.Read()
	if !ok {
		return "", false
	}

	value, ok := v.(string)
	return value, ok
}
