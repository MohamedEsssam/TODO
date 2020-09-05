package utils

import (
	"../models"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var dbInstance *sql.DB

func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "root:555FFaa@eerewqMohamed@111@/todo?multiStatements=true")

	dbInstance = db
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connect to DB successfully.")
	}

	return db
}

func Init() *sql.DB {
	return dbInstance
}

func CreateModels() {
	models.UserModel(dbInstance)
	models.TodoModel(dbInstance)
}
