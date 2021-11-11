package service

import (
	"golang-simple-boilerplate/model"
	"golang-simple-boilerplate/repository"
	"golang-simple-boilerplate/validation"
)

type BookServiceImpl struct {
	BookRepository repository.BookRepository
}

func NewBookService(BookRepository *repository.BookRepository) BookService {
	return &BookServiceImpl{
		BookRepository: *BookRepository,
	}
}

func (Service BookServiceImpl) CreateBook(Request model.BookRequest) (Response bool, Verified bool, Error error) {
	if Error = validation.BookValidation(Request); Error != nil {
		return Response, Verified, Error
	}
	Verified = true
	Response, Error = Service.BookRepository.CreateBook(Request)
	return Response, Verified, Error
}

func (Service BookServiceImpl) GetAllBook() (Response []model.BookResponse, Error error) {
	Response, Error = Service.BookRepository.GetAllBook()
	return Response, Error
}

func (Service BookServiceImpl) GetOneBook(id string) (Response model.BookResponse, Error error) {
	Response, Error = Service.BookRepository.GetOneBook(id)
	return Response, Error
}

func (Service BookServiceImpl) DeleteBook(id string) (Response bool, Error error) {
	Response, Error = Service.BookRepository.DeleteBook(id)
	return Response, Error
}

func (Service BookServiceImpl) UpdateBook(id string, Request model.BookRequest) (Response model.BookResponse, Error error) {
	Response, Error = Service.BookRepository.UpdateBook(id, Request)
	return Response, Error
}
