package app

import (
	"github.com/PS-07/bookstore_users-api/logger"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

// StartApp func
func StartApp() {
	mapUrls()

	logger.Info("about to start the application...")
	router.Run(":8080")
}
