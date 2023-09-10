package handlers

import (
	"html/template"
	"net/http"
)

func AboutUs(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/team" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		errorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}
	tmpl, err := template.ParseFiles("templates/aboutUs.html")
	if err != nil {
		errorHandler(w,r,http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, PageData{})
}
