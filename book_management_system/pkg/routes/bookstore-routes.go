package routes

import (
	"net/http"

	"book_management_system/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookstoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controllers.GetBooks).Methods(http.MethodGet)
	router.HandleFunc("/books", controllers.CreateBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{id}", controllers.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", controllers.UpdateBook).Methods(http.MethodPut, http.MethodPatch)
	router.HandleFunc("/books/{id}", controllers.DeleteBook).Methods(http.MethodDelete)
}
