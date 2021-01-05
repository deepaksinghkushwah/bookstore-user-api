package users

import (
	"encoding/json"
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

// Marshal return user based on request type
func (p *User) Marshal(isPublic bool) interface{} {
	if isPublic == true {
		return PublicUser{
			ID:          p.ID,
			DateCreated: p.DateCreated,
			Status:      p.Status,
		}
	}

	userJSON, _ := json.Marshal(p)
	var privateUser PrivateUser
	json.Unmarshal(userJSON, &privateUser)
	return privateUser
}

// Users is [] of users
type Users []User

// Marshal to multiple users
func (p Users) Marshal(isPublic bool) []interface{} {
	result := make([]interface{}, len(p))
	for index, user := range p {
		result[index] = user.Marshal(isPublic)
	}
	return result
}
