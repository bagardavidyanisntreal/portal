package dto

import (
	"fmt"
)

func NewHero(name string) *Hero {
	return &Hero{Name: name}
}

type Hero struct {
	Name  string
	Level int
}

func (h Hero) String() string {
	return fmt.Sprintf("%s level %d", h.Name, h.Level)
}

const (
	Axe             = "Axe"
	Mirana          = "Mirana"
	Lion            = "Lion"
	PhantomAssassin = "Phantom Assassin"
	Barathrum       = "Barathrum"
	Windranger      = "Windranger"
	Tinker          = "Tinker"
	Clockwerk       = "Clockwerk"
	Enchantress     = "Enchantress"
	Pudge           = "Pudge"
	Lina            = "Lina"
	AntiMage        = "Anti-Mage"
	Luna            = "Luna"
	Ursa            = "Ursa"
	DrowRanger      = "Drow Ranger"
	Techies         = "Techies"
)
