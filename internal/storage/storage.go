package storage

import (
	"sync"

	"github.com/ttvdmt/url_shortener/pkg/encode"
)

type Storage struct {
	Mu   *sync.RWMutex
	Urls map[string]string
}

func NewStorage() *Storage {
	return &Storage{
		Mu:   &sync.RWMutex{},
		Urls: make(map[string]string),
	}
}

func (s *Storage) Save(url string) string {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	code := encode.GenerateCode(6)
	s.Urls[code] = url
	return code
}

func (s *Storage) Get(code string) (string, bool) {
	s.Mu.RLock()
	defer s.Mu.RUnlock()

	url, exists := s.Urls[code]
	return url, exists
}
