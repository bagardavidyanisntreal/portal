package user

import (
	"fmt"

	"github.com/DavidBagaryan/portal"
)

// todo remove this useless construction
type uselessHandler struct{}

func (u uselessHandler) Support(_ portal.Message) bool {
	return true
}

func (u uselessHandler) Handle(msg portal.Message) {
	//time.Sleep(time.Second * 2) // simulation handler long work
	fmt.Printf("%s\n", msg.Data())
}
