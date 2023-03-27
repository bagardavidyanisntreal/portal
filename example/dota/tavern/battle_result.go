package tavern

import (
	"log"

	"github.com/bagardavidyanisntreal/portal/v2/example/dota/dto"
)

func NewBattlerResultHandler(tavern *Tavern) *BattleResultHandler {
	return &BattleResultHandler{tavern: tavern}
}

type BattleResultHandler struct {
	tavern *Tavern
}

func (h *BattleResultHandler) Handle(msg any) {
	battlerResult, ok := msg.(dto.BattleResult)
	if !ok || battlerResult.Empty() {
		return
	}

	coolDown, err := h.tavern.UpdateScore(battlerResult.Winner, battlerResult.Looser)
	if err != nil {
		log.Printf("[battler result handler error]: %s", err)
		return
	}

	log.Printf(
		"%s killz %s! See you %s after %v\n",
		battlerResult.Winner,
		battlerResult.Looser,
		battlerResult.Looser.Name,
		coolDown,
	)

	respawn := &dto.Respawn{Hero: battlerResult.Looser, CoolDown: coolDown}
	h.tavern.Respawn(respawn)
}
