package main

import (
	"belajar-golang-resful-api/app"
	"belajar-golang-resful-api/controller"
	"belajar-golang-resful-api/helper"
	"belajar-golang-resful-api/repository"
	"belajar-golang-resful-api/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
)

func main() {

	validate := validator.New()

	db := app.NewDB()

	categoryRepo := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepo, db, validate)
	CategoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", CategoryController.FindAll)
	router.GET("/api/categories/:categoryId", CategoryController.FindbyId)
	router.POST("/api/categories", CategoryController.Create)
	router.PUT("/api/categories/:categoryId", CategoryController.Update)
	router.DELETE("/api/categories/:categoryId", CategoryController.Delete)

	server := http.Server{
		Addr:    "localhost:8000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
