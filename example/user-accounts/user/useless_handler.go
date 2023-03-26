package user

import (
	"fmt"
	"time"
)

// todo remove this useless construction
type uselessHandler struct{}

func (u uselessHandler) Handle(msg any) {
	time.Sleep(time.Second * 2) // simulation handler long work
	fmt.Printf("%T: %v\n", msg, msg)
}
