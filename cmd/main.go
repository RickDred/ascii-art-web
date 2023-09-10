package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/RickDred/ascii-art-web/internal/handlers"
)

const (
	Addr = ":8000"
)

func main() {
	port := "3000"
	if val, ok := os.LookupEnv("PORT"); ok {
		port = val
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/ascii-art/export", handlers.Export)
	http.HandleFunc("/ascii-art", handlers.AsciiArt)
	http.HandleFunc("/usage", handlers.Usage)
	http.HandleFunc("/team", handlers.AboutUs)

	fmt.Println("Server starting on http://localhost:" + port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
