package users

import (
	"strconv"

	"github.com/deepaksinghkushwah/bookstore/bookstore-user-api/datasource/mysql/bookstore"
	"github.com/deepaksinghkushwah/bookstore/bookstore-user-api/utils/dates"
	"github.com/deepaksinghkushwah/bookstore/bookstore-user-api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

const (
	queryInsertUser = "INSERT INTO `users` (first_name, last_name, email, date_created) VALUES(?, ?, ?, ?)"
)

// Save to save user
func (user *User) Save() *errors.RestErr {
	if err := bookstore.BookStoreDBLink.Ping(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	stmt, err := bookstore.BookStoreDBLink.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	user.DateCreated = dates.GetNowString()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		return errors.NewInternalServerError("Error when trying to save user")
	}
	userID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError("Error when trying to retrive last insert id")
	}
	user.ID = userID
	current := userDB[user.ID]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError("Email already registered")
		}
		return errors.NewBadRequestError("User already exists")
	}
	return nil
}

// Get to get user from db
func (user *User) Get() *errors.RestErr {
	if err := bookstore.BookStoreDBLink.Ping(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	stmt, err := bookstore.BookStoreDBLink.Prepare("SELECT id, first_name, last_name, email, date_created FROM `users` WHERE id = ?")
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	err = stmt.QueryRow(user.ID).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

// GetAll return all users
func GetAll() (*[]User, *errors.RestErr) {
	var users []User
	stmt, err := bookstore.BookStoreDBLink.Prepare("SELECT id, first_name, last_name, email, date_created FROM `users`")
	if err != nil {
		return nil, errors.NewNotFoundError("No users found")
	}
	defer stmt.Close()
	results, err := stmt.Query()
	if err != nil {
		return nil, errors.NewNotFoundError("No users found")
	}
	defer results.Close()

	for results.Next() {
		var user User
		err = results.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated)
		if err != nil {
			return nil, errors.NewInternalServerError(err.Error())
		}
		users = append(users, user)
	}

	if len(users) > 0 {
		return &users, nil
	}

	return nil, errors.NewNotFoundError("No user found")

}

// PopulateUserTable populate user table
func PopulateUserTable() *errors.RestErr {
	stmt, err := bookstore.BookStoreDBLink.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	for i := 101; i <= 1000; i++ {
		var firstName = "Test"
		var lastName = strconv.Itoa(i)
		var email = firstName + lastName + "@localhost.com"
		var dateCreated = dates.GetNowString()
		_, err := stmt.Exec(firstName, lastName, email, dateCreated)
		if err != nil {
			return errors.NewInternalServerError("Error when trying to save user")
		}
	}
	return nil

}
