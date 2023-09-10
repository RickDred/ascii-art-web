package handlers

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/RickDred/ascii-art-web/internal/art"
	"github.com/RickDred/ascii-art-web/pkg/validator"
)

const (
	standardPath   = "./static/fonts/standard.txt"
	shadowPath     = "./static/fonts/shadow.txt"
	thinkertoyPath = "./static/fonts/thinkertoy.txt"
)

func AsciiArt(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		tmpl, err := template.ParseFiles("templates/asciiArt.html")
		if err != nil {
			errorHandler(w, r, http.StatusInternalServerError)
			return
		}

		asciiArt := art.Ascii("Example", standardPath)

		tmpl.Execute(w, PageData{Text: asciiArt})

	case http.MethodPost:
		tmpl, err := template.ParseFiles("templates/asciiArt.html")
		if err != nil {
			errorHandler(w, r, http.StatusInternalServerError)
			return
		}
		banner := r.FormValue("banner")
		inputText := r.FormValue("inputText")
		inputText = strings.ReplaceAll(inputText, "\r", "")

		// validator
		v := validator.New()

		// available options
		availableFonts := []string{"standard", "shadow", "thinkertoy"}

		if ok := v.IsTextAscii(inputText); !ok {
			errorHandler(w, r, http.StatusBadRequest)
			return
		}

		if ok := v.IsArrayContainsString(availableFonts, banner); !ok {
			errorHandler(w, r, http.StatusBadRequest)
			return
		}

		pathToBanner := ""

		switch banner {
		case "standard":
			pathToBanner = standardPath
		case "shadow":
			pathToBanner = shadowPath
		case "thinkertoy":
			pathToBanner = thinkertoyPath
		}

		asciiArt := art.Ascii(inputText, pathToBanner)

		tmpl.Execute(w, PageData{Text: asciiArt})
	default:
		errorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}
}
