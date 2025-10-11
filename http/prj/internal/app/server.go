package app

import (
	"fmt"
	"log"
	"net/http"
	"prj/pkg/config"
	"prj/pkg/database"

	"prj/internal/db"
	"prj/internal/item"
)

func Run() {
	config.Set()
	configs := config.Get()
	database.Connect()

	itemRepo := item.NewRepository(pool)
	itemHandler := item.NewHandler(itemRepo)

	mux := http.NewServeMux()
	mux.HandleFunc("/items", itemHandler.List)

	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("🚀 Server running at http://localhost%s\n", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
