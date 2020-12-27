package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

// StartApp start the application
func StartApp() {
	MapURL()
	router.Run(":8080")
}
