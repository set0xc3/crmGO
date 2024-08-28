package main

import (
	"crmgo/app/middleware"
	"crmgo/app/routes"
	"crmgo/db"
	"log"
	"net/http"
)

func main() {
	db.Init()
	defer db.DeInit()

	server := http.NewServeMux()
	routes.AddRoutes(server)

	stack := middleware.CreateStack(
		middleware.Logging,
	)

	httpServer := http.Server{
		Addr:    ":8080",
		Handler: stack(server),
	}

	log.Println("Starting server on port :8080")

	httpServer.ListenAndServe()
}
