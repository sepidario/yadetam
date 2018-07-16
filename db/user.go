package db

import (
	"github.com/jinzhu/gorm"
	"github.com/sepidario/yadetam/yadetam"
)

// UserService implements yadetam.UserService interface
type UserService struct {
	DB *gorm.DB
}

// CreateUser creates a user
func (s *UserService) CreateUser(user *yadetam.User) error {
	panic("not implemented")
}

// UpdateUser updates a user information (currently only Name)
func (s *UserService) UpdateUser(user *yadetam.User) error {
	panic("not implemented")
}

// DeleteUser deletes a user
func (s *UserService) DeleteUser(id string) error {
	panic("not implemented")
}
