package app

import "github.com/davidhwang-ij/go-users_api/controllers/users"

func mapUrls() {
	router.POST("/users", users.Create)
}
