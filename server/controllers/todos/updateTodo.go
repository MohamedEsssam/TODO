package todos

import (
	"../../services"
	"../../utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func Update(w http.ResponseWriter, r *http.Request) {
	db := utils.Init()
	vars := mux.Vars(r)

	userId := vars["userId"]

	var todo Todo
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &todo)

	todoId := todo.TodoID
	text := todo.Text
	status := todo.Status

	var result []byte
	if text != "" {

		result = services.UpdateText(db, todoId, userId, text)
	}
	if status != "" {
		result = services.UpdateStatus(db, todoId, userId, status)
	}

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
