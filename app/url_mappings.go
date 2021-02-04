package app

<<<<<<< HEAD
import "github.com/davidhwang-ij/go-demo_users_api_02/controllers/users"
=======
import "github.com/davidhwang-ij/go-demo_users_api/controllers/users"
>>>>>>> 905c897ea9f3eb8a10b67bf5f3c6f1ec13ccbe5e

func mapUrls() {
	router.POST("/users", users.Create)
	router.GET("/users/:user_id", users.Get)
	router.PUT("/users/:user_id", users.Update)
}
