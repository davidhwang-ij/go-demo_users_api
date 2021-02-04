package app

import "github.com/davidhwang-ij/go-demo_users_api_02/controllers/users"

func mapUrls() {
	router.POST("/users", users.Create)
	router.GET("/users/:user_id", users.Get)
	router.PUT("/users/:user_id", users.Update)
}
