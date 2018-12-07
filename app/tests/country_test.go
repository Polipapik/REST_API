package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Polipapik/REST_API/app/handlers"
	"github.com/Polipapik/REST_API/app/models"
	"github.com/stretchr/testify/assert"
)

func TestGetCountriesHandler(t *testing.T) {
	tmpcountry := models.Country{ID: 1, Name: "niceCHELIKI", Population: 1307}

	var m models.MockDB
	m.On("GetСountries").Return([]models.Country{
		tmpcountry}, nil).Once()

	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.GetCountries(&m, w, r)
	})

	hf.ServeHTTP(recorder, req)
	//t.Log("wowowowowo")

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := tmpcountry

	cs := []models.Country{}

	err = json.NewDecoder(recorder.Body).Decode(&cs)
	if err != nil {
		t.Fatal(err)
	}

	actual := cs[0]

	if !assert.True(t, (actual == expected)) {
		t.Errorf("Handler returned unexpected body: got %v want %v", actual, expected)
	}
	m.AssertExpectations(t)
}
