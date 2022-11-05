package account

import (
	"fmt"

	"github.com/bagardavidyanisntreal/portal"
	"github.com/bagardavidyanisntreal/portal/example/user"
)

type createAccountForNewUser struct {
	storage *Storage
}

func (h createAccountForNewUser) Support(msg portal.Message) bool {
	_, ok := msg.(*user.CreatedMessage)
	return ok
}

const startBalance = 100

func (h createAccountForNewUser) Handle(msg portal.Message) {
	data := msg.Data()
	u, ok := data.(user.User)
	if !ok {
		return
	}
	_, err := h.storage.Add(u, startBalance)
	if err != nil {
		fmt.Printf("an err occured on creating for user %s", u)
	}
}
