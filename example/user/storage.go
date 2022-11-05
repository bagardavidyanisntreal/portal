package user

import (
	"context"
	"fmt"
	"sync"

	"github.com/bagardavidyanisntreal/portal"
)

type portalGate interface {
	portal.Gate
}

type Storage struct {
	lock sync.RWMutex
	list map[int64]*User
	gate portalGate
}

func NewStorage(ctx context.Context, gate portalGate) *Storage {
	gate.Await(ctx, &uselessHandler{})
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
	s.gate.Send(user.CreatedMessage())
	s.lock.Lock()
	s.list[user.id] = user
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
