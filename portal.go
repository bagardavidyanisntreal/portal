package portal

import (
	"context"
	"log"
	"sync"
)

// Gate implementation to embed in case of distributed interfaces
// or just to import locally
type Gate interface {
	Send(msg Message)
	Await(ctx context.Context, handlers ...Handler)
}

// Handler func signature to pass through Gate.Await
type Handler interface {
	Support(msg Message) bool
	Handle(msg Message)
}

// Message data to pass through Gate
// todo add generics ?? no way cause generics are 💩
// may be in different package to choose between simple type assertion/listen all the messages
// or get typed data output for performance
type Message interface {
	Data() any
}

// Portal helps to connect services without coupling
// to pass a message use Send
// to receive a message use Await with specific handler func on it
type Portal struct {
	wg    sync.WaitGroup
	lock  sync.RWMutex
	subs  []chan Message
	input chan Message
}

const logfmt = "[PORTAL]: %s"

// New Portal constructor
// also runs monitor func under the hood
func New(ctx context.Context) *Portal {
	return (&Portal{
		input: make(chan Message),
	}).monitor(ctx)
}

func (p *Portal) monitor(ctx context.Context) *Portal {
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
				case inp := <-p.input:
					for _, sub := range p.subscriptions() {
						p.wg.Add(1)
						sub <- inp
					}
				}
			}
		}
	}()
	return p
}

func (p *Portal) subscriptions() []chan Message {
	subs := make([]chan Message, len(p.subs))
	p.lock.RLock()
	copy(subs, p.subs)
	p.lock.RUnlock()
	return subs
}

// Send sends message on input channel, which fans-out it on subscriptions after
// each subscription handler decides for itself whether to process the received message or not
func (p *Portal) Send(msg Message) {
	go func() {
		p.input <- msg
	}()
}

// Await subscribes specific handler on notification from portal input
// process runs on listener goroutine
func (p *Portal) Await(ctx context.Context, handlers ...Handler) {
	subscription := make(chan Message)
	p.lock.Lock()
	p.subs = append(p.subs, subscription)
	p.lock.Unlock()
	for _, handler := range handlers {
		p.listen(ctx, subscription, handler)
	}
}

func (p *Portal) listen(ctx context.Context, subscription <-chan Message, handler Handler) {
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
				case msg, open := <-subscription:
					if !open {
						return
					}
					if handler.Support(msg) {
						handler.Handle(msg)
					}
					p.wg.Done()
				}
			}
		}
	}()
}

// Close ends Portal working closing input channel and all the subscriptions
func (p *Portal) Close() {
	p.wg.Wait()
	p.lock.Lock()
	defer p.lock.Unlock()
	close(p.input)
	for _, sub := range p.subs {
		close(sub)
	}
}
