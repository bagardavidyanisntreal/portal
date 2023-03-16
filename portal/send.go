package portal

import "sync"

// Send sends message for all the existing subscribers at the moment
// and then each subscription handler decides itself whether to process the received message or not
func (p *Portal) Send(msg any) {
	if err := p.ctx.Err(); err != nil {
		p.closeSubs()
		return
	}

	subscriptions := p.subscriptions()
	p.wg.Add(len(subscriptions))

	for _, sub := range subscriptions {
		sub := sub
		go func() {
			sub <- msg
		}()
	}
}

var subsOnce = sync.Once{}

func (p *Portal) closeSubs() {
	subsOnce.Do(func() {
		p.lock.Lock()
		for _, sub := range p.subs {
			close(sub)
		}
		p.lock.Unlock()
	})
}

func (p *Portal) subscriptions() []chan any {
	p.lock.RLock()
	subs := make([]chan any, len(p.subs))
	copy(subs, p.subs)
	p.lock.RUnlock()

	return subs
}
