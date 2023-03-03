package portal

// Send sends message on input channel, which fans-out it on subscriptions
// after each subscription handler decides itself whether to process the received message or not
func (p *Portal) Send(msg any) {
	select {
	case <-p.done:
		p.closeInput()
		return
	default:
	}

	select {
	case <-p.done:
		p.closeInput()
		return
	case p.input <- msg:
	}
}

func (p *Portal) closeInput() {
	p.inpOnce.Do(func() {
		close(p.input)
	})
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
