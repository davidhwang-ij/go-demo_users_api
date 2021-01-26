package users

import (
	"net/http"

	"github.com/davidhwang-ij/go-demo_users_api/domain/users"
	"github.com/davidhwang-ij/go-demo_users_api/services"
	"github.com/davidhwang-ij/go-demo_users_api/utils/errors"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}
