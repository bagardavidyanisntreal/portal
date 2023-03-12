package portal

import "sync"

func (p *Portal) monitor() {
	defer p.closeSubs()
	for {
		if err := p.ctx.Err(); err != nil {
			return
		}

		select {
		case <-p.ctx.Done():
			return
		case envelope, open := <-p.input:
			if !open {
				return
			}
			go func() {
				for _, dest := range envelope.destinations {
					dest <- envelope.msg
				}
			}()
		}
	}
}

var subsOnce = sync.Once{}

func (p *Portal) closeSubs() {
	subsOnce.Do(func() {
		p.lock.Lock()
		defer p.lock.Unlock()

		for _, sub := range p.subs {
			close(sub)
		}
	})
}
