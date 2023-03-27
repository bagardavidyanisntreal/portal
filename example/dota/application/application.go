package application

import "github.com/bagardavidyanisntreal/portal/v2/example/dota/tavern"

type (
	portal interface {
		Send(any)
	}
	tavernSrv interface {
		PickHero(name string, hero string) error
		ScoreList(req *tavern.ScoreListRequest) *tavern.ScoreListResponse
	}
)

func New(portal portal, tavern tavernSrv) *Application {
	return &Application{portal: portal, tavern: tavern}
}

type Application struct {
	portal portal
	tavern tavernSrv
}
