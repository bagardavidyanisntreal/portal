package arena

import (
	"sync"

	"github.com/bagardavidyanisntreal/portal/v2/example/dota/dto"
)

type portalGate interface {
	Send(msg any)
}

func New(portal portalGate) *Arena {
	a := &Arena{portal: portal}
	go a.battle()
	return a
}

type Arena struct {
	heroes []*dto.Hero
	portal portalGate
	lock   sync.Mutex
}
