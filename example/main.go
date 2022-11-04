package main

import (
	"log"
	"net/http"

	"github.com/DavidBagaryan/portal"
	"github.com/DavidBagaryan/portal/example/account"
	"github.com/DavidBagaryan/portal/example/application"
	"github.com/DavidBagaryan/portal/example/user"
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
