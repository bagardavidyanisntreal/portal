package main

import (
	"log"
	"net/http"

	"github.com/bagardavidyanisntreal/portal/v2/example/dota/application"
	"github.com/bagardavidyanisntreal/portal/v2/example/dota/arena"
	"github.com/bagardavidyanisntreal/portal/v2/example/dota/dto"
	"github.com/bagardavidyanisntreal/portal/v2/example/dota/tavern"
	"github.com/bagardavidyanisntreal/portal/v2/portal"
)

const port = ":2022"

func main() {
	portalGate := portal.New()
	defer portalGate.Close()

	heroTavern := tavern.New(allHeroes(), portalGate)
	battleResultHandler := tavern.NewBattlerResultHandler(heroTavern)

	battleArena := arena.New(portalGate)
	heroSelectedHandler := arena.NewHeroSelectedHandler(battleArena)
	heroRebornHandler := arena.NewHeroRebornHandler(battleArena)

	portalGate.Subscribe(heroSelectedHandler, battleResultHandler, heroRebornHandler)

	app := application.New(portalGate, heroTavern)

	http.HandleFunc("/hero/choose/", app.ChoseHero)
	http.HandleFunc("/hero/scores/", app.Scores)

	log.Fatal(http.ListenAndServe(port, nil))
}

func allHeroes() map[string]struct{} {
	return map[string]struct{}{
		dto.Axe:             {},
		dto.Mirana:          {},
		dto.Lion:            {},
		dto.PhantomAssassin: {},
		dto.Barathrum:       {},
		dto.Windranger:      {},
		dto.Tinker:          {},
		dto.Clockwerk:       {},
		dto.Enchantress:     {},
		dto.Pudge:           {},
		dto.Lina:            {},
		dto.AntiMage:        {},
		dto.Luna:            {},
		dto.Ursa:            {},
		dto.DrowRanger:      {},
		dto.Techies:         {},
	}
}
