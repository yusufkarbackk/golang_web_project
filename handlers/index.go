package handlers

import (
	"golang_web_Project/auth"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	session, err := auth.Store.Get(r, "login_session")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	data := struct {
		Nama  string
		Saldo int
	}{
		Nama:  session.Values["nama"].(string),
		Saldo: session.Values["saldo"].(int),
	}
	tmpl, err := template.ParseFiles("templates/dashboard.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = tmpl.Execute(w, data)
}
