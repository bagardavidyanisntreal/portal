package portal

import (
	"sync"
)

// Gate implementation to embed in case of distributed interfaces
// or just to import locally
type Gate interface {
	Send(msg any)
	Await(handler func(msg any))
}

// Portal helps to connect services without coupling
// to pass a message use Send
// to receive a message use Await with specific handler func on it
type Portal struct {
	wg    sync.WaitGroup
	lock  sync.RWMutex
	subs  []chan any
	input chan any
}

// New Portal constructor
// also runs monitor func under the hood
func New() *Portal {
	return (&Portal{
		input: make(chan any),
	}).monitor()
}

func (b *Portal) monitor() *Portal {
	go func() {
		for {
			select {
			case inp := <-b.input:
				for _, sub := range b.subscriptions() {
					b.wg.Add(1)
					sub <- inp
				}
			}
		}
	}()
	return b
}

func (b *Portal) subscriptions() []chan any {
	subs := make([]chan any, len(b.subs))
	b.lock.RLock()
	copy(subs, b.subs)
	b.lock.RUnlock()
	return subs
}

// Send sends message on input channel, which fans-out it on subscriptions after
// each subscription handler decides for itself whether to process the received message or not
func (b *Portal) Send(msg any) {
	go func() {
		b.input <- msg
	}()
}

// Await subscribes specific handler on notification from portal input
// process runs on listener goroutine
func (b *Portal) Await(handler func(msg any)) {
	subscription := make(chan any)
	b.lock.Lock()
	b.subs = append(b.subs, subscription)
	b.lock.Unlock()
	b.listen(subscription, handler)
}

func (b *Portal) listen(sub <-chan any, handler func(msg any)) {
	go func() {
		for {
			select {
			case msg, open := <-sub:
				if !open {
					return
				}
				handler(msg)
				b.wg.Done()
			}
		}
	}()
}

// Close ends Portal working closing input channel and all the subscriptions
func (b *Portal) Close() {
	b.wg.Wait()
	b.lock.Lock()
	defer b.lock.Unlock()
	close(b.input)
	for _, sub := range b.subs {
		close(sub)
	}
}
