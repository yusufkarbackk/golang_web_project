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
	var user model.User

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	nik := r.PostForm.Get("nik")
	password := r.PostForm.Get("password")

	db := database.MySqlConnection()
	query := "select uuid, nik, nama, password, role, saldo from users where nik = ?"
	stmt, err := db.Prepare(query)

	if err != nil {
		panic(err)
	}

	err = stmt.QueryRow(nik).Scan(&user.Uuid, &user.Nik, &user.Nama, &user.Password, &user.Role, &user.Saldo)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		fmt.Println("nik salah")
		return
	} else {
		fmt.Println("nik benar")
		isPasswordValid := auth.VerifyPassword(password, user.Password)

		if isPasswordValid {
			session, err := auth.Store.Get(r, "login_session")
			if err != nil {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}
			session.Values["authenticated"] = true
			session.Values["nik"] = user.Nik
			session.Values["role"] = user.Role
			session.Values["saldo"] = user.Saldo
			session.Values["nama"] = user.Nama
			err = session.Save(r, w)

			role, ok := session.Values["role"].(string)

			if !ok {
				http.Error(w, "role not found in the session", http.StatusInternalServerError)
				return
			}

			if role == "user" {
				http.Redirect(w, r, "/home", http.StatusSeeOther)
				fmt.Println("login berhasil sebagai user")
			} else {
				http.Redirect(w, r, "/dashboard-user", http.StatusSeeOther)
				fmt.Println("login berhasil sebagai admin")
			}

		} else {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			fmt.Println("password salah")
		}
	}
}

func Logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	session, _ := auth.Store.Get(r, "login_session")

	session.Values["authenticated"] = false
	session.Save(r, w)

	fmt.Println("logout success")
}
