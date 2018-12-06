package models

//GetСountries comment
func (gorm *GormDB) GetСountries(start, count int) ([]Country, error) {
	countries := []Country{}
	var err error

	if count == 0 {
		err = gorm.DB.Find(&countries).Error
	} else {
		err = gorm.DB.Offset(start).Limit(count).Find(&countries).Error
	}

	if err != nil {
		return nil, err
	}

	return countries, nil
}

//GetCountry comment
func (gorm *GormDB) GetCountry(c *Country) error {
	return gorm.DB.Select("name, population").Find(&c).Error
}

//UpdateCountry comment
func (gorm *GormDB) UpdateCountry(c *Country) error {
	return gorm.DB.Save(&c).Error
}

//DeleteCountry comment
func (gorm *GormDB) DeleteCountry(c *Country) error {
	return gorm.DB.Delete(&c).Error
}

//CreateCountry comment
func (gorm *GormDB) CreateCountry(c *Country) error {
	err := gorm.DB.Create(&c).Error
	if err != nil {
		return err
	}
	return nil
}
