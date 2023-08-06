package main

import (
	"golang_web_Project/auth"
	"golang_web_Project/handlers"

	"golang_web_Project/middleware"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	// "google.golang.org/api/option"
)

func main() {

	router := httprouter.New()
	auth.CreateSession()
	Store := auth.GetSession()

	router.ServeFiles("/assets/*filepath", http.Dir("assets"))
	router.ServeFiles("/dashboard_assets/*filepath", http.Dir("dashboard_assets"))

	router.GET("/", middleware.AuthMiddleware(handlers.IndexHandler, Store))
	router.GET("/home", middleware.AuthMiddleware(handlers.IndexHandler, Store))
	// router.GET("/about", middleware.AuthMiddleware(handlers.AdminHandler, Store))
	router.GET("/dashboard", middleware.AuthMiddleware(handlers.ShowUsersHandler, Store))
	router.GET("/tambah-user", middleware.AuthMiddleware(handlers.ShowAddUserFormHandler, Store))
	router.GET("/login", handlers.ShowLoginForm)
	router.POST("/send_login", handlers.SubmitLogin)
	router.GET("/add-user", handlers.ShowFormHandler)
	router.POST("/add-user", handlers.SubmitFormHandler)

	log.Println("server started on port:8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
