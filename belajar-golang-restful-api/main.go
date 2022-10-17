package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"portofolio/belajar-golang-restful-api/app"
	"portofolio/belajar-golang-restful-api/controller"
	"portofolio/belajar-golang-restful-api/helper"
	"portofolio/belajar-golang-restful-api/middleware"
	"portofolio/belajar-golang-restful-api/repository"
	"portofolio/belajar-golang-restful-api/service"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
