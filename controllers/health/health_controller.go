package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health func
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, "status: ok")
}
