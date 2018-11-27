package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

//App comment
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

//Initialize comment
func (a *App) Initialize(host, port, user, password, dbname, sslmode string) {
	log.Println("Connection...")
	connectionString :=
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)

	var err error
	a.DB, err = gorm.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connection successful")

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

//Run comment
func (a *App) Run(port string) {
	log.Fatal(http.ListenAndServe(port, a.Router))
}

func (a *App) initializeRoutes() {

	a.Router.HandleFunc("/", a.getCountries).
		Queries(
			"username", "{username}",
			"email", "{email}",
		).
		Methods("GET")
	a.Router.HandleFunc("/", a.getCountries).Methods("GET")
	a.Router.HandleFunc("/countries", a.getCountries).Methods("GET")
	a.Router.HandleFunc("/country", a.createCountry).Methods("POST")
	a.Router.HandleFunc("/country/{id:[0-9]+}", a.getCountry).Methods("GET")
	a.Router.HandleFunc("/country/{id:[0-9]+}", a.updateCountry).Methods("PUT")
	a.Router.HandleFunc("/country/{id:[0-9]+}", a.deleteCountry).Methods("DELETE")

}

func (a *App) getCountries(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()

	count, _ := strconv.Atoi(v.Get("count"))
	start, _ := strconv.Atoi(v.Get("start"))

	if count > 10 {
		count = 10
	}
	if count < 1 {
		count = 0
	}
	if start < 0 {
		start = 0
	}

	countries, err := getСountriesPage(a.DB, start, count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, countries)
}

func (a *App) createCountry(w http.ResponseWriter, r *http.Request) {
	var c country
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&c); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := c.createCountry(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, c)
}

func (a *App) getCountry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid country ID")
		return
	}

	c := country{ID: id}
	if err := c.getCountry(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Country not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, c)
}

func (a *App) updateCountry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid country ID")
		return
	}

	var c country
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&c); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	c.ID = id

	if err := c.updateCountry(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, c)
}

func (a *App) deleteCountry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Country ID")
		return
	}

	c := country{ID: id}
	if err := c.deleteCountry(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
