package handlers

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func ShowErrorPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tmpl, err := template.ParseFiles("templates/errorPage.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = tmpl.Execute(w, nil)
}
