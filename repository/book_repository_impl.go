package repository

import (
	"golang-simple-boilerplate/entity"
	"golang-simple-boilerplate/model"
	"gorm.io/gorm"
)

type BookRepositoryImpl struct {
	Mysql gorm.DB
}

func NewBookRepository(Mysql *gorm.DB) BookRepository {
	return &BookRepositoryImpl{
		Mysql: *Mysql,
	}
}

func (Repository BookRepositoryImpl) CreateBook(Request model.BookRequest) (Response bool, Error error) {
	var book = entity.Book{
		Author: Request.Author,
		Title:  Request.Title,
		Year:   Request.Year,
	}

	if Error = Repository.Mysql.Create(&book).Error; Error != nil {
		return false, Error
	}

	return true, Error
}

func (Repository BookRepositoryImpl) GetAllBook() (Response []model.BookResponse, Error error) {
	var books []entity.Book
	if Error = Repository.Mysql.Find(&books).Error; Error != nil {
		return Response, Error
	}

	var bookResponse []model.BookResponse
	for _, book := range books {
		bookResponse = append(bookResponse, model.BookResponse{
			ID:        book.ID,
			CreatedAt: book.CreatedAt,
			UpdatedAt: book.UpdatedAt,
			Author:    book.Author,
			Title:     book.Title,
			Year:      book.Year,
		})
	}

	Response = bookResponse
	return Response, Error
}

func (Repository BookRepositoryImpl) GetOneBook(id string) (Response model.BookResponse, Error error) {
	var book entity.Book
	if Error = Repository.Mysql.Where("id = ?", id).First(&book).Error; Error != nil {
		return Response, Error
	}
	if book.Title == "" {
		return Response, Error
	}
	Response.ID = book.ID
	Response.CreatedAt = book.CreatedAt
	Response.UpdatedAt = book.UpdatedAt
	Response.Author = book.Author
	Response.Title = book.Title
	Response.Year = book.Year

	return Response, Error
}

func (Repository BookRepositoryImpl) DeleteBook(id string) (Response bool, Error error) {
	var book entity.Book

	if Error = Repository.Mysql.Where("id = ?", id).First(&book).Error; Error != nil {
		return Response, Error
	}

	if book.Title == "" {
		return false, Error
	}

	if Error = Repository.Mysql.Delete(&book).Error; Error != nil {
		return false, Error
	}

	return true, Error
}

func (Repository BookRepositoryImpl) UpdateBook(id string, Request model.BookRequest) (Response model.BookResponse, Error error) {
	book := new(entity.Book)

	if Error = Repository.Mysql.First(&book, id).Error; Error != nil {
		return Response, Error
	}

	if book.Title == "" {
		return Response, Error
	}

	book.Title = Request.Title
	book.Author = Request.Author
	book.Year = Request.Year

	if Error = Repository.Mysql.Save(&book).Error; Error != nil {
		return Response, Error
	}

	Response.ID = book.ID
	Response.CreatedAt = book.CreatedAt
	Response.UpdatedAt = book.UpdatedAt
	Response.Title = book.Title
	Response.Author = book.Author
	Response.Year = book.Year

	return Response, Error
}
