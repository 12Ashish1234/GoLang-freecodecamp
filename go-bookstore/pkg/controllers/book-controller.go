package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/12Ashish1234/GoLang-freecodecamp/go-bookstore/pkg/models"
	"github.com/12Ashish1234/GoLang-freecodecamp/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	// marshall is used convert data from database to json format
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK) // 200
	// Helps to send the response to frontend or postman.
	// here we are sending the "res" variable. The json of newBooks will be sent
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// we need to access the book ID. Hence the below line
	bookId := vars["bookId"]

	// Here parseint is used because the value from the bookID variable is in string type. strconv.ParseInt will convert it to int64 type.
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	// The GetBookByID function in models package return two variables. first one of type `Book`` and another one `db`.
	// Here we dont need the db variable. Hence we are using blank character for it `_`
	bookDetails, _ := models.GetBookById(ID)

	res, _ := json.Marshal(bookDetails)
	// ???? why is Header set in some places and not all?
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	// try to remove the & and see at the end what would happen!
	CreateBook := &models.Book{}
	// ParseBody is used beause we get a content from the user. So now we want to parse it so that the database can understand it
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()

	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Initial steps for delete will resemble the GetBookByID function. because we need to get the bookId and convert it to int64
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// ???? why is the below line used only some places?
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// creating an instance of book called updateBook
	var updateBook = &models.Book{}

	// below ParseBody will transfer the contents in request `r` to updateBook. now updateBook will have contents of the request body which we give in postman
	utils.ParseBody(r, updateBook)

	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	bookDetails, db := models.GetBookById(ID)

	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}

	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}

	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	// The new data has been updated in bookDetails variable. now we will save the details to the database. hence the below line
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
