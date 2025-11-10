package routes

import (
	"net/http"
	itemController "prj/internal/modules/item/controller"
	userController "prj/internal/modules/user/controller"
	"prj/pkg/middleware"
)

func ItemRoutes(router *http.ServeMux) {
	itemHandler := itemController.NewController()
	router.HandleFunc("/items", itemHandler.List)

	userHandler := userController.NewController()
	router.HandleFunc("/auth/register", userHandler.Register)
	router.HandleFunc("/auth/login", userHandler.Login)

	router.Handle("/items", middleware.AuthMiddleware(
		"super_secret_key")(http.HandlerFunc(itemHandler.List)),
	)
}
