package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"golang_web_Project/database/user_service"
	"golang_web_Project/handlers"
	"log"
	"net/http"
)

func main() {

	userservice.GetUser()

	router := httprouter.New()

	router.GET("/", handlers.IndexHandler)
	router.GET("/about", handlers.AboutHandler)
	router.GET("/add-user", handlers.ShowFormHandler)
	router.POST("/add-user", handlers.SubmitFormHandler)
	router.ServeFiles("/static/*filepath", http.Dir("static"))

	log.Println("server started on port:8080")
	log.Fatal(http.ListenAndServe(":8000", router))
}
