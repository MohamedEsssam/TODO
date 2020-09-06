package privateRoutes

import (
	"../../controllers/todos"
	"github.com/gorilla/mux"
)

func HandleTodoReq(router *mux.Router) {
	router.HandleFunc("/{userId}", todos.GetTodos).Methods("GET", "OPTIONS")
	router.HandleFunc("/{userId}/{todoId}", todos.GetTodo).Methods("GET", "OPTIONS")
	router.HandleFunc("/{userId}", todos.Create).Methods("POST", "OPTIONS")
	router.HandleFunc("/{userId}", todos.Update).Methods("PUT", "OPTIONS")
	router.HandleFunc("/{userId}", todos.Delete).Methods("DELETE", "OPTIONS")
}
