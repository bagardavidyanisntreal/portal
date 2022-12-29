package main

import (
	"context"
	"log"
	"net/http"

	"github.com/bagardavidyanisntreal/portal/example/account"
	"github.com/bagardavidyanisntreal/portal/example/application"
	"github.com/bagardavidyanisntreal/portal/example/user"
	"github.com/bagardavidyanisntreal/portal/portal"
)

const port = ":2022"

func main() {
	ctx := context.Background()
	portalGate := portal.New(ctx)
	defer portalGate.Close()
	users := user.NewStorage(ctx, portalGate)
	accounts := account.NewStorage(ctx, portalGate)
	app := application.New(users, accounts)

	http.HandleFunc("/add/user", app.AddUser)
	http.HandleFunc("/add/account", app.AddAccount)
	log.Fatal(http.ListenAndServe(port, nil))
}
