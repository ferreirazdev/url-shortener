package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Server struct {
	shortener *Shortener
}

func NewServer() *Server {
	return &Server{
		shortener: NewShortener(),
	}
}

func (s *Server) handleShorten(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	var req struct {
		URL string `json:"url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	type result struct {
		shortURL string
		err      error
	}

	resultChan := make(chan result, 1)
	go func() {
		select {
		case <-ctx.Done():
			resultChan <- result{"", ctx.Err()}
			return
		default:
			shortURL := s.shortener.Shorten(req.URL)
			resultChan <- result{shortURL, nil}
		}
	}()

	select {
	case <-ctx.Done():
		http.Error(w, "Request timeout", http.StatusRequestTimeout)
		return
	case res := <-resultChan:
		if res.err != nil {
			http.Error(w, res.err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{
			"short_url": res.shortURL,
		})
	}
}

func (s *Server) handleRetrieve(w http.ResponseWriter, r *http.Request) {
	shortCode := r.URL.Path[1:] // Remove leading /

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	resultChan := make(chan struct {
		url string
		err error
	}, 1)

	go func() {
		select {
		case <-ctx.Done():
			resultChan <- struct {
				url string
				err error
			}{"", ctx.Err()}
		default:
			url, err := s.shortener.Retrieve(shortCode)
			resultChan <- struct {
				url string
				err error
			}{url, err}
		}
	}()

	select {
	case <-ctx.Done():
		http.Error(w, "Request timeout", http.StatusRequestTimeout)
	case res := <-resultChan:
		if res.err != nil {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		http.Redirect(w, r, res.url, http.StatusFound)
	}
}

func main() {
	server := NewServer()

	http.HandleFunc("/shorten", server.handleShorten)
	http.HandleFunc("/", server.handleRetrieve)

	fmt.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
