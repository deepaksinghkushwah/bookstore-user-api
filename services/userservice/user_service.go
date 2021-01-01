package userservice

import (
	"bookstore/bookstore-user-api/domain/users"
	"bookstore/bookstore-user-api/utils/errors"
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
func AllUsers() (*[]users.User, *errors.RestErr) {
	results, err := users.GetAll()
	if err != nil {
		return nil, err
	}

	return results, nil
}

// PopulateUserTable to topup user table
func PopulateUserTable() *errors.RestErr {
	err := users.PopulateUserTable()
	if err != nil {
		return err
	}
	return nil
}

// UpdateUser service
func UpdateUser(isPartial bool, user *users.User) (*users.User, *errors.RestErr) {
	current, err := FindUser(user.ID)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName == "" {
			user.FirstName = current.FirstName
		}

		if user.LastName == "" {
			user.LastName = current.LastName
		}

		if user.Email == "" {
			user.Email = current.Email
		}

		if user.DateCreated == "" {
			user.DateCreated = current.DateCreated
		}

	}
	err = user.UpdateUser()

	if err != nil {
		return nil, err
	}
	return user, nil
}

// DeleteUser to delete user
func DeleteUser(userID int64) *errors.RestErr {
	user := &users.User{ID: userID}
	return user.DeleteUser()

}
