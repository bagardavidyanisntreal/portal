package portal

// Send sends message on input channel, which fans-out it on subscriptions
// after each subscription handler decides itself whether to process the received message or not
func (p *Portal) Send(msg any) {
	destinations, waiting := p.destinations()
	envelope := envelopeMsg{
		msg:          msg,
		destinations: destinations,
	}

	select {
	case <-p.done:
		p.closeInput(waiting)
		return
	default:
	}

	select {
	case <-p.done:
		p.closeInput(waiting)
		return
	case p.input <- envelope:
	}
}

func (p *Portal) destinations() ([]chan any, int) {
	p.lock.RLock()
	waiting := len(p.subs)
	subs := make([]chan any, waiting)
	copy(subs, p.subs)
	p.lock.RUnlock()

	p.wg.Add(waiting)
	return subs, waiting
}

func (p *Portal) closeInput(waiting int) {
	p.inputOnce.Do(func() {
		p.wg.Add(-1 * waiting)
		close(p.input)
	})
}
