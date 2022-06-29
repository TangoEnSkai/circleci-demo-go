package main

import (
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

var indexTmpl = template.Must(template.ParseFiles("templates/index.html"))

// HandleGet handles HTTP GET requests.
func HandleGet(w http.ResponseWriter, r *http.Request) {
	err := indexTmpl.Execute(w, nil)
	if err != nil {
		// status 500, internal server error
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// App holds a router for testing and production.
type App struct {
	Router *mux.Router
}

func (a *App) Init() {
	a.Router = mux.NewRouter()
	a.Router.HandleFunc("/", HandleGet).Methods("GET")
}
