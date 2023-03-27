package tavern

import (
	"time"

	"github.com/bagardavidyanisntreal/portal/v2/example/dota/dto"
)

func (t *Tavern) waitForRespawn() {
	for {
		respawn := <-t.respawn

		go func() {
			<-time.After(respawn.CoolDown)
			t.portal.Send(dto.HeroReborn{Hero: respawn.Hero})
		}()
	}
}

func (t *Tavern) Respawn(respawn *dto.Respawn) {
	if respawn == nil {
		return
	}

	// don't forget check if we still can pass hero to respawn

	t.respawn <- *respawn
}

const minCoolDown = 3 * time.Second

func (t *Tavern) UpdateScore(winner, looser *dto.Hero) (time.Duration, error) {
	t.lock.Lock()
	defer t.lock.Unlock()

	looserScore, ok := t.scoreByHero[looser.Name]
	if !ok {
		score, err := NewHeroScore(1, 0, looser)
		if err != nil {
			return 0, err
		}

		t.scoreByHero[looser.Name] = score
		looserScore = score
	} else {
		looserScore.deaths++
	}

	winnerScore, ok := t.scoreByHero[winner.Name]
	if !ok {
		newWinnerScore, err := NewHeroScore(0, 1, winner)
		if err != nil {
			return 0, err
		}

		t.scoreByHero[winner.Name] = newWinnerScore
	} else {
		winnerScore.AddFrag()
	}

	return minCoolDown + looserScore.CoolDownExt(), nil
}
