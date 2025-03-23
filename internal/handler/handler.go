package handler

import (
	"fmt"
	"net/http"

	"github.com/ttvdmt/url_shortener/internal/service"
	"github.com/ttvdmt/url_shortener/internal/storage"
)

func Init(storage storage.Storager) {
	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		service.Create(w, r, storage)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		service.Redirect(w, r, storage)
	})
}

func Listen(port string) error {
	if err := http.ListenAndServe(port, nil); err != nil {
		return fmt.Errorf("server cant start listen: %w", err)
	}

	return nil
}
