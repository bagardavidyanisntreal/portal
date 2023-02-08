package portal

import (
	"context"
	"sync"
)

// Gate implementation to embed in case of distributed interfaces
// or just to import locally
type Gate interface {
	Send(msg Message)
	Await(ctx context.Context, handlers ...Handler)
}

// Handler func signature to pass through Gate.Await
type Handler interface {
	Support(msg Message) bool
	Handle(msg Message)
}

// Message data to pass through Gate
// todo add generics ?? no way cause generics are 💩
// may be in different package to choose between simple type assertion/listen all the messages
// or get typed data output for performance
type Message interface {
	Data() any
}

// Portal helps to connect services without coupling
// to pass a message use Send
// to receive a message use Await with specific handler func on it
type Portal struct {
	wg *sync.WaitGroup
	mu sync.RWMutex

	subs  []*input
	input *input
}

// New Portal constructor
// also runs inputMonitor func under the hood
func New(ctx context.Context) *Portal {
	p := &Portal{
		wg:    new(sync.WaitGroup),
		input: newInput(),
	}
	p.inputMonitor(ctx)
	return p
}
