package users

import (
	"net/http"
	"strconv"

	"github.com/deepaksinghkushwah/bookstore-user-api/domain/users"
	"github.com/deepaksinghkushwah/bookstore-user-api/services"
	errors "github.com/deepaksinghkushwah/bookstore-utils-api/rest_errors"

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
	users, err := services.UserService.AllUsers()
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, users.Marshal(c.GetHeader("X-Public") == "true"))
}

// CreateUser create user
func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError(err.Error())
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.UserService.CreateUser(&user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result.Marshal(c.GetHeader("X-Public") == "true"))
}

// GetUser search for user with id
func GetUser(c *gin.Context) {
	userID, userErr := getUserID(c.Param("id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}

	user, err := services.UserService.GetUser(userID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, user.Marshal(c.GetHeader("X-Public") == "true"))

}

// PopulateUserTable db
func PopulateUserTable(c *gin.Context) {
	err := services.UserService.PopulateUserTable()
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
	current, err := services.UserService.UpdateUser(isPartial, &user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 1,
		"msg":    "User updated",
		"user":   current.Marshal(c.GetHeader("X-Public") == "true"),
	})

}

// DeleteUser to delete user
func DeleteUser(c *gin.Context) {
	userID, userErr := getUserID(c.Param("id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}

	deleteErr := services.UserService.DeleteUser(userID)

	if deleteErr != nil {
		c.JSON(deleteErr.Status, deleteErr)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 1,
		"msg":    "User deleted",
	})

}

// SearchUsers return search users
func SearchUsers(c *gin.Context) {
	status := c.Query("status")
	results, err := services.UserService.SearchUsers(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, results.Marshal(c.GetHeader("X-Public") == "true"))
}
