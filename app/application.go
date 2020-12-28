package app

import (
	"log"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApp start the application
func StartApp() {
	gin.ForceConsoleColor()
	MapURL()
	if err := router.Run(":8080"); err != nil {
		log.Fatalln(err)
	}
}
