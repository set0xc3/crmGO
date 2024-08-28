package routes

import (
	"crmgo/db"
	"fmt"
	"html"
	"net/http"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HOME, %q", html.EscapeString(r.URL.Path))
}

func HandleAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API, %q", html.EscapeString(r.URL.Path))
}

func HandleAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ABOUT, %q", html.EscapeString(r.URL.Path))
}

func HandleClients(w http.ResponseWriter, r *http.Request) {
	fmt.Println(db.ReadClientList())
}

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /", HandleIndex)
	mux.HandleFunc("GET /api", HandleAPI)
	mux.HandleFunc("GET /about", HandleAbout)
	mux.HandleFunc("GET /clients", HandleClients)
}
