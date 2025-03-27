package storage

import (
	"fmt"
	"reflect"
	"time"

	"github.com/ttvdmt/url_shortener/internal/config"
)

type Storager interface {
	Save(url string, ttl time.Duration) (string, error)
	Get(code string) (string, error)
	CleanUp() error
	Close() error
}

func New(cfg config.Config) (Storager, error) {
	if reflect.DeepEqual(cfg, config.Config{}) {
		return nil, fmt.Errorf("cant create new storage: empty config")
	}

	switch {
	case cfg.Storage.SQLite.DBPath != "":
		dbpath := cfg.Storage.SQLite.DBPath

		db, err := newSQLite(dbpath)
		if err != nil {
			return nil, fmt.Errorf("cant create new storage: %w", err)
		}

		return db, nil

	case cfg.Storage.PostgreSQL.Host != "":
		user := cfg.Storage.PostgreSQL.User
		password := cfg.Storage.PostgreSQL.Password
		host := cfg.Storage.PostgreSQL.Host
		port := cfg.Storage.PostgreSQL.Port
		dbname := cfg.Storage.PostgreSQL.DBName
		sslmode := cfg.Storage.PostgreSQL.SSLMode

		db, err := newPostgres(user, password, host, port, dbname, sslmode)
		if err != nil {
			return nil, fmt.Errorf("cant create new storage: %w", err)
		}

		return db, nil

	default:
		return nil, fmt.Errorf("cant create new storage: wrong config")
	}
}
