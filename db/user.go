package db

import (
	"github.com/jinzhu/gorm"
	"github.com/sepidario/yadetam/yadetam"
)

// UserService implements yadetam.UserService interface
type UserService struct {
	db *gorm.DB
}

// NewUserService creates a new user service
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db}
}

// CreateUser creates a user
func (s *UserService) CreateUser(user *yadetam.User) error {
	return s.db.Create(user).Error
}

// UpdateUser updates a user information (currently only Name)
func (s *UserService) UpdateUser(user *yadetam.User) error {
	return s.db.Save(user).Error
}

// DeleteUser deletes a user
func (s *UserService) DeleteUser(id string) error {
	return s.db.Delete(&yadetam.User{ID: id}).Error
}
