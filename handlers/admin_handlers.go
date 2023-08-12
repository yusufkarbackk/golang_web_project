package handlers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"golang_web_Project/auth"
	userservice "golang_web_Project/database/user_service"
	"golang_web_Project/model"
	"html/template"
	"net/http"
	"strconv"
)

func ShowUsersHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	tmpl, err := template.ParseFiles("templates/dashboardAdmin.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := struct {
		Users []model.UserNoPassword
	}{
		Users: userservice.GetUsers(),
	}
	fmt.Println(data)
	err = tmpl.Execute(w, data)
}

func ShowAddUserFormHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	baseTmpl, err := template.ParseFiles("templates/tambahPengguna.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = baseTmpl.Execute(w, nil)
}

func SubmitAddUserFormHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var userData model.User
	err := r.ParseForm()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	hashedPassword, error := auth.HashPassword(r.PostForm.Get("nik"))
	if err != nil {
		panic(error)
	}

	userData.Nama = r.PostForm.Get("nama")
	userData.Nik = r.PostForm.Get("nik")
	userData.Jenis_kelamin = r.PostForm.Get("gender")
	userData.Alamat = r.PostForm.Get("alamat")
	userData.Password = hashedPassword

	userservice.AddUser(&userData)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func ShowUpdateUserFormHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	baseTmpl, err := template.ParseFiles("templates/editPengguna.html")
	id, err := strconv.Atoi(r.URL.Query().Get("uuid"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := userservice.GetUser(id)
	fmt.Println(id)
	err = baseTmpl.Execute(w, data)
}
