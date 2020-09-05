package todos

import (
	"../../services"
	"../../utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	db := utils.Init()
	vars := mux.Vars(r)

	userId := vars["userId"]

	result := services.GetTodos(db, userId)
	if result == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("")
		return
	}

	todos := Todos{}
	json.Unmarshal(result, &todos)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}
