package users

import (
	"strings"

	"bookstore/bookstore-user-api/utils/errors"
)

// User struct
type User struct {
	ID          int64  `json:"id"`
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

// PublicUser for external api use
type PublicUser struct {
	ID          int64  `json:"id"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

// PrivateUser for internal use
type PrivateUser struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

// Validate user
func (p *User) Validate() *errors.RestErr {
	p.Email = strings.TrimSpace(strings.ToLower(p.Email))
	if p.Email == "" {
		return errors.NewBadRequestError("Email is not valid")
	}

	p.FirstName = strings.TrimSpace(p.FirstName)
	if p.FirstName == "" {
		return errors.NewBadRequestError("First name is not valid")
	}

	p.LastName = strings.TrimSpace(p.LastName)
	if p.LastName == "" {
		return errors.NewBadRequestError("Last name is not valid")
	}
	return nil
}
