package handlers

import (
	"github.com/spayder/url-shortener/internal/url"
	"html/template"
	"net/http"
	"strings"
)

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	originalUrl := r.PostFormValue("url")
	if originalUrl == "" {
		http.Error(w, "Missing url", http.StatusBadRequest)
	}

	if strings.HasPrefix(originalUrl, "http://") || strings.HasPrefix(originalUrl, "https://") {
		originalUrl = "https://" + originalUrl
	}

	hashedUrl := url.Shorten(originalUrl)

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
