package utils

import (
	"../models"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "root:555FFaa@eerewqMohamed@111@/todo?multiStatements=true")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connect to DB successfully.")
	}

	return db
}

func CreateModels(db *sql.DB) {
	models.UserModel(db)
	models.TodoModel(db)
}
