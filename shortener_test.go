package main

import "testing"

func TestShortenURL(t *testing.T) {
	shortener := NewShortener()

	originalURL := "https://www.google.com"
	shortURL := shortener.Shorten(originalURL)

	if shortURL == "" {
		t.Error("Expected a short URL, got empty string")
	}

	if shortURL == originalURL {
		t.Error("Short URL should be different from original")
	}
}
