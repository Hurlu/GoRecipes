package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Ingredient struct {
	gorm.Model
	Name string
	Unit Unit
}