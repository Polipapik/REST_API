package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type country struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Population int64  `json:"population"`
}

func getСountriesPage(db *gorm.DB, start, count int) ([]country, error) {
	if count == 0 {
		return getСountries(db)
	}

	rows, err := db.DB().Query(
		"SELECT id, name, population FROM countries LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	countries := []country{}

	for rows.Next() {
		var c country
		if err := rows.Scan(&c.ID, &c.Name, &c.Population); err != nil {
			return nil, err
		}
		countries = append(countries, c)
	}

	return countries, nil
}

func getСountries(db *gorm.DB) ([]country, error) {
	rows, err := db.DB().Query("SELECT * FROM countries")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	countries := []country{}

	for rows.Next() {
		var c country
		if err := rows.Scan(&c.ID, &c.Name, &c.Population); err != nil {
			return nil, err
		}
		countries = append(countries, c)
	}

	return countries, nil
}

func (c *country) getCountry(db *gorm.DB) error {
	return db.DB().QueryRow("SELECT name, population FROM countries WHERE id=$1",
		c.ID).Scan(&c.Name, &c.Population)
}

func (c *country) updateCountry(db *gorm.DB) error {
	_, err :=
		db.DB().Exec("UPDATE countries SET name=$1, population=$2 WHERE id=$3",
			c.Name, c.Population, c.ID)

	return err
}

func (c *country) deleteCountry(db *gorm.DB) error {
	_, err := db.DB().Exec("DELETE FROM countries WHERE id=$1", c.ID)

	return err
}

func (c *country) createCountry(db *gorm.DB) error {
	err := db.DB().QueryRow(
		"INSERT INTO countries(name, population) VALUES($1, $2) RETURNING id",
		c.Name, c.Population).Scan(&c.ID)

	if err != nil {
		return err
	}

	return nil
}
