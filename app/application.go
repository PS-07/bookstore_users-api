package app

import "github.com/gin-gonic/gin"

var router = gin.Default()

// StartApp func
func StartApp() {
	mapUrls()
	router.Run(":8080")
}
