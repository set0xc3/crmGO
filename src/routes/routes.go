package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/set0xc3/crmGO/src/db"
	"github.com/set0xc3/crmGO/src/view/home"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	home.Index().Render(r.Context(), w)
}

func HandleClients(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if r.PathValue("id") != "" {
			id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
			if err != nil {
				log.Fatal(err)
				return
			}
			ok, client := db.GetClientById(id)
			if ok {
				w.Header().Set("Content-Type", "application/json")

				if err := json.NewEncoder(w).Encode(client); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}
		} else {
			clients := db.GetClientList()

			w.Header().Set("Content-Type", "application/json")

			if err := json.NewEncoder(w).Encode(clients); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	} else if r.Method == "POST" {
		var client db.Client

		err := json.NewDecoder(r.Body).Decode(&client)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		db.CreateClient(client)
	} else if r.Method == "DELETE" {
		id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)

		if err != nil {
			log.Fatal(err)
			return
		}

		db.DeleteClient(id)
	} else if r.Method == "UPDATE" {
		var client db.Client

		err := json.NewDecoder(r.Body).Decode(&client)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
		ok := db.UpdateClientById(id, client)
		if ok == false {
			log.Fatal("Failed UpdateClientById")
			return
		}

		if err != nil {
			log.Fatal(err)
			return
		}
	}
}

func RegisterRoutes(
	mux *http.ServeMux,
) {
	mux.HandleFunc("/", HandleIndex)

	mux.HandleFunc("GET /api/v1/clients", HandleClients)
	mux.HandleFunc("POST /api/v1/clients", HandleClients)
	mux.HandleFunc("GET /api/v1/clients/{id}", HandleClients)
	mux.HandleFunc("DELETE /api/v1/clients/{id}", HandleClients)
	mux.HandleFunc("UPDATE /api/v1/clients/{id}", HandleClients)
}

func NewServer() http.Handler {
	mux := http.NewServeMux()

	RegisterRoutes(
		mux,
	)

	var handler http.Handler = mux
	return handler
}
