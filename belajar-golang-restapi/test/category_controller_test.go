package test

import (
	"belajar-golang-restapi/controller"
	"belajar-golang-restapi/service"
	"belajar-golang-restapi/repository"
	"os"
	"database/sql"
	"belajar-golang-restapi/helper"
	"belajar-golang-restapi/middleware"
	"net/http"
	"net/http/httptest"
	"time"
	"github.com/go-playground/validator/v10"
	"belajar-golang-restapi/app"
	"testing"
	"github.com/stretchr/testify/assert"
	"strings"
	"io"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
)

func setupTestDB() *sql.DB {
	password := os.Getenv("password")
	db, err := sql.Open("mysql", "root:"+password+"@tcp(localhost:3306)/golang_test")
	helper.PanicIfError(err)
	
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter() http.Handler {
	db := setupTestDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)
}

func TestCreateCategorySuccess(t *testing.T) {
	router := setupRouter()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", strings.NewReader(`{"name": "Gadget"}`))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusCreated, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 201, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"])
}