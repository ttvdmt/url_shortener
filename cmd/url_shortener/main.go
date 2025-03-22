package main

import (
	"github.com/ttvdmt/url_shortener/internal/handler"
	"github.com/ttvdmt/url_shortener/internal/storage"
)

func main() {
	st := storage.NewStorage()

	handler.Init(st)
	handler.Listen(":8080")
}
