package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Polipapik/REST_API/app/models"

	"github.com/Polipapik/REST_API/app/handlers"
	"github.com/jinzhu/gorm"

	"github.com/gorilla/mux"
)

//App comment
type App struct {
	Router *mux.Router
	Counry models.CountryAPI
}

//Initialize comment
func (a *App) Initialize(host, port, user, password, dbname, sslmode string) {
	log.Println("Connection...")
	connectionString :=
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)

	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connection successful")

	a.Counry = &models.GormDB{DB: db}
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

//Run comment kek
func (a *App) Run(port string) {
	log.Fatal(http.ListenAndServe(port, a.Router))
}

func (a *App) initializeRoutes() {

	a.Router.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			handlers.GetCountry(a.Counry, w, r)
		}).
		Queries( // just a try
			"username", "{username}",
			"email", "{email}",
		).
		Methods("GET")
	a.Router.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			handlers.GetCountries(a.Counry, w, r)
		}).
		Methods("GET")
	a.Router.HandleFunc("/countries",
		func(w http.ResponseWriter, r *http.Request) {
			handlers.GetCountries(a.Counry, w, r)
		}).
		Methods("GET")
	a.Router.HandleFunc("/country",
		func(w http.ResponseWriter, r *http.Request) {
			handlers.CreateCountry(a.Counry, w, r)
		}).
		Methods("POST")
	a.Router.HandleFunc("/country/{id:[0-9]+}",
		func(w http.ResponseWriter, r *http.Request) {
			handlers.GetCountry(a.Counry, w, r)
		}).
		Methods("GET")
	a.Router.HandleFunc("/country/{id:[0-9]+}",
		func(w http.ResponseWriter, r *http.Request) {
			handlers.UpdateCountry(a.Counry, w, r)
		}).
		Methods("PUT")
	a.Router.HandleFunc("/country/{id:[0-9]+}",
		func(w http.ResponseWriter, r *http.Request) {
			handlers.DeleteCountry(a.Counry, w, r)
		}).
		Methods("DELETE")
}
