package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Ping only
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "Ping received")
}
