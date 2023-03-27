package tavern

import (
	"sync"

	"github.com/bagardavidyanisntreal/portal/v2/example/dota/dto"
)

type portal interface {
	Send(msg any)
}

func New(heroesToPick map[string]struct{}, portal portal) *Tavern {
	t := &Tavern{
		heroesToPick:   heroesToPick,
		pickByUsername: make(map[string]string),
		scoreByHero:    make(map[string]*HeroScore),
		respawn:        make(chan dto.Respawn),
		portal:         portal,
	}

	go t.waitForRespawn()
	return t
}

type Tavern struct {
	heroesToPick   map[string]struct{}
	pickByUsername map[string]string
	scoreByHero    map[string]*HeroScore
	respawn        chan dto.Respawn
	portal         portal
	lock           sync.RWMutex
}
