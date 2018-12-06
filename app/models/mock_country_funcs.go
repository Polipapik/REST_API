package models

import (
	"log"
)

//GetСountries comment
func (m *MockDB) GetСountries() ([]Country, error) {
	// tmpcountry := Country{ID: 1, Name: "niceCHELIKI", Population: 1307}
	// countries := []Country{
	// 	tmpcountry}
	// return countries, nil

	rets := m.Called()
	log.Println(m)
	return rets.Get(0).([]Country), rets.Error(1)
}

//GetCountry comment
func (m *MockDB) GetCountry(c *Country) error {
	return m.Called(c).Error(0)
}

//UpdateCountry comment
func (m *MockDB) UpdateCountry(c *Country) error {
	return m.Called(c).Error(0)
}

//DeleteCountry comment
func (m *MockDB) DeleteCountry(c *Country) error {
	return m.Called(c).Error(0)
}

//CreateCountry comment
func (m *MockDB) CreateCountry(c *Country) error {
	return m.Called(c).Error(0)
}
