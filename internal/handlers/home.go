package handlers

import (
	"html/template"
	"net/http"
)

type PageData struct {
	StatusCode int
	StatusText string
	Text       string
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		errorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}
	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		errorHandler(w,r,http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, PageData{})
}
