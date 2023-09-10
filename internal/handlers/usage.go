package handlers

import (
	"html/template"
	"net/http"
)

func Usage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/usage" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		errorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}
	tmpl, err := template.ParseFiles("templates/usage.html")
	if err != nil {
		errorHandler(w,r,http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, PageData{})
}
