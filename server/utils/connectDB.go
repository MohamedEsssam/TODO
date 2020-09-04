package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "root:555FFaa@eerewqMohamed@111@/todo")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connect to DB successfully.")
	}

	return db
}
