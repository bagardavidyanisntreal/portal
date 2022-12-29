package portal

import (
	"context"
	"log"
)

// inputMonitor monitors all the input data to fun out the subscribers
// NOTICE
// in future version this input inputMonitor func
// will be removed from under the hood New Portal constructor call
// to prevent context mess-up
func (p *Portal) inputMonitor(ctx context.Context) {
	const logfmt = "[portal input inputMonitor]: %s"
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Printf(logfmt, ctx.Err())
				return
			default:
				select {
				case <-ctx.Done():
					log.Printf(logfmt, ctx.Err())
					return
				case msg, open := <-p.input.hub:
					if !open {
						return
					}
					for _, sub := range p.subscriptions() {
						sub.send(msg)
					}
				}
			}
		}
	}()
}

func (p *Portal) subscriptions() []*input {
	subs := make([]*input, len(p.subs))
	p.lock.RLock()
	copy(subs, p.subs)
	p.lock.RUnlock()
	return subs
}
