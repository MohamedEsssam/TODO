package main

import (
	// "fmt"
	// "log"
	// "net/http"
	// "./routes"
	"./utils"
)

func main() {

	db := utils.ConnectDB()
	utils.CreateModels(db)

	// fmt.Println("Starting server on the port 8000...")
	// log.Fatal(http.ListenAndServe(":8000", Router))
}
