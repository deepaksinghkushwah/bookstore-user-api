package services

import (
	"github.com/deepaksinghkushwah/bookstore/bookstore-user-api/domain/users"
	"github.com/deepaksinghkushwah/bookstore/bookstore-user-api/utils/errors"
)

// CreateUser to create user in db
func CreateUser(user *users.User) (*users.User, *errors.RestErr) {
	err := user.Validate()
	if err != nil {
		return nil, err
	}
	if err = user.Save(); err != nil {
		return nil, err
	}
	return user, nil
}

// FindUser find user
func FindUser(userID int64) (*users.User, *errors.RestErr) {
	/*if userID <= 0 {
		return nil, errors.NewBadRequestError("Invalid user id")
	}*/
	result := users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return &result, nil

}

// AllUsers return all users
func AllUsers() (*map[int64]*users.User, *errors.RestErr) {
	results, err := users.GetAll()
	if err != nil {
		return nil, err
	}

	return &results, nil
}