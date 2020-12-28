package users

import (
	"net/http"
	"strconv"

	"github.com/deepaksinghkushwah/bookstore/bookstore-user-api/domain/users"
	"github.com/deepaksinghkushwah/bookstore/bookstore-user-api/services"
	"github.com/deepaksinghkushwah/bookstore/bookstore-user-api/utils/errors"
	"github.com/gin-gonic/gin"
)

// GetUsers return list of users
func GetUsers(c *gin.Context) {
	users, err := services.AllUsers()
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
	result, saveErr := services.CreateUser(&user)
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

	user, err := services.FindUser(userID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, user)

}

// PopulateUserTable
func PopulateUserTable(c *gin.Context) {
	err := services.PopulateUserTable()
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "Db populated",
	})
}
