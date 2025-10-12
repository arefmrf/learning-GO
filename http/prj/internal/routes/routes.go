package routes

import (
	"net/http"
	itemController "prj/internal/modules/item/controller"
	userController "prj/internal/modules/user/controller"
)

func ItemRoutes(router *http.ServeMux) {
	itemHandler := itemController.NewController()
	router.HandleFunc("/items", itemHandler.List)

	userHandler := userController.NewController()
	router.HandleFunc("/auth/register", userHandler.Register)
	router.HandleFunc("/auth/login", userHandler.Login)
}
