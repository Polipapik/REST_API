package models

//Country comment
type Country struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Population int64  `json:"population"`
}

//CountryAPI comment
type CountryAPI interface {
	GetСountries() ([]Country, error)
	GetCountry(c *Country) error
	UpdateCountry(c *Country) error
	DeleteCountry(c *Country) error
	CreateCountry(c *Country) error
}
