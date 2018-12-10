package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"

	"github.com/Polipapik/REST_API/app/models"
	"github.com/Polipapik/REST_API/app/utils"
	"github.com/gorilla/mux"
)

//GetCountries comment OK
func GetCountries(db models.CountryAPI, w http.ResponseWriter, r *http.Request) {

	v := r.URL.Query()
	name := v.Get("name")

	cs, err := db.GetСountries(name)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, cs)
}

//GetCountry comment
func GetCountry(db models.CountryAPI, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid country ID")
		return
	}

	c := models.Country{ID: id}
	if err := db.GetCountry(&c); err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			utils.RespondWithError(w, http.StatusNotFound, "Country not found")
		default:
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, c)
}

//UpdateCountry comment
func UpdateCountry(db models.CountryAPI, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid country ID")
		return
	}

	var c models.Country
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&c); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	c.ID = id

	if err := db.UpdateCountry(&c); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, c)
}

//DeleteCountry comment
func DeleteCountry(db models.CountryAPI, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Country ID")
		return
	}

	c := models.Country{ID: id}
	if err := db.DeleteCountry(&c); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

//CreateCountry comment
func CreateCountry(db models.CountryAPI, w http.ResponseWriter, r *http.Request) {
	var c models.Country
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&c); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := db.CreateCountry(&c); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, c)
}
