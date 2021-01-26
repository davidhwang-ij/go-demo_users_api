package users_db

import (
	"database/sql"
	"fmt"
	"log"
)

var (
	Client   *sql.DB
	username = "root"
	password = ""
	host     = "127.0.0.1:3306"
	schema   = "users_db_01"
)

func init() {
	// username:password@tcp(host)/user_schema
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)?charset=utf8", username, password, host, schema)

	var err error
	Client, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("database successfully configured")
}