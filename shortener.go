package main

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"sync"
)

type Shortener struct {
	mu   sync.RWMutex
	urls map[string]string
}

func NewShortener() *Shortener {
	return &Shortener{
		urls: make(map[string]string),
	}
}

func (s *Shortener) Shorten(originalURL string) string {
	s.mu.Lock()
	defer s.mu.Unlock()

	bytes := make([]byte, 6)
	rand.Read(bytes)
	shortCode := base64.URLEncoding.EncodeToString(bytes)

	s.urls[shortCode] = originalURL
	return shortCode
}

func (s *Shortener) Retrieve(shortCode string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	originalURL, exists := s.urls[shortCode]
	if !exists {
		var ErrURLNotFound = errors.New("URL not found")
		return "", ErrURLNotFound
	}

	return originalURL, nil
}
