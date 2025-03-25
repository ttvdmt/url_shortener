package storage

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/ttvdmt/url_shortener/pkg/encode"
)

type SQL_Storage struct {
	Mu *sync.RWMutex
	Db *sql.DB
}

func NewSQLStorage(path string) (*SQL_Storage, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("cant open database: %v", err)
	}

	db.Exec("PRAGMA journal_mode=WAL;")

	return &SQL_Storage{Db: db, Mu: &sync.RWMutex{}}, nil
}

func (s *SQL_Storage) Save(url string, ttl time.Duration) (string, error) {
	code := encode.GenerateCode(6)

	death_time := time.Now().Add(ttl).Format("2006-01-02")
	if _, err := s.Db.Exec("INSERT INTO urls (code, original_url, death_time) VALUES (?, ?, ?)", code, url, death_time); err != nil {
		return "", fmt.Errorf("cant save URL: %w", err)
	}

	return code, nil
}

func (s *SQL_Storage) Get(code string) (string, error) {
	var originalURL string
	err := s.Db.QueryRow("SELECT original_url FROM urls WHERE code = ?", code).Scan(&originalURL)

	if err != nil {
		return "", fmt.Errorf("cant get URL: %w", err)
	}

	return originalURL, nil
}

func (s *SQL_Storage) Close() error {
	return s.Db.Close()
}

func (s *SQL_Storage) CleanUp() error {
	_, err := s.Db.Exec("DELETE FROM urls WHERE death_time < date('now')")
	if err != nil {
		return fmt.Errorf("cant clean old urls: %w", err)
	}

	return nil
}
