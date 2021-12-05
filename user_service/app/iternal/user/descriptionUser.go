package user

type User struct {
	id       int
	UUID     string
	Email    string
	Password string
}

type CreateUser struct {
	Email          string
	Password       string
	RepeatPassword string
}

func (user *User) GetName() string {
	return user.UUID
}
func (user *User) GetEmail() string {
	return user.Email
}
