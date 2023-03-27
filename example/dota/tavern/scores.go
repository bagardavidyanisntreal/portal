package tavern

import "sort"

type ScoreListRequest struct {
	Hero string
}

type ScoreListResponse struct {
	Scores []*HeroScore
}

func (t *Tavern) ScoreList(req *ScoreListRequest) *ScoreListResponse {
	resp := &ScoreListResponse{}

	if req != nil && req.Hero != "" {
		t.lock.RLock()
		heroScore, ok := t.scoreByHero[req.Hero]
		t.lock.RUnlock()

		if !ok {
			return resp
		}

		resp.Scores = append(resp.Scores, heroScore)
		return resp
	}

	t.lock.RLock()
	for _, score := range t.scoreByHero {
		resp.Scores = append(resp.Scores, score)
	}
	t.lock.RUnlock()

	sort.Slice(resp.Scores, func(i, j int) bool {
		return resp.Scores[i].frags > resp.Scores[j].frags
	})

	return resp
}
