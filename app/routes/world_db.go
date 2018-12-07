package routes

import (
	"net/http"

	"github.com/Polipapik/REST_API/app/handlers"
	"github.com/Polipapik/REST_API/app/models"
	"github.com/gorilla/mux"
)

//Env comment
type Env struct {
	Router  *mux.Router
	Country models.CountryAPI
}

//InitializeRoutes comment
func (env *Env) InitializeRoutes() {

	env.Router.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			handlers.GetCountry(env.Country, w, r)
		}).
		Queries( // just env try
			"username", "{username}",
			"email", "{email}",
		).
		Methods("GET")
	env.Router.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			handlers.GetCountries(env.Country, w, r)
		}).
		Methods("GET")
	env.Router.HandleFunc("/countries",
		func(w http.ResponseWriter, r *http.Request) {
			handlers.GetCountries(env.Country, w, r)
		}).
		Methods("GET")
	env.Router.HandleFunc("/country",
		func(w http.ResponseWriter, r *http.Request) {
			handlers.CreateCountry(env.Country, w, r)
		}).
		Methods("POST")
	env.Router.HandleFunc("/country/{id:[0-9]+}",
		func(w http.ResponseWriter, r *http.Request) {
			handlers.GetCountry(env.Country, w, r)
		}).
		Methods("GET")
	env.Router.HandleFunc("/country/{id:[0-9]+}",
		func(w http.ResponseWriter, r *http.Request) {
			handlers.UpdateCountry(env.Country, w, r)
		}).
		Methods("PUT")
	env.Router.HandleFunc("/country/{id:[0-9]+}",
		func(w http.ResponseWriter, r *http.Request) {
			handlers.DeleteCountry(env.Country, w, r)
		}).
		Methods("DELETE")
}
