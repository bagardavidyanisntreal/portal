package account

import (
	"context"
	"fmt"
	"sync"

	"github.com/bagardavidyanisntreal/portal/example/user"
	"github.com/bagardavidyanisntreal/portal/portal"
)

type Storage struct {
	lock sync.RWMutex
	list map[int64]*Account
	gate portal.Gate
}

func NewStorage(ctx context.Context, gate portal.Gate) *Storage {
	srg := &Storage{
		gate: gate,
		list: make(map[int64]*Account),
	}
	gate.Await(ctx, &createAccountForNewUser{storage: srg})
	return srg
}

func (s *Storage) Add(user user.User, balance int64) (*Account, error) {
	existed, _ := s.Get(user.GetID())
	if existed != nil {
		return nil, fmt.Errorf("account already exists, %s", user)
	}
	account, err := New(&user, WithBalance(balance), WithPrivileges(Privileges(1)))
	if err != nil {
		return nil, err
	}
	s.lock.Lock()
	s.list[user.GetID()] = account
	s.lock.Unlock()
	s.gate.Send(account.CreatedNotify())
	return account, nil
}

func (s *Storage) Get(userID int64) (*Account, error) {
	got, ok := s.list[userID]
	if !ok {
		return nil, fmt.Errorf("account not found by userID: %d", userID)
	}
	return got, nil
}
