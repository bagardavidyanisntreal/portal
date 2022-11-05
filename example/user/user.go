package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	id   int64
	name string
	age  uint
}

func New(name string, age uint) (*User, error) {
	if age > 150 {
		return nil, errors.New("it's too much for any person: have more then 150 yo")
	}
	if len(name) > 250 || len(name) < 2 {
		return nil, errors.New("username cannot be less then 2 or more then 250 symbols length")
	}
	return &User{
		id:   time.Now().Unix(),
		name: name,
		age:  age,
	}, nil
}

func (u User) GetID() int64 {
	return u.id
}

func (u User) String() string {
	return fmt.Sprintf(
		`{"id": %d, name": "%s", "age": %d}`,
		u.id,
		u.name,
		u.age,
	)
}

func (u User) CreatedMessage() *CreatedMessage {
	return &CreatedMessage{
		user: u,
	}
}
