package service

import (
	"fmt"
	"net/http"

	"github.com/ttvdmt/url_shortener/internal/storage"
)

func Create(w http.ResponseWriter, r *http.Request, storage *storage.Storage) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	originalURL := r.FormValue("url")
	if originalURL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	code := storage.Save(originalURL)
	shortURL := fmt.Sprintf("http://%s/%s", r.Host, code)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Short URL: %s\n", shortURL)
}

func Redirect(w http.ResponseWriter, r *http.Request, storage *storage.Storage) {
	code := r.URL.Path[1:]
	originalURL, exists := storage.Get(code)
	if !exists {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}
