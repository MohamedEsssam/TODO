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

func Create(db *sql.DB, text string, userId string) []byte {
	var query = "SELECT UUID()"

	uid, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println(query)
		defer uid.Close()
	}
	todo := Todo{}
	for uid.Next() {
		if err := uid.Scan(&todo.TodoID); err != nil {
			panic(err.Error())
		}
	}

	query = "INSERT INTO todoList (todoId, text, createdAt, userId) VALUES (?, ?, NOW(), ?)"
	insert, err := db.Query(query, todo.TodoID, text, userId)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println(query)
		defer insert.Close()
	}

	return GetTodo(db, todo.TodoID, userId)
}

func UpdateText(db *sql.DB, todoId string, userId string, text string) []byte {
	todo := GetTodo(db, todoId, userId)
	if todo == nil {
		return nil
	}

	const query = "UPDATE todoList SET text = ? WHERE todoId = ? AND userId = ?"
	update, err := db.Query(query, text, todoId, userId)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println(query)
		defer update.Close()
	}

	return GetTodo(db, todoId, userId)
}

func UpdateStatus(db *sql.DB, todoId string, userId string, status string) []byte {
	todo := GetTodo(db, todoId, userId)
	if todo == nil {
		return nil
	}

	const query = "UPDATE todoList SET status = ? WHERE todoId = ? AND userId = ?"
	update, err := db.Query(query, status, todoId, userId)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println(query)
		defer update.Close()
	}

	return GetTodo(db, todoId, userId)
}

func Delete(db *sql.DB, todoId string, userId string) []byte {
	todo := GetTodo(db, todoId, userId)
	if todo == nil {
		return nil
	}

	const query = "DELETE FROM todoList WHERE todoId = ? AND userId = ?"
	delete, err := db.Query(query, todoId, userId)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println(query)
		defer delete.Close()
	}

	return todo
}
