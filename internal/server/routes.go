package server

import (
	// "encoding/json"
	// "log"
	"net/http"

	"github.com/a-h/templ"

	"github.com/set0xc3/crmGO/cmd/web"
)

func (s *Server) RegisterRoutes() http.Handler {
	fileServer := http.FileServer(http.FS(web.Files))
	mux := http.NewServeMux()
	mux.Handle("/", templ.Handler(web.Base()))

	mux.HandleFunc("/health", s.healthHandler)

	mux.Handle("/assets/", fileServer)
	mux.Handle("/web", templ.Handler(web.HelloForm()))
	mux.HandleFunc("/hello", web.HelloWebHandler)

	return mux
}

func (s *Server) IndexHandler(w http.ResponseWriter, r *http.Request) {
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
}
