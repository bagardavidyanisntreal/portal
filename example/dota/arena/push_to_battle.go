package arena

import "github.com/bagardavidyanisntreal/portal/v2/example/dota/dto"

func (a *Arena) PushToBattle(hero *dto.Hero) {
	if hero == nil {
		return
	}

	a.lock.Lock()
	defer a.lock.Unlock()

	a.heroes = append(a.heroes, hero)
}
