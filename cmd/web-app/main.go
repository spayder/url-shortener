package main

import (
	"github.com/spayder/url-shortener/internal/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.Show)
	http.HandleFunc("/shorten", handlers.ShortenHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
