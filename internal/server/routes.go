package server

import (
	// "encoding/json"
	// "log"
	"net/http"

	"github.com/a-h/templ"
	// "github.com/aarol/reload"

	"github.com/set0xc3/crmGO/cmd/web"
)

func (s *Server) RegisterRoutes() http.Handler {
	fileServer := http.FileServer(http.FS(web.Files))
	mux := http.NewServeMux()
	mux.Handle("/", templ.Handler(web.Base()))
	mux.Handle("/assets/", fileServer)

	return mux
}

func (s *Server) IndexHandler(w http.ResponseWriter, r *http.Request) {
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
}
