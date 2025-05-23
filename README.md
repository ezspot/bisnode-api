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

### Using PowerShell (Recommended)

```powershell
# Search for a person by mobile number (POST with JSON body)
$body = @{ mobileNumber = "12345678" } | ConvertTo-Json
irm -Uri "http://localhost:8080/api/v1/directory/persons/search" `
  -Method Post `
  -Body $body `
  -ContentType "application/json"

# Search for an organization by number (GET with query parameter)
irm -Uri "http://localhost:8080/api/v1/directory/organizations/search?orgNo=123456789" `
  -Method Get

# Health check
irm -Uri "http://localhost:8080/health" -Method Get
```

### Using cURL (if installed)

```bash
# Search for a person by mobile number
$body='{"mobileNumber":"12345678"}'
curl -X POST http://localhost:8080/api/v1/directory/persons/search \
  -H "Content-Type: application/json" \
  -d $body

# Search for an organization by number
curl -X GET "http://localhost:8080/api/v1/directory/organizations/search?orgNo=123456789"

# Health check
curl -X GET http://localhost:8080/health
```

## Response Format

### Person Search Response
```json
{
  "type": "Person",
  "organizationnumber": "931698400",
  "lastname": "Doe",
  "firstname": "John",
  "streetname": "Example Street",
  "houseno": "123",
  "zipcode": "1234",
  "city": "Oslo"
}
```

### Organization Search Response
```json
{
  "type": "Company",
  "organizationnumber": "123456789",
  "lastname": "Example Company AS",
  "streetname": "Business Street",
  "houseno": "456",
  "zipcode": "1234",
  "city": "Oslo"
}
```

### Error Response
```json
{
  "error": "Not Found",
  "message": "No results found"
}
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
