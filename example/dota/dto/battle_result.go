package dto

type BattleResult struct {
	Winner, Looser *Hero
}

func (r BattleResult) Empty() bool {
	return r.Winner == nil || r.Looser == nil
}
