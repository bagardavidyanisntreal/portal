package portal

import "log"

// Close ends Portal working closing input channel and all the subscriptions
func (p *Portal) Close() {
	log.Println("stopping portal...")
	p.done <- struct{}{}
}
