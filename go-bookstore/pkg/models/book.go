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

// Now we are creating model functions to make changes in the database
// We receive variable of type Book b, and we return b of type book after creating the book
func (b *Book) CreateBook() *Book {
	// The functions given by gorm will create the SQL queries. We just need to use those functions as shown below
	db.NewRecord(b)
	db.Create((&b))
	return b
}

// Here []Book denotes the return type. As this is GetAllBooks function, we will be returning all the books in the Database.
func GetAllBooks() []Book {
	// creating variable Books which is of type slice of Book (Denoted by []Book)
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
