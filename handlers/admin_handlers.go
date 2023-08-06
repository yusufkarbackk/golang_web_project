package handlers

import (
	"fmt"
	userservice "golang_web_Project/database/user_service"
	"golang_web_Project/model"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
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
		Users: userservice.GetUser(),
	}
	fmt.Println(data)
	err = tmpl.Execute(w, data)
}

func ShowAddUserFormHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tmpl, err := template.ParseFiles("templates/tambahPengguna.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
}
