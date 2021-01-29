package services

import (
	"github.com/deepaksinghkushwah/bookstore-user-api/domain/users"
	errors "github.com/deepaksinghkushwah/bookstore-utils-api/rest_errors"
)

var (
	// UserService var
	UserService userServiceInterface = &userService{}
)

type userService struct{}

type userServiceInterface interface {
	CreateUser(*users.User) (*users.User, *errors.RestErr)
	GetUser(int64) (*users.User, *errors.RestErr)
	AllUsers() (users.Users, *errors.RestErr)
	PopulateUserTable() *errors.RestErr
	UpdateUser(bool, *users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	SearchUsers(string) (users.Users, *errors.RestErr)
}

// CreateUser to create user in db
func (u *userService) CreateUser(user *users.User) (*users.User, *errors.RestErr) {
	err := user.Validate()
	if err != nil {
		return nil, err
	}
	if err = user.Save(); err != nil {
		return nil, err
	}
	return user, nil
}

// GetUser find user with user id
func (u *userService) GetUser(userID int64) (*users.User, *errors.RestErr) {
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
func (u *userService) AllUsers() (users.Users, *errors.RestErr) {
	results, err := users.GetAll()
	if err != nil {
		return nil, err
	}

	return results, nil
}

// PopulateUserTable to topup user table
func (u *userService) PopulateUserTable() *errors.RestErr {
	err := users.PopulateUserTable()
	if err != nil {
		return err
	}
	return nil
}

// UpdateUser service
func (u *userService) UpdateUser(isPartial bool, user *users.User) (*users.User, *errors.RestErr) {
	current, err := UserService.GetUser(user.ID)
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
func (u *userService) DeleteUser(userID int64) *errors.RestErr {
	user := &users.User{ID: userID}
	return user.DeleteUser()

}

// SearchUsers user with params
func (u *userService) SearchUsers(status string) (users.Users, *errors.RestErr) {
	var user *users.User
	results, err := user.SearchUsers(status)
	if err != nil {
		return nil, err
	}
	return results, nil
}
