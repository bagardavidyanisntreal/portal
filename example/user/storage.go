package user

import (
	"fmt"
	"sync"

	"github.com/bagardavidyanisntreal/portal/portal"
)

type portalGate interface {
	portal.Gate
}

type Storage struct {
	lock sync.RWMutex
	list map[int64]*User
	gate portalGate
}

func NewStorage(gate portalGate) *Storage {
	gate.Subscribe(&uselessHandler{})
	return &Storage{
		gate: gate,
		list: make(map[int64]*User),
	}
}

func (s *Storage) Add(name string, age uint) (*User, error) {
	user, err := New(name, age)
	if err != nil {
		return nil, err
	}
	s.gate.Send(user)
	s.lock.Lock()
	s.list[user.ID] = user
	s.lock.Unlock()
	return user, nil
}

func (s *Storage) Get(userID int64) (*User, error) {
	got, ok := s.list[userID]
	if !ok {
		return nil, fmt.Errorf("user not found: %d", userID)
	}
	return got, nil
}
