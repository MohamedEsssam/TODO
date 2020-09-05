package todos

import (
	"../../services"
	"../../utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetTodo(w http.ResponseWriter, r *http.Request) {
	db := utils.Init()
	vars := mux.Vars(r)

	userId := vars["userId"]
	todoId := vars["todoId"]

	result := services.GetTodo(db, todoId, userId)
	if result == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("")
		return
	}

	todo := Todo{}
	json.Unmarshal(result, &todo)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todo)
}
