package repository

import "golang-simple-boilerplate/model"

type BookRepository interface {
	CreateBook(Request model.BookRequest) (Response bool, Error error)
	GetAllBook() (Response []model.BookResponse, Error error)
	GetOneBook(id string) (Response model.BookResponse, Error error)
	DeleteBook(id string) (Response bool, Error error)
	UpdateBook(id string, Request model.BookRequest) (Response model.BookResponse, Error error)
}
