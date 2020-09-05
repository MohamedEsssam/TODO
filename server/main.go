package main

import (
	"./routes/public"
	"./utils"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	utils.ConnectDB()
	router := mux.NewRouter()

	routes.HandleUserReq(router)

	fmt.Println("Starting server on the port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}
