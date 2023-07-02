package handlers

import (
	"net/http"
	"text/template"
)

type Er struct {
	Status int
	text   string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, http.StatusNotFound)
		return
	}
	tmpl := template.Must(template.ParseFiles("frontend/templates/index.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ErrorHandler(w http.ResponseWriter, i int) {
	w.WriteHeader(i)
	tmpl, err := template.ParseFiles("frontend/templates/error.html")
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	er := Er{i, http.StatusText(i)}
	err = tmpl.Execute(w, er)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
}
