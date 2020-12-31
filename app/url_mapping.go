package app

import (
	"bookstore/bookstore-user-api/controllers/ping"
	"bookstore/bookstore-user-api/controllers/users"
)

// MapURL map urls
func MapURL() {
	router.GET("/", ping.Ping)
	router.GET("/users", users.GetUsers)
	router.POST("/users", users.CreateUser)
	router.GET("/users/:id", users.FindUser)
	router.PUT("/users", users.UpdateUser)
	router.GET("/populate-db", users.PopulateUserTable)
}
