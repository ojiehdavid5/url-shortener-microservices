# URL Shortener Microservices

A lightweight Go microservices application that provides URL shortening and redirection functionality. Built with the high-performance Fiber web framework.

## Architecture

The application consists of two independent microservices:

### 1. URL Service (Port 8081)
Handles URL creation and retrieval. Stores shortened URLs in memory with randomly generated 6-character codes.

### 2. Redirect Service (Port 8082)
Handles redirects by querying the URL Service and performing HTTP 302 redirects to the original URLs.

## Prerequisites

- Go 1.16 or higher
- Terminal/shell access

## Installation

1. Clone the repository:
```bash
cd /Users/user/go/src/url-shortener-microservices
```

2. Install dependencies for each service:
```bash
# URL Service
cd url-service
go mod download

# Redirect Service
cd ../redirect-service
go mod download
```

## Running the Services

Start both services in separate terminal windows:

**Terminal 1 - URL Service:**
```bash
cd url-service
go run main.go
```
The service will start on `http://localhost:8081`

**Terminal 2 - Redirect Service:**
```bash
cd redirect-service
go run main.go
```
The service will start on `http://localhost:8082`

## API Endpoints

### URL Service (Port 8081)

#### Create a shortened URL
```
POST /urls
Content-Type: application/json

{
  "original_url": "https://www.example.com/very/long/url"
}
```

**Response (200 OK):**
```json
{
  "code": "aBc123",
  "original_url": "https://www.example.com/very/long/url"
}
```

#### Retrieve original URL
```
GET /urls/:code
```

**Response (200 OK):**
```json
{
  "original_url": "https://www.example.com/very/long/url"
}
```

**Response (404 Not Found):**
```
Not found
```

### Redirect Service (Port 8082)

#### Redirect to original URL
```
GET /:code
```

Redirects (HTTP 302) to the original URL associated with the code.

**Response (404 Not Found):**
```
Invalid URL
```

## Example Usage

### Create a shortened URL:
```bash
curl -X POST http://localhost:8081/urls \
  -H "Content-Type: application/json" \
  -d '{"original_url": "https://www.github.com"}'
```

Response:
```json
{
  "code": "aB3xYz",
  "original_url": "https://www.github.com"
}
```

### Use the shortened URL:
```bash
curl -L http://localhost:8082/aB3xYz
```

This will redirect to `https://www.github.com`

## Project Structure

```
url-shortener-microservices/
├── go.mod                    # Root module file
├── README.md
├── url-service/
│   ├── go.mod              # URL Service dependencies
│   └── main.go             # URL Service implementation
└── redirect-service/
    ├── go.mod              # Redirect Service dependencies
    └── main.go             # Redirect Service implementation
```

## Dependencies

Both services use:
- **Fiber** (github.com/gofiber/fiber/v2) - Fast HTTP web framework

## Notes

- URLs are stored in-memory and will be lost when the service restarts
- The 6-character codes are randomly generated with a mix of uppercase, lowercase, and numeric characters
- For production use, consider adding:
  - Persistent database storage (PostgreSQL, MongoDB, etc.)
  - Input validation and URL sanitization
  - Rate limiting
  - Error handling improvements
  - Logging and monitoring
  - Docker containerization

## License

[Add your license here]
