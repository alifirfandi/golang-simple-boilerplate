package controller

import (
	"github.com/gofiber/fiber/v2"
	"golang-simple-boilerplate/exception"
	"golang-simple-boilerplate/middleware"
	"golang-simple-boilerplate/model"
	"golang-simple-boilerplate/service"
)

type BookController struct {
	BookService service.BookService
}

func NewBookController(BookService *service.BookService) BookController {
	return BookController{
		BookService: *BookService,
	}
}

func (Controller BookController) Route(App fiber.Router) {
	router := App.Group("/book")
	router.Get("/get", middleware.CheckToken(), Controller.AllBook)
	router.Get("/get/:id", middleware.CheckToken(), Controller.OneBook)
	router.Post("/create", middleware.CheckToken(), Controller.CreateBook)
	router.Delete("/delete/:id", middleware.CheckToken(), Controller.DeleteBook)
	router.Put("/update/:id", middleware.CheckToken(), Controller.UpdateBook)
}

func (Controller BookController) CreateBook(c *fiber.Ctx) error {
	book := new(model.BookRequest)
	if err := c.BodyParser(book); err != nil {
		return exception.ErrorHandler(c, err)
	}

	response, verified, err := Controller.BookService.CreateBook(*book)
	if verified {
		return c.Status(fiber.StatusOK).JSON(model.Response{
			Code:   fiber.StatusOK,
			Status: "OK",
			Data:   response,
			Error:  nil,
		})
	}
	return exception.ErrorHandler(c, err)
}

func (Controller BookController) AllBook(c *fiber.Ctx) error {
	response, err := Controller.BookService.GetAllBook()
	if err != nil {
		return exception.ErrorHandler(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(model.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
		Error:  nil,
	})
}

func (Controller BookController) OneBook(c *fiber.Ctx) error {
	response, err := Controller.BookService.GetOneBook(c.Params("id"))

	if err != nil {
		return exception.ErrorHandler(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(model.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
		Error:  nil,
	})
}

func (Controller BookController) DeleteBook(c *fiber.Ctx) error {
	response, err := Controller.BookService.DeleteBook(c.Params("id"))

	if err != nil {
		return exception.ErrorHandler(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(model.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
		Error:  nil,
	})
}

func (Controller BookController) UpdateBook(c *fiber.Ctx) error {
	request := new(model.BookRequest)
	if err := c.BodyParser(request); err != nil {
		return exception.ErrorHandler(c, err)
	}

	response, err := Controller.BookService.UpdateBook(c.Params("id"), model.BookRequest{
		Author: request.Author,
		Title:  request.Title,
		Year:   request.Year,
	})

	if err != nil {
		return exception.ErrorHandler(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(model.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
		Error:  nil,
	})
}
