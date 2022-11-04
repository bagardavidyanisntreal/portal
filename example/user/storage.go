package user

import (
	"fmt"
	"sync"
	"time"

	"github.com/DavidBagaryan/portal"
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
	gate.Await(func(msg any) {
		time.Sleep(time.Second * 2) // simulation handler long work todo remove
		fmt.Printf("%s", msg)
	})
	return &Storage{
		gate: gate,
		list: make(map[int64]*User),
	}
}

func (s *Storage) Add(name string, age uint) (*User, error) {
	u, err := New(name, age)
	if err != nil {
		return nil, err
	}
	s.gate.Send(*u)
	s.lock.Lock()
	s.list[u.id] = u
	s.lock.Unlock()
	return u, nil
}

func (s *Storage) Get(userID int64) (*User, error) {
	got, ok := s.list[userID]
	if !ok {
		return nil, fmt.Errorf("user not found: %d", userID)
	}
	return got, nil
}
