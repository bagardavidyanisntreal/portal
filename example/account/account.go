package account

import (
	"errors"
	"fmt"

	"github.com/DavidBagaryan/portal/example/user"
)

type Account struct {
	balance    int64
	privileges Privileges
	user       *user.User
}

func New(user *user.User, opts ...AccOpt) (*Account, error) {
	if user == nil {
		return nil, errors.New("cannot creat acc from undefined user")
	}
	acc := &Account{
		user: user,
	}
	var err error
	for _, opt := range opts {
		acc, err = opt(acc)
		if err != nil {
			return nil, err
		}
	}
	return acc, nil
}

func (a Account) String() string {
	return fmt.Sprintf(
		`{"user": %s, "balance": %d, "priveleges": %d}`,
		a.user,
		a.balance,
		a.privileges,
	)
}

type Privileges int

const (
	Intern Privileges = iota + 1
	Junior
	Middle
	Senior
)

func (a Account) IsIntern() bool {
	return !(a.privileges > Intern)
}

func (a Account) IsJunior() bool {
	return !(a.privileges > Junior)
}

func (a Account) IsMiddle() bool {
	return !(a.privileges > Middle)
}

func (a Account) IsSenior() bool {
	return !(a.privileges > Senior)
}

type AccOpt func(acc *Account) (*Account, error)

func WithBalance(balance int64) AccOpt {
	return func(acc *Account) (*Account, error) {
		if acc.IsIntern() && balance > 100 {
			return nil, errors.New("intern cannot get more then 100")
		}
		if acc.IsJunior() && balance > 300 {
			return nil, errors.New("junior cannot get more then 300")
		}
		if acc.IsMiddle() && balance > 500 {
			return nil, errors.New("middle cannot get more then 500")
		}
		if acc.IsSenior() && balance > 700 {
			return nil, errors.New("senior cannot get more then 700")
		}
		acc.balance = balance
		return acc, nil
	}
}

func WithPrivileges(privileges Privileges) AccOpt {
	return func(acc *Account) (*Account, error) {
		acc.privileges = privileges
		return acc, nil
	}
}
