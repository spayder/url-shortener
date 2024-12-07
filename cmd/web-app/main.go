package main

import (
	"database/sql"
	"github.com/spayder/url-shortener/internal/db"
	"github.com/spayder/url-shortener/internal/handlers"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	sqlite, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer sqlite.Close()

	err = db.CreateTable(sqlite)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			handlers.Show(w, r)
		} else {
			handlers.Proxy(sqlite)(w, r)
		}
	})
	http.HandleFunc("/shorten", handlers.ShortenHandler(sqlite))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
