package app

import (
	"fmt"
	"html"
	"net/http"
)

func handleAPI(r *http.Request) {
}

func addRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	// mux.HandleFunc("/api/", handleAPI())
	// mux.HandleFunc("/about", handleAbout())
	// mux.HandleFunc("/", handleIndex())
	// mux.HandleFunc("/admin", adminOnly(handleAdminIndex()))
}
