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
	router.GET("/users/:id", users.GetUser)

	router.GET("/users-search", users.SearchUsers)

	router.PUT("/users/:id", users.UpdateUser)
	router.PATCH("/users/:id", users.UpdateUser)
	router.DELETE("/users/:id", users.DeleteUser)
	router.GET("/populate-db", users.PopulateUserTable)
}
