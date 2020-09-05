package main

import (
	"./routes/private"
	"./routes/public"
	"./utils"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	utils.ConnectDB()
	// utils.CreateModels()

	router := mux.NewRouter()

	privateRoutes.HandleTodoReq(router)
	publicRoutes.HandleUserReq(router)

	fmt.Println("Starting server on the port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}
