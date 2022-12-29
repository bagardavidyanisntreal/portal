package portal

import "log"

// Close ends Portal working closing input channel and all the subscriptions
func (p *Portal) Close() {
	const logfmt = "[portal close]: %s"

	log.Printf(logfmt, "stopping...")
	log.Printf(logfmt, "waiting handler jobs to be done...")
	p.wg.Wait()
	p.input.close()

	log.Printf(logfmt, "closing subscriber channels...")
	p.lock.Lock()
	defer p.lock.Unlock()
	for _, sub := range p.subs {
		sub.close()
	}

	log.Printf(logfmt, "gracefully stopped!")
}
