package storage

import (
	"time"
)

type Storager interface {
	Save(url string, ttl time.Duration) (string, error)
	Get(code string) (string, error)
	CleanUp() error
}
