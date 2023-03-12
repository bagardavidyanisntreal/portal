package portal

import "sync"

// Send sends message on input channel, which fans-out it on subscriptions
// after each subscription handler decides itself whether to process the received message or not
func (p *Portal) Send(msg any) {
	destinations := p.destinations()
	p.wg.Add(len(destinations))

	envelope := envelopeMsg{msg: msg, destinations: destinations}

	if err := p.ctx.Err(); err != nil {
		p.closeInput(len(destinations))
		return
	}

	p.input <- envelope
}

func (p *Portal) destinations() []chan any {
	p.lock.RLock()
	subs := make([]chan any, len(p.subs))
	copy(subs, p.subs)
	p.lock.RUnlock()

	return subs
}

var inputOnce = sync.Once{}

func (p *Portal) closeInput(destinationCount int) {
	inputOnce.Do(func() {
		p.wg.Add(-1 * destinationCount)
		close(p.input)
	})
}
