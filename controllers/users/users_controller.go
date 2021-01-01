package users

import (
	"bookstore/bookstore-user-api/domain/users"
	"bookstore/bookstore-user-api/services/userservice"
	"bookstore/bookstore-user-api/utils/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getUserID(userIDParam string) (int64, *errors.RestErr) {
	userID, userErr := strconv.ParseInt(userIDParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("Invalid user id")
	}
	return userID, nil
}

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
	userID, userErr := getUserID(c.Param("id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
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

// UpdateUser controllers
func UpdateUser(c *gin.Context) {
	userID, userErr := getUserID(c.Param("id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}

	var user users.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadGateway, errors.NewBadRequestError("Invalid user id"))
		return
	}
	user.ID = userID
	isPartial := c.Request.Method == http.MethodPatch
	current, err := userservice.UpdateUser(isPartial, &user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 1,
		"msg":    "User updated",
		"user":   current,
	})

}

// DeleteUser to delete user
func DeleteUser(c *gin.Context) {
	userID, userErr := getUserID(c.Param("id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}

	deleteErr := userservice.DeleteUser(userID)

	if deleteErr != nil {
		c.JSON(deleteErr.Status, deleteErr)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 1,
		"msg":    "User deleted",
	})

}
