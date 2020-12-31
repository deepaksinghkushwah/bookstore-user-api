package users

import (
	"bookstore/bookstore-user-api/domain/users"
	"bookstore/bookstore-user-api/services/userservice"
	"bookstore/bookstore-user-api/utils/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUsers return list of users
func GetUsers(c *gin.Context) {
	users, err := userservice.AllUsers()
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, &users)
}

// CreateUser create user
func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError(err.Error())
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := userservice.CreateUser(&user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

// FindUser search for user with id
func FindUser(c *gin.Context) {
	userID, userErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("Invalid user id")
		c.JSON(err.Status, err)
		return
	}

	user, err := userservice.FindUser(userID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, user)

}

// PopulateUserTable db
func PopulateUserTable(c *gin.Context) {
	err := userservice.PopulateUserTable()
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "Db populated",
	})
}
