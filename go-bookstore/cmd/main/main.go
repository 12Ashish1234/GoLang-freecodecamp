package main

/*
*	in main we are going to do two things:
*	1. create the server localhost.
*	2. tell golang where our routers reside i.e., bookstore-routes.go file
 */

import (
	"log"
	"net/http"

	"github.com/12Ashish1234/GoLang-freecodecamp/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	// http.ListenAndServe() helps create the server in mentioned port. In this example localhost:9010
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
