package broadcastr

type Broadcaster interface {
	Send(interface{}) error
	Subscribe() Subscriber
	Close()
}

func NewBroadcaster() Broadcaster {
	return NewAsyncBroadcaster()
}
