package tavern

import (
	"errors"
	"fmt"
	"time"

	"github.com/bagardavidyanisntreal/portal/v2/example/dota/dto"
)

func NewHeroScore(death, frags int, hero *dto.Hero) (*HeroScore, error) {
	if hero == nil {
		return nil, errors.New("hero score cannot be without a hero")
	}

	return &HeroScore{
		deaths: death,
		frags:  frags,
		hero:   hero,
	}, nil
}

type HeroScore struct {
	deaths int
	frags  int
	hero   *dto.Hero
}

func (s HeroScore) CoolDownExt() time.Duration {
	return time.Duration(s.deaths/2) * time.Second
}

func (s HeroScore) String() string {
	return fmt.Sprintf("{%s frags: %d, deaths: %d}", s.hero, s.frags, s.deaths)
}

func (s *HeroScore) AddFrag() {
	if s == nil {
		return
	}

	s.frags++
	if s.frags%3 == 0 {
		s.hero.Level++
	}
}
