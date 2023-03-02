package portal

// Send sends message on input channel, which fans-out it on subscriptions
// after each subscription handler decides itself whether to process the received message or not
func (p *Portal) Send(msg any) {
	send(msg, p.input, p.done)
}

func send(msg any, input chan any, done chan struct{}) {
	select {
	case <-done:
		return
	case input <- msg:
	}
}
