package tavern

import (
	"fmt"
)

func (t *Tavern) PickHero(username string, hero string) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	picked, ok := t.pickByUsername[username]
	if ok {
		return fmt.Errorf("%s already picked %s", username, picked)
	}

	_, ok = t.heroesToPick[hero]
	if !ok {
		return fmt.Errorf("hero %s is not available", hero)
	}

	t.pickByUsername[username] = hero
	delete(t.heroesToPick, hero)

	return nil
}
