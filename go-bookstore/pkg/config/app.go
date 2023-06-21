/*
*	Purpose of this file is for returning a db variable so that other files can talk to the database
 */
package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	// This is for opening connection to mysql. ashish is the username, ashishd@57@ is the password and simplerest is the table name
	// d, err := gorm.Open("mysql", "root:Ashu2124#/simplerest?charset=utf8&parseTime=True&loc=Local")
	dsn := "root:Ashu2124#@tcp(127.0.0.1:3306)/simplerest?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}
	// now whatever is there in d, we will transfer it to db variable
	db = d
}

func GetDB() *gorm.DB {
	return db
}
