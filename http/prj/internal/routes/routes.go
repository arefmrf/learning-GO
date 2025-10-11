package routes

import (
	"net/http"
	"prj/internal/modules/item/controller"
)

func ItemRoutes(router *http.ServeMux) {
	itemHandler := controller.NewController()
	router.HandleFunc("/items", itemHandler.List)
}
