package main

import (
	"fmt"
	"belajar-golang-restapi/helper"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	// "belajar-golang-restapi/app"
	// "belajar-golang-restapi/controller"
	// "belajar-golang-restapi/repository"
	// "belajar-golang-restapi/service"
	"belajar-golang-restapi/middleware"

	// "github.com/go-playground/validator/v10"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr: "localhost:3000",
		Handler: authMiddleware,
	}
}

func main() {
	// db := app.NewDB()
	// validate := validator.New()

	// categoryRepository := repository.NewCategoryRepository()
	// categoryService := service.NewCategoryService(categoryRepository, db, validate)
	// categoryController := controller.NewCategoryController(categoryService)

	// router := app.NewRouter(categoryController)
	// authMiddleware := middleware.NewAuthMiddleware(router)

	// server := NewServer(authMiddleware)

	// fmt.Println("Server is running on port 3000")

	// err := server.ListenAndServe()
	// if err != nil {
	// 	panic(err)
	// }
	server := InitializeServer()
	fmt.Println("Server is running on port 3000")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}