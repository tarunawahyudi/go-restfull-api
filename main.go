package main

import (
	"belajar-golang-restfull-api/app"
	"belajar-golang-restfull-api/controller"
	"belajar-golang-restfull-api/exception"
	"belajar-golang-restfull-api/helper"
	"belajar-golang-restfull-api/middleware"
	"belajar-golang-restfull-api/repository"
	"belajar-golang-restfull-api/service"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {

	validate := validator.New()
	db := app.NewDB()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    ":3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
