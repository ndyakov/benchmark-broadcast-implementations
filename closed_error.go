package broadcastr

type ClosedError struct{}

func (c ClosedError) Error() string {
	return "This broadcaster is closed!"
}

func (c ClosedError) String() string {
	return c.Error()
}

func newClosedError() error {
	return new(ClosedError)
}
