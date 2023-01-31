package portal

import (
	"context"
	"log"
)

// Await subscribes specific handler on notification from portal input
// process runs on listener goroutine
func (p *Portal) Await(ctx context.Context, handlers ...Handler) {
	if len(handlers) == 0 {
		return
	}

	var newSubs []*input
	for _, handler := range handlers {
		newSub := newInput()
		newSubs = append(newSubs, newSub)
		p.listen(ctx, newSub, handler)
	}

	p.lock.Lock()
	defer p.lock.Unlock()
	p.subs = append(p.subs, newSubs...)
}

func (p *Portal) listen(ctx context.Context, subscription *input, handler Handler) {
	const logfmt = "[portal await listener]: %s"
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
				case msg, open := <-subscription.hub:
					if !open {
						return
					}
					if handler.Support(msg) {
						handler.Handle(msg)
						p.wg.Done()
					}
				}
			}
		}
	}()
}
