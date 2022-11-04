package account

import (
	"fmt"
	"sync"

	"github.com/DavidBagaryan/portal"
	"github.com/DavidBagaryan/portal/example/user"
)

type Storage struct {
	lock sync.RWMutex
	list map[int64]*Account
	gate portal.Gate
}

func NewStorage(gate portal.Gate) *Storage {
	s := &Storage{
		gate: gate,
		list: make(map[int64]*Account),
	}
	gate.Await(func(msg any) {
		u, ok := msg.(user.User)
		if !ok {
			return
		}
		_, err := s.Add(u, 100)
		if err != nil {
			fmt.Printf("an err occured on creating account for user %s", u)
		}
	})
	return s
}

func (s *Storage) Add(user user.User, balance int64) (*Account, error) {
	existed, _ := s.Get(user.GetID())
	if existed != nil {
		return nil, fmt.Errorf("account already exists, %s", user)
	}
	acc, err := New(&user, WithBalance(balance), WithPrivileges(Privileges(1)))
	if err != nil {
		return nil, err
	}
	s.lock.Lock()
	s.list[user.GetID()] = acc
	s.lock.Unlock()
	s.gate.Send(fmt.Sprintf("created new account %s", acc))
	return acc, nil
}

func (s *Storage) Get(userID int64) (*Account, error) {
	got, ok := s.list[userID]
	if !ok {
		return nil, fmt.Errorf("account not found by userID: %d", userID)
	}
	return got, nil
}
