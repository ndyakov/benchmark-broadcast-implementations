package broadcastr

type payload struct {
	link  chan *payload
	value interface{}
}

type LinkedBroadcaster struct {
	closed    bool
	input     chan interface{}
	subscribe chan (chan (chan *payload))
	lockc     chan struct{}
}

func NewLinkedBroadcaster() *LinkedBroadcaster {
	b := new(LinkedBroadcaster)
	b.input = make(chan interface{})
	b.lockc = make(chan struct{}, 1)
	b.subscribe = make(chan (chan (chan *payload)))
	b.closed = false
	go b.listen()
	return b
}

func (b *LinkedBroadcaster) listen() {
	current := make(chan *payload, 1)
	for {
		select {
		case v, ok := <-b.input:
			if !ok {
				close(current)
				return
			}
			link := make(chan *payload, 1)
			p := new(payload)
			p.link = link
			p.value = v
			current <- p
			current = link
		case s, ok := <-b.subscribe:
			if !ok {
				close(current)
				return
			}
			s <- current
		}
	}
}

func (b *LinkedBroadcaster) Send(v interface{}) (err error) {
	b.lock()
	defer b.unlock()
	if b.closed {
		return newClosedError()
	}

	b.input <- v
	return nil
}

func (b *LinkedBroadcaster) Subscribe() Subscriber {
	c := make(chan chan *payload)
	b.subscribe <- c
	return newLinkedSubscriber(<-c)
}

func (b *LinkedBroadcaster) Close() {
	b.lock()
	defer b.unlock()
	b.closed = true
	close(b.input)
	close(b.subscribe)
}

func (b *LinkedBroadcaster) lock() {
	b.lockc <- struct{}{}
}

func (b *LinkedBroadcaster) unlock() {
	<-b.lockc
}
