package portal

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
			close(p.input)
			for _, sub := range p.subs {
				close(sub)
			}
			return
		case msg, open := <-p.input:
			if !open {
				return
			}
			for _, sub := range p.subs {
				send(msg, sub, p.done)
			}
		}
	}
}
