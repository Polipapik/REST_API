package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Polipapik/REST_API/app/models"
	"github.com/Polipapik/REST_API/app/util"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

//GetCountries comment
func GetCountries(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
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

	countries, err := models.GetСountries(db, start, count)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusOK, countries)
}

//CreateCountry comment
func CreateCountry(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var c models.Country
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&c); err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := c.CreateCountry(db); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusCreated, c)
}

//GetCountry comment
func GetCountry(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid country ID")
		return
	}

	c := models.Country{ID: id}
	if err := c.GetCountry(db); err != nil {
		switch err {
		case sql.ErrNoRows:
			util.RespondWithError(w, http.StatusNotFound, "Country not found")
		default:
			util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	util.RespondWithJSON(w, http.StatusOK, c)
}

//UpdateCountry comment
func UpdateCountry(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid country ID")
		return
	}

	var c models.Country
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&c); err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	c.ID = id

	if err := c.UpdateCountry(db); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusOK, c)
}

//DeleteCountry comment
func DeleteCountry(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid Country ID")
		return
	}

	c := models.Country{ID: id}
	if err := c.DeleteCountry(db); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
