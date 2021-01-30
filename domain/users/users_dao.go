package users

import (
	"github.com/davidhwang-ij/go-demo_users_api/datasource/mysql/users_db"
	"github.com/davidhwang-ij/go-demo_users_api/utils/errors"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, password) VALUES (?, ?, ?, ?);"
)

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Password)
	if saveErr != nil {
		return errors.NewInternalServerError("database error")
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError("database error")
	}
	user.Id = userId

	return nil
}
