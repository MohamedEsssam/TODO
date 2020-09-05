package main

import (
	"./routes"
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
	routes.Routes(router)

	fmt.Println("Starting server on the port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}
