package runtime

// Bond ..
type Bond interface {
	Errors() chan error
	Recv() chan string
	Send() chan string
	Close() chan error
}

// Stream ..
type Stream interface {
	New() chan Message
}

// Message ..
type Message struct {
	Payload []byte
	Key     []byte
	Origin  []byte
}
