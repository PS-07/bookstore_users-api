package services

import (
	"github.com/PS-07/bookstore_users-api/domain/users"
	"github.com/PS-07/bookstore_users-api/utils/errors"
)

// CreateUser func
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUser func
func GetUser(userID int64) (*users.User, *errors.RestErr) {
	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}
