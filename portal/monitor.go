package portal

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
		case envelope, open := <-p.input:
			if !open {
				return
			}
			p.notify(envelope.msg, envelope.destinations)
		}
	}
}

func (p *Portal) notify(msg any, destinations []chan any) {
	for _, dest := range destinations {
		select {
		case <-p.done:
			p.closeSubs()
			return
		default:
		}

		select {
		case <-p.done:
			p.closeSubs()
			return
		case dest <- msg:
		}
	}
}

func (p *Portal) closeSubs() {
	p.subsOnce.Do(func() {
		p.lock.Lock()
		defer p.lock.Unlock()
		for _, sub := range p.subs {
			close(sub)
		}
	})
}
