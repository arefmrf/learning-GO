package bootstrap

import (
	"fmt"
	"log"
	"net/http"
	"prj/pkg/config"
	"prj/pkg/database"
	"prj/pkg/routing"
)

type Server struct {
}

func Run() {
	config.Set()
	configs := config.Get()
	database.Connect()

	routing.Init()
	routing.RegisterRoutes()

	addr := fmt.Sprintf(":%s", configs.Server.Port)
	log.Printf("🚀 Server running at http://localhost%s\n", addr)

	if err := http.ListenAndServe(addr, routing.GetRouter()); err != nil {
		log.Fatal(err)
	}
}
