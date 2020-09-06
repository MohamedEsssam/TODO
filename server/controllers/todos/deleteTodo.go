package todos

import (
	"../../services"
	"../../utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	db := utils.Init()
	vars := mux.Vars(r)

	userId := vars["userId"]

	var todo Todo
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &todo)

	todoId := todo.TodoID

	result := services.Delete(db, todoId, userId)
	if result == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("")
		return
	}

	var res Todo
	json.Unmarshal(result, &res)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
