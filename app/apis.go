package app

import (
	"github.com/PS-07/bookstore_users-api/controllers/health"
	"github.com/PS-07/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/health", health.Health)

	router.GET("/users/:user_id", users.Get)
	router.POST("/users", users.Create)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
}
