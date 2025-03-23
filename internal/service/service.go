package service

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/ttvdmt/url_shortener/internal/storage"
)

func Create(w http.ResponseWriter, r *http.Request, storage storage.Storager) {
	if storage == nil {
		fmt.Println("empty storage")
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	originalURL := r.FormValue("url")
	if originalURL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	code, err := storage.Save(originalURL)
	if err != nil {
		fmt.Println("cant create short URL%w", err)
		return
	}
	shortURL := fmt.Sprintf("http://%s/%s", r.Host, code)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Short URL: %s\n", shortURL)
}

func Redirect(w http.ResponseWriter, r *http.Request, storage storage.Storager) {
	if storage == nil {
		fmt.Println("empty storage")
		return
	}

	code := r.URL.Path[1:]
	originalURL, err := storage.Get(code)
	if err == sql.ErrNoRows {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	if err != nil {
		fmt.Println("cant redirect url: %w", err)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}
