package users

import (
	"fmt"

	"github.com/deepaksinghkushwah/bookstore/bookstore-user-api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

// Save to save user
func (user *User) Save() *errors.RestErr {
	current := userDB[user.ID]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError("Email already registered")
		}
		return errors.NewBadRequestError("User already exists")
	}
	userDB[user.ID] = user
	return nil
}

// Get to get user from db
func (user *User) Get() *errors.RestErr {
	result := userDB[user.ID]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("User id %d not found", user.ID))
	}
	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	return nil
}

// GetAll return all users
func GetAll() (map[int64]*User, *errors.RestErr) {
	if userDB != nil {
		return userDB, nil
	}
	return nil, errors.NewNotFoundError("No user found")

}
