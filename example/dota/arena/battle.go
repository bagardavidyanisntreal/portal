package arena

import (
	"math/rand"
	"time"

	"github.com/bagardavidyanisntreal/portal/v2/example/dota/dto"
)

const (
	minSeconds = 2
	maxSeconds = 5
)

func (a *Arena) battle() {
	for {
		rand.Seed(time.Now().UnixNano())
		rndBattleRound := minSeconds + rand.Intn(maxSeconds-minSeconds+1)

		<-time.After(time.Duration(rndBattleRound) * time.Second)

		winner, looser := a.toTavern()
		if winner == nil || looser == nil {
			continue
		}

		a.portal.Send(dto.BattleResult{Winner: winner, Looser: looser})
	}
}

func (a *Arena) toTavern() (winner *dto.Hero, looser *dto.Hero) {
	a.lock.Lock()
	defer a.lock.Unlock()

	if len(a.heroes) == 0 || len(a.heroes) == 1 {
		return nil, nil
	}

	rand.Seed(time.Now().UnixNano())
	winnerID := rand.Int() % len(a.heroes)
	looserID := rand.Int() % len(a.heroes)

	if winnerID == looserID {
		return nil, nil
	}

	winner, looser = a.heroes[winnerID], a.heroes[looserID]

	rand.Seed(time.Now().UnixNano())
	if rand.Intn(2) == 1 {
		winner, looser, looserID = looser, winner, winnerID
	}

	if looserID == len(a.heroes)-1 {
		a.heroes = append(a.heroes[:looserID])
	} else {
		a.heroes = append(a.heroes[:looserID], a.heroes[looserID+1:]...)
	}

	if winner.Level < looser.Level {
		winner.Level++
	}

	return winner, looser
}
