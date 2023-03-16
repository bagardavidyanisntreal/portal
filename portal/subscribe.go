package portal

// Subscribe subscribes specific handler on notification from portal input
// process runs on listener goroutine
func (p *Portal) Subscribe(handlers ...Handler) {
	if len(handlers) == 0 {
		return
	}

	subscribers := make([]chan any, len(handlers))
	for i, handler := range handlers {
		subscriber := make(chan any)
		subscribers[i] = subscriber

		go p.listen(subscriber, handler)
	}

	p.lock.Lock()
	p.subs = append(p.subs, subscribers...)
	p.lock.Unlock()
}

func (p *Portal) listen(subscription chan any, handler Handler) {
	for {
		if err := p.ctx.Err(); err != nil {
			return
		}

		select {
		case <-p.ctx.Done():
			return
		case msg, open := <-subscription:
			if !open {
				return
			}
			handler.Handle(msg)
			p.wg.Done()
		}
	}
}
