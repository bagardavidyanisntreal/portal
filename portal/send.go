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
