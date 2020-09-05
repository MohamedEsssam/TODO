package routes

import (
	"./private"
	"./public"
	"github.com/gorilla/mux"
)

func Routes(router *mux.Router) {

	privateRoutes.HandleTodoReq(router)
	publicRoutes.HandleUserReq(router)

}
