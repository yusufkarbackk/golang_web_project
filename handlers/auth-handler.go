package handlers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"golang_web_Project/auth"
	"golang_web_Project/database"
	"golang_web_Project/model"
	"html/template"
	"net/http"
	
)

func ShowLoginForm(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tmpl, err := template.ParseFiles("templates/login.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = tmpl.Execute(w, nil)
}

func SubmitLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	nama := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	db := database.MySqlConnection()
	query := "select id, nama, email, password from users where nama = ?"
	stmt, err := db.Prepare(query)
	if err != nil {
		panic(err)
	}

	var user model.User

	err = stmt.QueryRow(nama).Scan(&user.Id, &user.Nama, &user.Email, &user.Password)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		fmt.Println("username salah")
		return
	} else {
		fmt.Println("username benar")
		isPasswordValid := auth.VerifyPassword(password, user.Password)

		if isPasswordValid {
			session, err := auth.Store.Get(r, "login_session")
			if err != nil {
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
			session.Values["authenticated"] = true

			err = session.Save(r, w)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			fmt.Println("login berhasil")
		} else {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			fmt.Println("password salah")
		}
	}
}

func ShowRegisterForm(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tmpl, err := template.ParseFiles("templates/login.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = tmpl.Execute(w, nil)
}
