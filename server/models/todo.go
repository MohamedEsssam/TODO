package models

import (
	"database/sql"
	"fmt"
)

func TodoModel(db *sql.DB) {
	const query = "DROP TABLE IF EXISTS todoList; CREATE TABLE todoList (todoId varchar(36) NOT NULL PRIMARY KEY, text VARCHAR(255) NOT NULL, createdAt DATETIME NOT NULL, userId varchar(36) NOT NULL, CONSTRAINT fk_userId FOREIGN KEY (userId) REFERENCES user(userId) ON DELETE RESTRICT ON UPDATE CASCADE)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;"
	_, err := db.Exec(query)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println(query)
	}
}
