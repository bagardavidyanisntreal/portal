package account

import (
	"fmt"

	"github.com/bagardavidyanisntreal/portal/v2/example/user"
)

type createAccountForNewUser struct {
	storage *Storage
}

const startBalance = 100

func (h createAccountForNewUser) Handle(msg any) {
	u, ok := msg.(*user.User)
	if !ok {
		return
	}
	_, err := h.storage.Add(u, startBalance)
	if err != nil {
		fmt.Printf("an err occured on creating for user %v\n", u)
	}
}
