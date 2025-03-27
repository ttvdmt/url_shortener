package service

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/ttvdmt/url_shortener/internal/storage"
	urlvalidation "github.com/ttvdmt/url_shortener/internal/url_validation"
)

func Create(w http.ResponseWriter, r *http.Request, storage storage.Storager, default_ttl string) {
	if storage == nil {
		fmt.Println("cant create short url: empty storage")
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	parsed_url := r.FormValue("url")
	if parsed_url == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	if !urlvalidation.IsValid(parsed_url) {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	ttl := r.FormValue("ttl")
	if ttl == "" {
		ttl = default_ttl
	}

	parsed_ttl, err := time.ParseDuration(ttl)
	if err != nil {
		http.Error(w, "Wrong TTL", http.StatusBadRequest)
		return
	}

	code, err := storage.Save(parsed_url, parsed_ttl)
	if err != nil {
		fmt.Println("cant create short url%w", err)
		return
	}

	short_url := fmt.Sprintf("http://%s/%s", r.Host, code)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Short URL: %s\n", short_url)
}

func Redirect(w http.ResponseWriter, r *http.Request, storage storage.Storager) {
	if storage == nil {
		fmt.Println("cant redirect url: empty storage")
		return
	}

	code := r.URL.Path[1:]
	original_url, err := storage.Get(code)
	if err == sql.ErrNoRows {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	if err != nil {
		fmt.Println("cant redirect url: %w", err)
		return
	}

	http.Redirect(w, r, original_url, http.StatusFound)
}
