package main

import (
	"fmt"

	"github.com/ttvdmt/url_shortener/internal/config"
	"github.com/ttvdmt/url_shortener/internal/handler"
	"github.com/ttvdmt/url_shortener/internal/storage"
)

func main() {
	cfg, err := config.Load("../../config/config.yaml")
	if err != nil {
		fmt.Println(err)
	}

	st, err := storage.NewSQLStorage(cfg.Storage.SQLite.DBPath)
	if err != nil {
		fmt.Println(err)
	}
	defer st.Close()

	handler.Init(st, cfg)
	fmt.Println("Server is ready")

	handler.CleanUp(st, cfg)
	handler.Listen(cfg)
}
