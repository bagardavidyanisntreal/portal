package application

import (
	"net/http"

	"github.com/bagardavidyanisntreal/portal/v2/example/dota/tavern"
)

func (a *Application) Scores(writer http.ResponseWriter, request *http.Request) {
	tavernReq := &tavern.ScoreListRequest{Hero: request.URL.Query().Get("hero")}
	scores := a.tavern.ScoreList(tavernReq)

	if len(scores.Scores) == 0 {
		response404(writer)
		return
	}

	response200(writer, scores)
}
