package handlers

import (
	"html/template"
	"net/http"
)

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	text := http.StatusText(status)

	tmpl := template.Must(template.ParseFiles("templates/error.html"))

	tmpl.Execute(w, PageData{StatusCode: status, StatusText: text})
}
