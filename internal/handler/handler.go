package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ttvdmt/url_shortener/internal/config"
	"github.com/ttvdmt/url_shortener/internal/service"
	"github.com/ttvdmt/url_shortener/internal/storage"
)

func Init(storage storage.Storager, cfg config.Config) {
	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		service.Create(w, r, storage, cfg.App.TTL)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		service.Redirect(w, r, storage)
	})
}

func Listen(cfg config.Config) error {
	host := cfg.App.Host
	port := cfg.App.Port

	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), nil); err != nil {
		return fmt.Errorf("server cant start listen: %w", err)
	}

	return nil
}

func CleanUp(storage storage.Storager, cfg config.Config) error {
	period, err := time.ParseDuration(cfg.App.CleanUp_Period)
	if err != nil {
		return fmt.Errorf("cant start cleaning: %w", err)
	}

	go func() {
		timer := time.NewTicker(period)
		for range timer.C {
			if err := storage.CleanUp(); err != nil {
				timer.Stop()
			}
		}
	}()

	return nil
}
