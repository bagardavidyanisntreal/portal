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
	defer p.lock.Unlock()
	p.subs = append(p.subs, subscribers...)
}

func (p *Portal) listen(subscription chan any, handler Handler) {
	for {
		select {
		case <-p.done:
			return
		default:
		}

		select {
		case <-p.done:
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
