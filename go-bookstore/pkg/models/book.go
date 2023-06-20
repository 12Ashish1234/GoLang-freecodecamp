package models

import (
	"github.com/12Ashish1234/GoLang-freecodecamp/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()

	// Auto migrate will automatically migrate your schema, to keep your schema up to date.
	// Auto migrate will create tables, missing foreign keys, constraints, columns and indexes. It will change existing column's type if it's size, precision, nullable changed.
	// it WILL NOT delete unused columns to protect your data.
	// Refer:- https://t.ly/NbZ2
	db.AutoMigrate(&Book{})
}
