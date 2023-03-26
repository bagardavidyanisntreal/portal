package user

import (
	"encoding/json"
	"errors"
	"time"
)

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

func New(name string, age uint) (*User, error) {
	if age > 150 {
		return nil, errors.New("it's too much for any person: have more then 150 yo")
	}
	if len(name) > 250 || len(name) < 2 {
		return nil, errors.New("username cannot be less then 2 or more then 250 symbols length")
	}
	return &User{
		ID:   time.Now().Unix(),
		Name: name,
		Age:  age,
	}, nil
}

func (u User) String() string {
	b, _ := json.Marshal(u)
	return string(b)
}
