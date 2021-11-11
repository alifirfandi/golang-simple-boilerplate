package service

import "golang-simple-boilerplate/model"

type BookService interface {
	CreateBook(Request model.BookRequest) (Response bool, Verified bool, Error error)
	GetAllBook() (Response []model.BookResponse, Error error)
	GetOneBook(Request string) (Response model.BookResponse, bookExist bool, Error error)
	DeleteBook(id string) (Response bool, bookExist bool, Error error)
	UpdateBook(id string, Request model.BookRequest) (Response model.BookResponse, bookExist bool, Error error)
}
