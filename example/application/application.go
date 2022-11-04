package application

import (
	"github.com/DavidBagaryan/portal/example/account"
	"github.com/DavidBagaryan/portal/example/user"
)

type (
	users interface {
		Add(name string, age uint) (*user.User, error)
		Get(userID int64) (*user.User, error)
	}
	accounts interface {
		Add(user user.User, balance int64) (*account.Account, error)
	}
)

type Application struct {
	users    users
	accounts accounts
}

func New(users users, accounts accounts) *Application {
	return &Application{
		users:    users,
		accounts: accounts,
	}
}
