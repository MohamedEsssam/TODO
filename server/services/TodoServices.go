package services

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type Todo struct {
	TodoID string `json:"todoId"`
	Text   string `json:"text"`
	Status string `json:"status"`
	Date   string `json:"date"`
	UserID string `json:"userId"`
}

type Todos struct {
	Data []Todo `json:"data"`
}

func GetTodos(db *sql.DB, userId string) []byte {
	const query = "SELECT * FROM todoList WHERE userId = ?"
	rows, err := db.Query(query, userId)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println(query)
		defer rows.Close()
	}

	todo := Todo{}
	todos := Todos{}
	hasResults := false
	for rows.Next() {
		hasResults = true
		if err := rows.Scan(&todo.TodoID, &todo.Text, &todo.Status,
			&todo.Date, &todo.UserID); err != nil {
			panic(err.Error())
		}
		todos.Data = append(todos.Data, todo)
	}

	if !hasResults {
		return nil
	}

	res, _ := json.Marshal(todos)

	return res
}

func GetTodo(db *sql.DB, todoId string, userId string) []byte {
	const query = "SELECT * FROM todoList WHERE todoId = ? AND userId = ?"
	row, err := db.Query(query, todoId, userId)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println(query)
		defer row.Close()
	}

	todo := Todo{}
	hasResults := false
	for row.Next() {
		hasResults = true
		if err := row.Scan(&todo.TodoID, &todo.Text, &todo.Status,
			&todo.Date, &todo.UserID); err != nil {
			panic(err.Error())
		}
	}

	if !hasResults {
		return nil
	}

	res, _ := json.Marshal(todo)

	return res

}

func Create(db *sql.DB, text string, userId string) {

}

func Update(db *sql.DB, todoId string, userId string) {

}

func Delete(db *sql.DB, todoId string, userId string) {

}
