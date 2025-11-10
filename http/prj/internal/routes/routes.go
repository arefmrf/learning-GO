package routes

import (
	"net/http"
	itemController "prj/internal/modules/item/controller"
	userController "prj/internal/modules/user/controller"
	"prj/pkg/config"
	"prj/pkg/middleware"
)

func ItemRoutes(router *http.ServeMux) {
	itemHandler := itemController.NewController()
	userHandler := userController.NewController()
	configs := config.Get()

	// Public routes
	router.HandleFunc("/auth/register", userHandler.Register)
	router.HandleFunc("/auth/login", userHandler.Login)

	router.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			itemHandler.List(w, r)
		case http.MethodPost:
			// Apply auth middleware manually for POST only
			middleware.AuthMiddleware(configs.Server.JwtSecret)(http.HandlerFunc(itemHandler.Create)).ServeHTTP(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
