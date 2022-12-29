package portal

import (
	"log"
	"sync"
	"sync/atomic"
)

type input struct {
	hub    chan Message
	lock   sync.Mutex
	closed atomic.Bool
}

func newInput() *input {
	return &input{
		hub: make(chan Message),
	}
}

func (i *input) send(msg Message) {
	const logfmt = "[portal input send]: %s"

	if i.closed.Load() {
		log.Printf(logfmt, "exiting, input is closed")
		return
	}

	i.lock.Lock()
	i.hub <- msg
	i.lock.Unlock()
}

func (i *input) close() {
	const logfmt = "[input close]: %s"

	log.Printf(logfmt, "closing...")

	i.lock.Lock()
	defer i.lock.Unlock()

	i.closed.Store(true)
	log.Printf(logfmt, "closed")
}
