package arena

import (
	"log"

	"github.com/bagardavidyanisntreal/portal/v2/example/dota/dto"
)

func NewHeroSelectedHandler(arena *Arena) *HeroSelectedHandler {
	return &HeroSelectedHandler{arena: arena}
}

type HeroSelectedHandler struct {
	arena *Arena
}

func (h *HeroSelectedHandler) Handle(msg any) {
	heroSelected, ok := msg.(dto.HeroSelected)
	if !ok || heroSelected.Hero == nil || heroSelected.Username == "" {
		return
	}

	log.Printf(
		"A new hero %s selected by %s. Let's see what will happend...\n",
		heroSelected.Hero.Name,
		heroSelected.Username,
	)

	h.arena.PushToBattle(heroSelected.Hero)
}
