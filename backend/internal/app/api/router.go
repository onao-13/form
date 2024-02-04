package api

import (
	"backend/internal/app/controller"
	"github.com/gorilla/mux"
	"net/http"
)

// NewHandler REST API сервиса
func NewHandler(people controller.PeopleEmployment, web controller.Web) *mux.Router {
	r := mux.NewRouter()

	// API
	r.HandleFunc("/api/send", people.Send)

	// WEB
	r.HandleFunc("/", web.FormPage)
	fileServer := http.FileServer(http.Dir("frontend/static/"))
	r.PathPrefix("/static").Handler(http.StripPrefix("/static", fileServer))

	return r
}
