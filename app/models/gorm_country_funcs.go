package models

//GetСountries comment
func (g *GormDB) GetСountries(nameFilter string) ([]Country, error) {
	cs := []Country{}

	if err := g.DB.Where("name LIKE ?", "%"+nameFilter+"%").Find(&cs).Error; err != nil {
		return nil, err
	}

	return cs, nil
}

//GetCountry comment
func (g *GormDB) GetCountry(c *Country) error {
	return g.DB.Select("name, population").Find(&c).Error
}

//UpdateCountry comment
func (g *GormDB) UpdateCountry(c *Country) error {
	return g.DB.Save(&c).Error
}

//DeleteCountry comment
func (g *GormDB) DeleteCountry(c *Country) error {
	return g.DB.Delete(&c).Error
}

//CreateCountry comment
func (g *GormDB) CreateCountry(c *Country) error {
	err := g.DB.Create(&c).Error
	if err != nil {
		return err
	}
	return nil
}
