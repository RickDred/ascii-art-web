package handlers

import (
	"io/ioutil"
	"net/http"
	"os"
)

func Export(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art/export" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		errorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		errorHandler(w, r, http.StatusInternalServerError)
	}

	text := r.FormValue("result")

	tmpfile, err := ioutil.TempFile("", "ascii-art.txt")
	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError)
		return
	}
	defer tmpfile.Close()

	if _, err = tmpfile.WriteString(text); err != nil {
		errorHandler(w, r, http.StatusInternalServerError)
		return
	}

	if err = tmpfile.Chmod(os.FileMode(0644)); err != nil {
		errorHandler(w, r, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename=ascii-art.txt")
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))

	http.ServeFile(w, r, tmpfile.Name())
}
