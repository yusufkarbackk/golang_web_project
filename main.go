package main

import (
	"golang_web_Project/auth"
	"golang_web_Project/handlers"
	"golang_web_Project/middleware"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)


func main() {

	router := httprouter.New()
	auth.CreateSession()
	Store := auth.GetSession()

	router.GET("/", handlers.IndexHandler)
	router.GET("/about", handlers.AboutHandler)
	router.GET("/login", handlers.ShowLoginForm)
	router.POST("/login", handlers.SubmitLogin)
	router.GET("/add-user", middleware.FormMiddleware(handlers.ShowFormHandler, Store))
	router.POST("/add-user", handlers.SubmitFormHandler)
	router.ServeFiles("/static/*filepath", http.Dir("static"))

	log.Println("server started on port:8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
