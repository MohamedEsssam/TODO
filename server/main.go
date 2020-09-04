package main

import (
	"fmt"
	// "log"
	// "net/http"
	// "./routes"
	"./services"
	"./utils"
)

func main() {
	db := utils.ConnectDB()
	// utils.CreateModels(db)

	// result := services.Login(db, "email1@yahoo.com", "1111")
	// if result == "" {
	// 	fmt.Println("sasa")
	// } else {
	// 	fmt.Println(result)
	// }

	result := services.Register(db, "Mohamed", "email1@gmail.com", "1111")
	if result == "" {
		fmt.Println("sasa")
	} else {
		fmt.Println(result)
	}

	// fmt.Println("Starting server on the port 8000...")
	// log.Fatal(http.ListenAndServe(":8000", Router))
}
