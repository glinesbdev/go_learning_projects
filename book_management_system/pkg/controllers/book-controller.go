package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"book_management_system/pkg/models"
	"book_management_system/pkg/utils"
	"github.com/gorilla/mux"
)

var newBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := newBook.AllBooks()
	res, err := json.Marshal(books)

	if utils.ResponseError(w, err) {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	bookId, err := strconv.ParseInt(id, 0, 0)

	if utils.ResponseError(w, err) {
		return
	}

	book, err := models.FindBookById(bookId)

	if utils.ResponseError(w, err) {
		return
	}

	res, err := json.Marshal(book)

	if utils.ResponseError(w, err) {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	utils.ParseBody(r, book)

	book, err := book.CreateBook()

	if utils.ResponseError(w, err) {
		return
	}

	res, err := json.Marshal(book)

	if utils.ResponseError(w, err) {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	utils.ParseBody(r, book)
	id_param := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(id_param, 0, 0)

	if utils.ResponseError(w, err) {
		return
	}

	err = models.UpdateBook(id, book)

	if utils.ResponseError(w, err) {
		return
	}

	res, err := json.Marshal(book)

	if utils.ResponseError(w, err) {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	bookId, err := strconv.ParseInt(id, 0, 0)

	if utils.ResponseError(w, err) {
		return
	}

	models.DeleteBook(bookId)
	w.Header().Set("Content-Type", "application/json")
}
