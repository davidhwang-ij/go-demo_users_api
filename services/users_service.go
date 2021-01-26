package services

import (
	"github.com/davidhwang-ij/go-demo_users_api/domain/users"
	"github.com/davidhwang-ij/go-demo_users_api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	// Validate
	if err := user.Validate(); err != nil {
		return nil, err
	}
	// Save to database
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}
