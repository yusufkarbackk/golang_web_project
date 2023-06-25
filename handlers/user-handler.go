package handlers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"golang_web_Project/auth"
	"golang_web_Project/database/user_service"
	"golang_web_Project/model"
	"html/template"
	"net/http"
)

func ShowFormHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tmpl, err := template.ParseFiles("templates/user-form.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	users := userservice.GetUser()
	fmt.Println(users)

	err = tmpl.Execute(w, users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func SubmitFormHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var userData model.User
	err := r.ParseForm()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	hashedPassword, error := auth.HashPassword(r.PostForm.Get("password"))
	if err != nil {
		panic(error)
	}

	userData.Nama = r.PostForm.Get("username")
	userData.Email = r.PostForm.Get("email")
	userData.Password = hashedPassword

	userservice.AddUser(&userData)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
