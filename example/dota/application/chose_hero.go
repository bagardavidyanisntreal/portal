package application

import (
	"fmt"
	"net/http"

	"github.com/bagardavidyanisntreal/portal/v2/example/dota/dto"
)

func (a *Application) ChoseHero(writer http.ResponseWriter, request *http.Request) {
	username := request.URL.Query().Get("username")
	if username == "" {
		writer.WriteHeader(http.StatusBadRequest)
		response400(writer, "username cannot be empty")
		return
	}

	hero := request.URL.Query().Get("hero")
	if hero == "" {
		response400(writer, "hero is not selected")
		return
	}

	err := a.tavern.PickHero(username, hero)
	if err != nil {
		response400(writer, err.Error())
		return
	}

	a.portal.Send(dto.HeroSelected{Username: username, Hero: dto.NewHero(hero)})
	response200(writer, fmt.Sprintf("hero %s selected by %s!", hero, username))
}
