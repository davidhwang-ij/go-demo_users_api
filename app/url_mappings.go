package app

import "github.com/davidhwang-ij/go-demo_users_api/controllers/users"

func mapUrls() {
	router.POST("/users", users.Create)
	router.GET("/users/:user_id", users.Get)
}
