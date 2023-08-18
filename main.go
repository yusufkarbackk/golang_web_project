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

	router.ServeFiles("/assets/*filepath", http.Dir("assets"))
	router.ServeFiles("/dashboard_assets/*filepath", http.Dir("dashboard_assets"))

	router.GET("/", middleware.AuthMiddleware(handlers.IndexHandler, Store))
	router.GET("/home", middleware.AuthMiddleware(handlers.IndexHandler, Store))
	// router.GET("/about", middleware.AuthMiddleware(handlers.AdminHandler, Store))
	router.GET("/dashboard-user", middleware.AuthMiddleware(handlers.ShowUsersHandler, Store))
	router.GET("/dashboard-transaksi", middleware.AuthMiddleware(handlers.ShowTransactionsHandler, Store))
	router.GET("/deposit", middleware.AuthMiddleware(handlers.ShowAddDepositHandler, Store))
	router.POST("/deposit", middleware.AuthMiddleware(handlers.SubmitDepositHandler, Store))
	router.GET("/withdraw", middleware.AuthMiddleware(handlers.ShowAddWithdrawFormHandler, Store))
	router.POST("/withdraw", middleware.AuthMiddleware(handlers.SubmitWithdrawHandler, Store))
	router.GET("/add-user", handlers.ShowAddUserFormHandler)
	router.POST("/add-user", handlers.SubmitAddUserFormHandler)
	router.GET("/update-user", middleware.AuthMiddleware(handlers.ShowUpdateUserFormHandler, Store))
	router.PUT("/update-user", middleware.AuthMiddleware(handlers.SubmitUpdateUserHandler, Store))
	router.POST("/delete-user", middleware.AuthMiddleware(handlers.DeleteUserHandler, Store))
	router.GET("/login", handlers.ShowLoginForm)
	router.POST("/login", handlers.SubmitLogin)
	router.GET("/logout", middleware.AuthMiddleware(handlers.Logout, Store))
	router.GET("/error", handlers.ShowErrorPage)

	log.Println("server started on port:8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
