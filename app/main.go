package main

import (
	"crmgo/app/middleware"
	"log"
	"net/http"
)

func NewServer() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", http.NotFoundHandler())

	mux.HandleFunc("GET /api", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("api"))
	})
	mux.HandleFunc("GET /about", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("about"))
	})
	mux.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("admin"))
	})

	var handler http.Handler = mux
	return handler
}

func main() {
	srv := NewServer()

	httpServer := http.Server{
		Addr:    ":8080",
		Handler: middleware.Logging(srv),
	}

	log.Println("Starting server on port :8080")

	httpServer.ListenAndServe()
}
