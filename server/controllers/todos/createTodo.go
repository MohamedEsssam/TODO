package todos

import (
	"../../services"
	"../../utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	db := utils.Init()
	vars := mux.Vars(r)

	userId := vars["userId"]

	var todo Todo
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &todo)

	text := todo.Text

	result := services.Create(db, text, userId)
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
