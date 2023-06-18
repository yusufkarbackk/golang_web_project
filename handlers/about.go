package handlers

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func AboutHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := struct {
		Title string
	}{
		Title: "About",
	}

	tmpl, err := template.ParseFiles("templates/about.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
