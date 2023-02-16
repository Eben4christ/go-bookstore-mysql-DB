package main

import (
	"log"
	"net/http"

	//_ "github.com.jinzhu/gorm/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	
	"github.com/eben4christ/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main(){
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}