package users

import (
	"../../services"
	"../../utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	db := utils.Init()

	var user User
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &user)

	email := user.Email
	password := user.Password

	result := services.Login(db, email, password)
	if result == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("")
		return
	}

	var res User
	json.Unmarshal(result, &res)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
