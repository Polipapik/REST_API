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
	countries := []models.Country{
		{ID: 1, Name: "niceCHELIKI", Population: 1307},
		{ID: 2, Name: "trueGays", Population: 228322}}

	var m models.MockDB
	m.On("GetСountries").Return(countries, nil).Once()

	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.GetCountries(&m, w, r)
	})
	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := countries

	cs := []models.Country{}
	if err = json.NewDecoder(recorder.Body).Decode(&cs); err != nil {
		t.Fatal(err)
	}
	actual := cs

	rly := true
	if len(actual) == len(expected) {
		for i := 0; i < len(actual); i++ {
			if actual[i] != expected[i] {
				rly = false
			}
		}
	}

	if !assert.True(t, rly) {
		t.Errorf("Handler returned unexpected body: got %v want %v", actual, expected)
	}
	m.AssertExpectations(t)
}
