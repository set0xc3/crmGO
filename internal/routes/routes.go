package routes

import (
	"net/http"

	"github.com/set0xc3/crmGO/internal/view/home"
)

func RegisterRoutes(
	mux *http.ServeMux,
) {
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		home.Index().Render(r.Context(), w)
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
