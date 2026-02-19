package main

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
)

type Shortener struct {
	urls map[string]string
}

func NewShortener() *Shortener {
	return &Shortener{
		urls: make(map[string]string),
	}
}

func (s *Shortener) Shorten(originalURL string) string {
	bytes := make([]byte, 6)
	rand.Read(bytes)
	shortCode := base64.URLEncoding.EncodeToString(bytes)

	s.urls[shortCode] = originalURL
	return shortCode
}

func (s *Shortener) Retrieve(shortCode string) (string, error) {
	originalURL, exists := s.urls[shortCode]
	if !exists {
		var ErrURLNotFound = errors.New("URL not found")
		return "", ErrURLNotFound
	}

	return originalURL, nil
}
