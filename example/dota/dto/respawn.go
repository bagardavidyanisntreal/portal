package dto

import "time"

type Respawn struct {
	Hero     *Hero
	CoolDown time.Duration
}
