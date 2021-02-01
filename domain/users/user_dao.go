package users

import (
	"fmt"

	"github.com/PS-07/bookstore_users-api/utils/errors"
)

var userDB = make(map[int64]*User)

// Get func
func (user *User) Get() *errors.RestErr {
	result := userDB[user.ID]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.ID))
	}
	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

// Save func
func (user *User) Save() *errors.RestErr {
	current := userDB[user.ID]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.ID))
	}
	userDB[user.ID] = user
	return nil
}
