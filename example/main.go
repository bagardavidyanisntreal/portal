package main

import (
	"log"
	"net/http"

	"github.com/bagardavidyanisntreal/portal/v2/example/account"
	"github.com/bagardavidyanisntreal/portal/v2/example/application"
	"github.com/bagardavidyanisntreal/portal/v2/example/user"
	"github.com/bagardavidyanisntreal/portal/v2/portal"
)

const port = ":2022"

func main() {
	portalGate := portal.New()
	defer portalGate.Close()

	users := user.NewStorage(portalGate)
	accounts := account.NewStorage(portalGate)
	app := application.New(users, accounts)

	http.HandleFunc("/add/user", app.AddUser)
	http.HandleFunc("/add/account", app.AddAccount)
	log.Fatal(http.ListenAndServe(port, nil))
}
