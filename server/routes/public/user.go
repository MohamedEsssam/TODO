package routes

import (
	"../../controllers/users"
	"github.com/gorilla/mux"
)

func HandleUserReq(router *mux.Router) {
	router.HandleFunc("/login", users.Login).Methods("POST", "OPTIONS")
	router.HandleFunc("/register", users.Register).Methods("POST", "OPTIONS")

}
