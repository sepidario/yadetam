package yadetam

// User defines a user
type User struct {
	ID   string
	OS   string
	Name string
	Base
}

// UserService is the interface that intracts with a user
type UserService interface {
	CreateUser(user *User) error
	UpdateUser(user *User) error
	DeleteUser(id string) error
}
