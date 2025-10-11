package routing

import (
	"net/http"
	"prj/internal/routes"
)

var router *http.ServeMux

func Init() {
	router = http.NewServeMux()
}

func GetRouter() *http.ServeMux {
	return router
}

func RegisterRoutes() {
	routes.ItemRoutes(GetRouter())
}
