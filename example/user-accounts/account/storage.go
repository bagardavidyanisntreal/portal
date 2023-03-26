package account

import (
	"fmt"
	"sync"

	"github.com/bagardavidyanisntreal/portal/v2/example/user-accounts/user"
	"github.com/bagardavidyanisntreal/portal/v2/portal"
)

type Storage struct {
	lock sync.RWMutex
	list map[int64]*Account
	gate portal.Gate
}

func NewStorage(gate portal.Gate) *Storage {
	srg := &Storage{
		gate: gate,
		list: make(map[int64]*Account),
	}
	gate.Subscribe(&createAccountForNewUser{storage: srg})
	return srg
}

func (s *Storage) Add(user *user.User, balance int64) (*Account, error) {
	existed, _ := s.Get(user.ID)
	if existed != nil {
		return nil, fmt.Errorf("account already exists, %v\n", user)
	}
	account, err := New(user, WithBalance(balance), WithPrivileges(Privileges(1)))
	if err != nil {
		return nil, err
	}
	s.lock.Lock()
	s.list[user.ID] = account
	s.lock.Unlock()
	s.gate.Send(account)
	return account, nil
}

func (s *Storage) Get(userID int64) (*Account, error) {
	s.lock.RLock()
	got, ok := s.list[userID]
	s.lock.RUnlock()
	if !ok {
		return nil, fmt.Errorf("account not found by userID: %d", userID)
	}
	return got, nil
}
