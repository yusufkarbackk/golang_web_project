package handlers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"golang_web_Project/database/user_service"
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
	err := r.ParseForm()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	fmt.Println(username)
	fmt.Println(password)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
