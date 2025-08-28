package main

import (
	"fmt"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"belajar-golang-restapi/app"
	"belajar-golang-restapi/controller"
	"belajar-golang-restapi/repository"
	"belajar-golang-restapi/service"
	"belajar-golang-restapi/middleware"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr: "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	fmt.Println("Server is running on port 3000")

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}