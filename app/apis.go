package app

import (
	"github.com/PS-07/bookstore_users-api/controllers/health"
	"github.com/PS-07/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/health", health.Health)

	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
}
