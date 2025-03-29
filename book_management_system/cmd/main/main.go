package main

import (
	"log"
	"net/http"

	"book_management_system/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookstoreRoutes(router)
	http.Handle("/", router)

	log.Println("Server started on: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
