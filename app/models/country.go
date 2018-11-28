package models

import "github.com/jinzhu/gorm"

//Country comment
type Country struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Population int64  `json:"population"`
}

//GetСountries comment
func GetСountries(db *gorm.DB, start, count int) ([]Country, error) {
	countries := []Country{}
	var err error

	if count == 0 {
		err = db.Find(&countries).Error
	} else {
		err = db.Offset(start).Limit(count).Find(&countries).Error
	}

	if err != nil {
		return nil, err
	}

	return countries, nil
}

//GetCountry comment
func (c *Country) GetCountry(db *gorm.DB) error {
	return db.Select("name, population").Find(&c).Error
}

//UpdateCountry comment
func (c *Country) UpdateCountry(db *gorm.DB) error {
	return db.Save(&c).Error
}

//DeleteCountry comment
func (c *Country) DeleteCountry(db *gorm.DB) error {
	return db.Delete(&c).Error
}

//CreateCountry comment
func (c *Country) CreateCountry(db *gorm.DB) error {
	err := db.Create(&c).Error
	if err != nil {
		return err
	}
	return nil
}
