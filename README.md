# Bisnode Directory Search API

A Go service that provides an HTTP API for searching person and company information using Bisnode's Directory Search API. The service supports searching by mobile number for individuals and by organization number for companies.

## Prerequisites

- Go 1.22 or later
- Bisnode API credentials (username and password)

## Setup

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd bisnode-directory-search
   ```

2. Update the configuration:
   - Copy `config.json.example` to `config.json`
   - Update the `config.json` with your Bisnode API credentials

3. Install dependencies:
   ```bash
   go mod tidy
   ```

## Running the Application

```bash
go run cmd/api/main.go
```

The API will be available at `http://localhost:8080`

## API Endpoints

### Search for a Person by Mobile Number

```http
POST /api/v1/directory/persons/search
Content-Type: application/json

{
  "mobileNumber": "12345678"
}
```

### Search for an Organization by Number

Using query parameter:
```http
GET /api/v1/directory/organizations/search?orgNo=123456789
```

Or using JSON body:
```http
POST /api/v1/directory/organizations/search
Content-Type: application/json

{
  "organizationNumber": "123456789"
}
```

### Health Check

```http
GET /health
```

## Example Usage

### Using cURL

```bash
# Search person by mobile number
curl -X POST http://localhost:8080/api/v1/directory/persons/search \
  -H "Content-Type: application/json" \
  -d '{"mobileNumber":"12345678"}'

# Search organization by number (query param)
curl -X GET "http://localhost:8080/api/v1/directory/organizations/search?orgNo=123456789"

# Search organization by number (JSON body)
curl -X POST http://localhost:8080/api/v1/directory/organizations/search \
  -H "Content-Type: application/json" \
  -d '{"organizationNumber":"123456789"}'
```

### Using PowerShell

```powershell
# Search person by mobile number
irm -Uri "http://localhost:8080/api/v1/directory/persons/search" `
  -Method Post `
  -Body '{"mobileNumber":"12345678"}' `
  -ContentType "application/json"

# Search organization by number (query param)
irm -Uri "http://localhost:8080/api/v1/directory/organizations/search?orgNo=123456789" `
  -Method Get

# Search organization by number (JSON body)
irm -Uri "http://localhost:8080/api/v1/directory/organizations/search" `
  -Method Post `
  -Body '{"organizationNumber":"123456789"}' `
  -ContentType "application/json"
```

## Configuration

Create a `config.json` file in the root directory:

```json
{
  "bisnode": {
    "base_url": "https://apps.bisnode.no/api",
    "client_id": "your_username_here",
    "client_secret": "your_password_here"
  }
}
```

Or use environment variables:
- `BISNODE_BASE_URL`
- `BISNODE_CLIENT_ID`
- `BISNODE_CLIENT_SECRET`

## Building

```bash
# Build for current platform
go build -o bin/api cmd/api/main.go

# Cross-compile for Linux
GOOS=linux GOARCH=amd64 go build -o bin/api-linux-amd64 cmd/api/main.go
```

## License

MIT
