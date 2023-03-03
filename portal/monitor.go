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
		case msg, open := <-p.input:
			if !open {
				return
			}
			p.notify(msg)
		}
	}
}

func (p *Portal) notify(msg any) {
	p.subsLock.Lock()
	defer p.subsLock.Unlock()
	for _, sub := range p.subs {
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
		case sub <- msg:
		}
	}
}

func (p *Portal) closeSubs() {
	p.subsOnce.Do(func() {
		p.subsLock.Lock()
		defer p.subsLock.Unlock()
		for _, sub := range p.subs {
			close(sub)
		}
	})
}
