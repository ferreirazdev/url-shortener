package main

import (
	"fmt"
	"sync"
	"testing"
)

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

func TestRetrieveURL(t *testing.T) {
	shortener := NewShortener()

	originalURL := "https://example.com/test"
	shortURL := shortener.Shorten(originalURL)

	retrieved, err := shortener.Retrieve(shortURL)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if retrieved != originalURL {
		t.Errorf("Expected %s, got %s", originalURL, retrieved)
	}
}

func TestRetrieveNonExistentURL(t *testing.T) {
	shortener := NewShortener()

	_, err := shortener.Retrieve("nonexistent")
	if err == nil {
		t.Error("Expected error for non-existent URL")
	}
}

func TestConcurrentShorten(t *testing.T) {
	shortener := NewShortener()

	const numGoroutines = 100
	urls := make(map[string]string)

	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			url := fmt.Sprintf("https://example.com/%d", i)
			short := shortener.Shorten(url)

			mu.Lock()
			urls[short] = url
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	if len(urls) != numGoroutines {
		t.Errorf("Expected %d URLs, got %d", numGoroutines, len(urls))
	}
}
