package portal

import (
	"context"
	"log"
	"sync"
)

// Gate implementation to embed in case of distributed interfaces or just to import locally
type Gate interface {
	Send(msg any)
	Subscribe(handlers ...Handler)
}

// Handler func signature to pass through Gate.Await
type Handler interface {
	Handle(msg any)
}

// New Portal constructor
func New() *Portal {
	ctx, cancel := context.WithCancel(context.Background())
	return &Portal{ctx: ctx, cancel: cancel}
}

// Portal helps to connect services without coupling
// to pass a message use Send
// to receive a message use Subscribe with specific handler func on it
type Portal struct {
	ctx    context.Context
	cancel context.CancelFunc

	subs []chan any
	wg   sync.WaitGroup
	lock sync.RWMutex
}

// Close signals about Portal working ending
func (p *Portal) Close() {
	log.Println("stopping portal...")
	p.wg.Wait()
	p.cancel()
}
