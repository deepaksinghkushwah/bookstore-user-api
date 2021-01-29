package users

import (
	"strconv"

	"errors"

	"github.com/deepaksinghkushwah/bookstore-user-api/datasource/mysql/bookstore"
	"github.com/deepaksinghkushwah/bookstore-utils-api/cryptoutils"
	"github.com/deepaksinghkushwah/bookstore-utils-api/dates"
	"github.com/deepaksinghkushwah/bookstore-utils-api/loggers"
	RestErrors "github.com/deepaksinghkushwah/bookstore-utils-api/rest_errors"
)

var (
	userDB = make(map[int64]*User)
)

const (
	queryInsertUser       = "INSERT INTO `users` (first_name, last_name, email, date_created, password, status) VALUES(?, ?, ?, ?, ?, ?)"
	queryUpdateUser       = "UPDATE `users` SET first_name = ?, last_name = ?, email = ?, date_created = ? WHERE id = ?"
	queryDeleteUser       = "DELETE FROM `users` WHERE id = ?"
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, date_created, password, `status`  FROM `users` WHERE `status` = ?"
	queryFindUserByID     = "SELECT id, first_name, last_name, email, date_created, password, `status`  FROM `users` WHERE `id` = ?"
	queryFindAllUsers     = "SELECT id, first_name, last_name, email, date_created, password, `status`  FROM `users`"
)

// Save to save user
func (user *User) Save() *RestErrors.RestErr {
	if err := bookstore.BookStoreDBLink.Ping(); err != nil {
		return RestErrors.NewInternalServerError(err.Error(), errors.New("database error"))
	}
	stmt, err := bookstore.BookStoreDBLink.Prepare(queryInsertUser)
	if err != nil {
		return RestErrors.NewInternalServerError(err.Error(), errors.New("database error"))
	}
	defer stmt.Close()
	user.DateCreated = dates.GetNowDBString()
	user.Password = cryptoutils.GetMD5(user.Password)

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Password, user.Status)
	if err != nil {
		return RestErrors.NewInternalServerError(err.Error(), errors.New("database error"))
	}
	userID, err := insertResult.LastInsertId()
	if err != nil {
		return RestErrors.NewInternalServerError("Error when trying to retrive last insert id", errors.New("database error"))
	}
	user.ID = userID
	current := userDB[user.ID]
	if current != nil {
		if current.Email == user.Email {
			return RestErrors.NewBadRequestError("Email already registered")
		}
		return RestErrors.NewBadRequestError("User already exists")
	}
	return nil
}

// Get to get user from db
func (user *User) Get() *RestErrors.RestErr {
	if err := bookstore.BookStoreDBLink.Ping(); err != nil {
		return RestErrors.NewInternalServerError(err.Error(), errors.New("database error"))
	}
	stmt, err := bookstore.BookStoreDBLink.Prepare(queryFindUserByID)
	if err != nil {
		loggers.Error("error when trying to prepare user statment", err)
		return RestErrors.NewInternalServerError("Database error", errors.New("database error"))
	}

	defer stmt.Close()

	err = stmt.QueryRow(user.ID).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Password, &user.Status)
	if err != nil {
		loggers.Error("error when trying to execute user statment", err)
		return RestErrors.NewInternalServerError("Database error", errors.New("database error"))
	}
	return nil
}

// GetAll return all users
func GetAll() (Users, *RestErrors.RestErr) {
	var users []User
	stmt, err := bookstore.BookStoreDBLink.Prepare(queryFindAllUsers)
	if err != nil {
		return nil, RestErrors.NewNotFoundError("No users found")
	}
	defer stmt.Close()
	results, err := stmt.Query()
	if err != nil {
		return nil, RestErrors.NewNotFoundError("No users found")
	}
	defer results.Close()

	for results.Next() {
		var user User
		err = results.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Password, &user.Status)
		if err != nil {
			loggers.Error("error when trying to prepare user statment", err)
			return nil, RestErrors.NewInternalServerError("Database error", errors.New("database error"))
		}
		users = append(users, user)
	}

	if len(users) > 0 {
		return users, nil
	}

	return nil, RestErrors.NewNotFoundError("No user found")

}

// PopulateUserTable populate user table
func PopulateUserTable() *RestErrors.RestErr {
	stmt, err := bookstore.BookStoreDBLink.Prepare(queryInsertUser)
	if err != nil {
		loggers.Error("error when trying to prepare table statment", err)
		return RestErrors.NewInternalServerError("Database error", errors.New("database error"))
	}
	defer stmt.Close()

	for i := 1; i <= 1000; i++ {
		var firstName = "Test"
		var lastName = strconv.Itoa(i)
		var email = firstName + lastName + "@localhost.com"
		var dateCreated = dates.GetNowDBString()
		var password = cryptoutils.GetMD5("123456")
		var status = 1
		_, err := stmt.Exec(firstName, lastName, email, dateCreated, password, status)
		if err != nil {
			return RestErrors.NewInternalServerError("Error when trying to save user", errors.New("database error"))
		}
	}
	return nil

}

// UpdateUser to update user
func (user *User) UpdateUser() *RestErrors.RestErr {
	err := user.Validate()
	if err != nil {
		return err
	}

	stmt, prepErr := bookstore.BookStoreDBLink.Prepare(queryUpdateUser)
	if prepErr != nil {
		loggers.Error("error when trying to prepare user statment", prepErr)
		return RestErrors.NewInternalServerError("Database error", errors.New("database error"))
	}

	defer stmt.Close()

	_, updateErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.ID)
	if updateErr != nil {
		loggers.Error("error when trying to prepare user statment", updateErr)
		return RestErrors.NewInternalServerError("Update error", errors.New("database error"))
	}

	return nil
}

// DeleteUser to delete user
func (user *User) DeleteUser() *RestErrors.RestErr {
	stmt, err := bookstore.BookStoreDBLink.Prepare(queryDeleteUser)
	if err != nil {
		loggers.Error("error when trying to prepare user statment", err)
		return RestErrors.NewInternalServerError("Delete error", errors.New("database error"))
	}
	defer stmt.Close()
	_, deleteErr := stmt.Exec(user.ID)
	if deleteErr != nil {
		loggers.Error("error when trying to prepare user statment", deleteErr)
		return RestErrors.NewInternalServerError("Delete error", errors.New("database error"))
	}
	return nil

}

// SearchUsers to find user with params
func (user *User) SearchUsers(status string) (Users, *RestErrors.RestErr) {
	stmt, err := bookstore.BookStoreDBLink.Prepare(queryFindUserByStatus)
	if err != nil {
		loggers.Error("error when trying to prepare user statment", err)
		return nil, RestErrors.NewInternalServerError("Datebase error", errors.New("database error"))

	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		loggers.Error("error when trying to prepare user statment", err)
		return nil, RestErrors.NewInternalServerError("Update error", errors.New("database error"))
	}

	defer rows.Close()
	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Password, &user.Status); err != nil {
			loggers.Error("error when trying to scan user", err)
			return nil, RestErrors.NewInternalServerError("Update error", errors.New("database error"))

		}
		results = append(results, user)
	}
	return results, nil
}
