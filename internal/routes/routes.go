package routes

import (
	"io"
	"net/http"
)

func RegisterRoutes(
	mux *http.ServeMux,
) {
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello...")
	})
}

func NewServer() http.Handler {
	mux := http.NewServeMux()

	RegisterRoutes(
		mux,
	)

	var handler http.Handler = mux
	return handler
}
