package account

import "fmt"

type CreatedNotify struct {
	account Account
}

func (d CreatedNotify) Data() any {
	return fmt.Sprintf("created new account %s", d.account)
}
