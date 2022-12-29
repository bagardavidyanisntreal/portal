package portal

// Send sends message on input channel, which fans-out it on subscriptions after
// each subscription handler decides itself whether to process the received message or not
func (p *Portal) Send(msg Message) {
	subs := p.subscriptions()
	p.wg.Add(len(subs))
	go p.input.send(msg)
}
