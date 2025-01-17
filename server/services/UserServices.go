package services

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

// User is..
type User struct {
	UserID string `json:"userId"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

// Login is..
func Login(db *sql.DB, email string, password string) []byte {
	const query = "SELECT userId, name, email FROM user WHERE email = ? AND password = ?"
	row, err := db.Query(query, email, password)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println(query)
		defer row.Close()
	}

	user := User{}
	hasResults := false
	for row.Next() {
		hasResults = true
		if err := row.Scan(&user.UserID, &user.Name, &user.Email); err != nil {
			panic(err.Error())
		}
	}

	if !hasResults {
		return nil
	}

	res, _ := json.Marshal(user)

	return res
}

// Register is..
func Register(db *sql.DB, name string, email string, password string) []byte {
	var user User
	json.Unmarshal([]byte(getUserByEmail(db, email)), &user)
	if (User{} != user) {
		return nil
	}

	const query = "INSERT INTO user (userId, name, email, password)VALUES(UUID(), ?, ?, ?)"
	insert, err := db.Prepare(query)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println(query)
	}
	insert.Exec(name, email, password)
	defer insert.Close()

	return getUserByEmail(db, email)
}

// getUserByEmail is..
func getUserByEmail(db *sql.DB, email string) []byte {
	const query = "SELECT userId, name, email FROM user WHERE email = ?"
	row, err := db.Query(query, email)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println(query)
		defer row.Close()
	}

	user := User{}
	hasResults := false
	for row.Next() {
		hasResults = true
		if err := row.Scan(&user.UserID, &user.Name, &user.Email); err != nil {
			panic(err.Error())
		}
	}

	if !hasResults {
		return nil
	}

	res, _ := json.Marshal(user)

	return res
}
