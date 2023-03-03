package portal

import (
	"log"
	"sync"
)

// Gate implementation to embed in case of distributed interfaces
// or just to import locally
type Gate interface {
	Send(msg any)
	Subscribe(handlers ...Handler)
}

// Handler func signature to pass through Gate.Await
type Handler interface {
	Handle(msg any)
}

// Portal helps to connect services without coupling
// to pass a message use Send
// to receive a message use Subscribe with specific handler func on it
type Portal struct {
	done  chan struct{}
	input chan any
	subs  []chan any

	inputCloser sync.Once
	subsCloser  sync.Once
}

// New Portal constructor
// also runs monitor for input
func New() *Portal {
	p := &Portal{
		done:  make(chan struct{}),
		input: make(chan any),
	}

	go p.monitor()
	return p
}

func (p *Portal) monitor() {
	for {
		select {
		case <-p.done:
			return
		default:
		}

		select {
		case <-p.done:
			return
		case msg, open := <-p.input:
			if !open {
				return
			}
			p.notify(msg)
		}
	}
}

// Close signals about Portal working ending
func (p *Portal) Close() {
	log.Println("stopping portal...")
	close(p.done)
}
