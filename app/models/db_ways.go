package models

import (
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/mock"
)

//GormDB comment
type GormDB struct {
	DB *gorm.DB
}

//MockDB comment
type MockDB struct {
	mock.Mock
}
