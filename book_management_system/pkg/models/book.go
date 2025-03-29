package models

import (
	"book_management_system/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"Name"`
	Author      string `json:"Author"`
	Publication string `json:"Publication"`
}

func init() {
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&Book{})
}

func (book *Book) CreateBook() (*Book, error) {
	db.NewRecord(book)
	db.Create(&book)

	if db.Error != nil {
		return nil, db.Error
	}

	return book, nil
}

func (book *Book) AllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func UpdateBook(id int64, book *Book) error {
	foundBook := db.First(book, id)

	if foundBook.Error != nil {
		return foundBook.Error
	}

	db.Model(&foundBook).Updates(book)

	return nil
}

func FindBookById(id int64) (*Book, error) {
	var book Book
	db.Where("ID = ?", id).Find(&book)

	if db.Error != nil {
		return nil, db.Error
	}

	return &book, nil
}

func DeleteBook(id int64) Book {
	var book Book
	db.Where("ID = ?", id).Delete(book)
	return book
}
