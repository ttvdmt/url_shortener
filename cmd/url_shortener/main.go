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

	st, err := storage.NewSQLStorage(cfg.Database)
	if err != nil {
		fmt.Println(err)
	}
	defer st.Close()

	handler.Init(st)
	fmt.Println("Server is ready")

	handler.Listen(cfg.Port)
}
