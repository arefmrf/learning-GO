package main

import (
	"fmt"
	"log"
	"net/http"
	"site/middleware"
	"site/router"
	"site/routes"
)

func main() {
	r := router.NewRouter()

	r.Handle("/", routes.Home, middleware.Recover, middleware.Logging)

	r.Handle("/api/users", routes.Users, middleware.Recover, middleware.Logging, middleware.Auth)

	fmt.Println("Server running on http://localhost:8080")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
