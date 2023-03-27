package arena

import (
	"log"

	"github.com/bagardavidyanisntreal/portal/v2/example/dota/dto"
)

func NewHeroRebornHandler(arena *Arena) *HeroRebornHandler {
	return &HeroRebornHandler{arena: arena}
}

type HeroRebornHandler struct {
	arena *Arena
}

func (h *HeroRebornHandler) Handle(msg any) {
	heroSelected, ok := msg.(dto.HeroReborn)
	if !ok || heroSelected.Hero == nil {
		return
	}

	log.Printf("Welcome back %s and Goodluck!\n", heroSelected.Hero)

	h.arena.PushToBattle(heroSelected.Hero)
}
