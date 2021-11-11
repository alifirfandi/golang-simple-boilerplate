package main

import (
	"golang-simple-boilerplate/config"
	"golang-simple-boilerplate/controller"
	"golang-simple-boilerplate/exception"
	"golang-simple-boilerplate/repository"
	"golang-simple-boilerplate/service"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	mysql := config.MysqlConnection()

	authRepository := repository.NewAuthRepository(mysql)
	userRepository := repository.NewUserRepository(mysql)
	bookRepository := repository.NewBookRepository(mysql)

	authService := service.NewAuthService(&authRepository)
	userService := service.NewUserService(&userRepository)
	bookService := service.NewBookService(&bookRepository)

	authController := controller.NewAuthController(&authService)
	userController := controller.NewUserController(&userService)
	bookController := controller.NewBookController(&bookService)

	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	v1 := app.Group("/api/v1")
	authController.Route(v1)
	userController.Route(v1)
	bookController.Route(v1)

	// Start App
	err := app.Listen(os.Getenv("PORT"))
	exception.PanicIfNeeded(err)
}
