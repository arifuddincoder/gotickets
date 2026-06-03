package user

type Repository interface {
	CreateUser(user *User) error
}
