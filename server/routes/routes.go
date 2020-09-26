package routes

import (
	"./private"
	"./public"
	"github.com/gorilla/mux"
)

func Routes(router *mux.Router) {
	publicRoutes.HandleUserReq(router)
	privateRoutes.HandleTodoReq(router)
}
