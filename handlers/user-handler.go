package handlers

import (
	"fmt"
	"golang_web_Project/auth"
	"golang_web_Project/database/user_service"
	"golang_web_Project/model"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func ShowFormHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	tmpl, err := template.ParseFiles("templates/user-form.html")

	users := userservice.GetUser()
	fmt.Println(users)

	err = tmpl.Execute(w, users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// if err != nil {
// 	http.Redirect(w, r, "/", http.StatusSeeOther)
// 	fmt.Println("Session habis")
// } else {
// 	fmt.Println("session valid")
// 	fmt.Println(session)
// 	fmt.Println(session.IsNew)
// }

// if session.IsNew {
// 	fmt.Println("session valid")
// 	fmt.Println(session)
// } else {
// 	fmt.Println("Session habis")
// }

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

	userData.Nama = r.PostForm.Get("nama")
	userData.Nik = r.PostForm.Get("nik")
	userData.Password = hashedPassword

	userservice.AddUser(&userData)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
