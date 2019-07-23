package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Recipe struct {
	gorm.Model
	Text string
	Picture []byte
}
