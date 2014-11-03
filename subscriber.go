package broadcastr

type Subscriber interface {
	Read() (interface{}, bool)
	ReadSync() (interface{}, bool)
	ReadInt() (int, bool)
	ReadString() (string, bool)
	Closed() bool
	Close()
	Unsubscribe()
}
