package app

import (
	"log"

	loggers "github.com/deepaksinghkushwah/bookstore-utils-api/loggers"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApp start the application
func StartApp() {
	gin.ForceConsoleColor()
	MapURL()
	loggers.GetLogger().Info("about to start app")
	if err := router.Run(":8080"); err != nil {
		log.Fatalln(err)
	}
}
