package handlers

import (
	"database/sql"
	"github.com/spayder/url-shortener/internal/db"
	"github.com/spayder/url-shortener/internal/url"
	"html/template"
	"net/http"
	"strings"
)

func ShortenHandler(sqlite *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		originalUrl := r.PostFormValue("url")
		if originalUrl == "" {
			http.Error(w, "Missing url", http.StatusBadRequest)
		}

		if !strings.HasPrefix(originalUrl, "http://") || !strings.HasPrefix(originalUrl, "https://") {
			originalUrl = "https://" + originalUrl
		}

		hashedUrl := url.Shorten(originalUrl)

		if err := db.CreateURL(sqlite, hashedUrl, originalUrl); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		data := map[string]string{
			"ShortURL": hashedUrl,
		}

		t, err := template.ParseFiles("internal/views/shorten.html")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = t.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func Proxy(sqlite *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shortUrl := r.URL.Path[1:]

		if shortUrl == "" {
			http.Error(w, "Missing url", http.StatusBadRequest)
			return
		}

		originalURL, err := db.GetOriginURL(sqlite, shortUrl)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Redirect(w, r, originalURL, http.StatusPermanentRedirect)
	}
}
