package handlers

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := struct {
		Title string
	}{
		Title: "Home golang",
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = tmpl.Execute(w, data)
}
