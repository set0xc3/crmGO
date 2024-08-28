package routes

import (
	"crmgo/db"
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"
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
		i, err := strconv.ParseInt(r.PathValue("id"), 10, 64)

		if err != nil {
			log.Fatal(err)
			return
		}

		db.DeleteClient(i)
	}
}

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /clients", HandleClients)
	mux.HandleFunc("GET /clients/{id}", HandleClients)
	mux.HandleFunc("DELETE /clients/{id}", HandleClients)
	mux.HandleFunc("POST /clients", HandleClients)
}
