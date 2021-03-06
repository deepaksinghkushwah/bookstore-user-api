package users

import (
	"strings"

	errors "github.com/deepaksinghkushwah/bookstore-utils-api/rest_errors"
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
