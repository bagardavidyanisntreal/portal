package user

// CreatedMessage notifies about new User created with user data
type CreatedMessage struct {
	user User
}

func (n CreatedMessage) Data() any {
	return n.user
}
