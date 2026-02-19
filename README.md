# URL Shortener

A simple, thread-safe URL shortener service written in Go. This service allows you to shorten long URLs and retrieve the original URLs using short codes.

## Features

- **URL Shortening**: Convert long URLs into short, base64-encoded codes
- **URL Retrieval**: Retrieve original URLs using short codes with automatic redirects
- **Thread-Safe**: Concurrent-safe operations using mutex locks
- **Timeout Handling**: Request timeouts for both shortening and retrieval operations
- **In-Memory Storage**: Fast, in-memory storage using Go maps

## Requirements

- Go 1.25.1 or later

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd url-shortener
```

2. Install dependencies (if any):
```bash
go mod download
```

## Usage

### Running the Server

Start the server:
```bash
go run main.go shortener.go
```

The server will start on port `8080`. You should see:
```
Server starting on :8080
```

### API Endpoints

#### Shorten a URL

**POST** `/shorten`

Request body:
```json
{
  "url": "https://www.example.com/very/long/url/path"
}
```

Response:
```json
{
  "short_url": "dGVzdA=="
}
```

Example using curl:
```bash
curl -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{"url": "https://www.google.com"}'
```

#### Retrieve Original URL

**GET** `/{shortCode}`

The server will automatically redirect to the original URL.

Example:
```bash
curl -L http://localhost:8080/dGVzdA==
```

Or simply visit `http://localhost:8080/{shortCode}` in your browser.

## Testing

Run the test suite:
```bash
go test
```

Run tests with verbose output:
```bash
go test -v
```

The test suite includes:
- URL shortening functionality
- URL retrieval functionality
- Error handling for non-existent URLs
- Concurrent operations testing

## Project Structure

```
url-shortener/
├── main.go           # HTTP server and request handlers
├── shortener.go      # Core shortening logic and storage
├── shortener_test.go # Test suite
├── go.mod           # Go module definition
└── README.md        # This file
```

## Implementation Details

- **Storage**: URLs are stored in an in-memory map. Data is lost when the server restarts.
- **Short Codes**: Generated using cryptographically secure random bytes (6 bytes) encoded in base64 URL-safe format.
- **Concurrency**: All operations are protected by read-write mutexes to ensure thread safety.
- **Timeouts**: 
  - Shortening requests have a 5-second timeout
  - Retrieval requests have a 2-second timeout

## Limitations

- URLs are stored in memory only - data is lost on server restart
- No persistence layer (database)
- No URL validation
- No expiration mechanism for short URLs
- No analytics or tracking

## Future Enhancements

Potential improvements:
- Add database persistence
- URL validation
- Custom short code support
- URL expiration
- Analytics and click tracking
- Rate limiting
- Authentication/authorization
