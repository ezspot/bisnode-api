# Bisnode API Client

A Go client for interacting with the Bisnode API with comprehensive Swagger/OpenAPI documentation.

## Features

- **Person Search**: Look up individuals by mobile number
- **Organization Search**: Find organizations by organization number
- **Vehicle Search**: Search for motor vehicles by license number or VIN
- **Interactive Documentation**: Full API documentation with Swagger UI
- **RESTful Endpoints**: Consistent API design following REST principles

## Table of Contents

- [Prerequisites](#prerequisites)
- [Quick Start](#quick-start)
- [API Documentation](#api-documentation)
- [API Reference](#api-reference)
  - [Authentication](#authentication)
  - [Persons](#search-for-a-person)
  - [Organizations](#search-for-an-organization)
  - [Vehicles](#search-for-a-vehicle)
  - [Health Check](#health-check)

## Prerequisites

- Go 1.22 or later
- Bisnode API credentials (username and password)

## Quick Start

1. **Clone and configure**:
   ```bash
   git clone <repository-url>
   cd bisnode-directory-search
   cp config.json.example config.json
   # Edit config.json with your credentials
   ```

2. **Install dependencies**:
   ```bash
   go mod tidy
   ```

3. **Run the server**:
   ```bash
   go run cmd/api/main.go
   ```

The API will be available at `http://localhost:8080`

## API Documentation

Access the interactive Swagger UI for complete API documentation:

```
http://localhost:8080/swagger/index.html
```

## API Reference

### Authentication

All endpoints require Basic Authentication. Include the `Authorization` header in your requests:

```http
Authorization: Basic <base64-encoded-username:password>
```

### Search for a Person

#### By Mobile Number (GET)
```http
GET /api/v1/directory/persons/search?mobileNumber=12345678
```

#### By Mobile Number (POST)
```http
POST /api/v1/directory/persons/search
Content-Type: application/json

{
  "mobileNumber": "12345678"
}
```

### Search for an Organization

#### By Organization Number (GET)
```http
GET /api/v1/directory/organizations/search?organizationNumber=123456789
```

#### By Organization Number (POST)
```http
POST /api/v1/directory/organizations/search
Content-Type: application/json

{
  "organizationNumber": "123456789"
}
```

### Search for a Motor Vehicle

Search by license number (GET):
```http
GET http://localhost:8080/api/v1/motor-vehicles/search?licenseNumber=AB12345
```

Search by VIN (GET):
```http
GET http://localhost:8080/api/v1/motor-vehicles/search?vin=WBAKG7C5XBE123456
```

Search using POST with JSON body:
```http
POST http://localhost:8080/api/v1/motor-vehicles/search
Content-Type: application/json

{
  "licenseNumber": "AB12345"
  // or "vin": "WBAKG7C5XBE123456"
}
```

### Health Check

```http
GET /health
```

Returns `200 OK` when the service is running.

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

# Search for a motor vehicle by license number
irm -Uri "http://localhost:8080/api/v1/motor-vehicles/search?licenseNumber=AB12345" `
  -Method Get

# Search for a motor vehicle by VIN
irm -Uri "http://localhost:8080/api/v1/motor-vehicles/search?vin=WBAKG7C5XBE123456" `
  -Method Get

# Search for a motor vehicle using POST
$body = @{ licenseNumber = "AB12345" } | ConvertTo-Json
# or $body = @{ vin = "WBAKG7C5XBE123456" } | ConvertTo-Json
irm -Uri "http://localhost:8080/api/v1/motor-vehicles/search" `
  -Method Post `
  -Body $body `
  -ContentType "application/json"

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

Create a `config.json` file in the root directory based on the `config.example.json` template:

```json
{
  "bisnode": {
    "base_url": "https://api.bisnode.no",
    "client_id": "your_username_here",
    "client_secret": "your_password_here"
  }
}
```

Or use environment variables:
- `BISNODE_BASE_URL` (default: `https://api.bisnode.no`)
- `BISNODE_CLIENT_ID`
- `BISNODE_CLIENT_SECRET`

Note: Ensure your account has access to the specific Bisnode API services you intend to use.

### Additional Features

Other Bisnode API endpoints can be implemented on request. Please contact support for more information.

## API Documentation

This project includes Swagger/OpenAPI documentation that's automatically generated from the code. To access the interactive API documentation:

1. Start the server:
   ```bash
   go run cmd/api/main.go
   ```

2. Open your browser and navigate to:
   ```
   http://localhost:8080/swagger/index.html
   ```

The Swagger UI provides:
- Interactive API documentation
- Try-it-out functionality for all endpoints
- Request/response schemas
- Authentication requirements

## Building

```bash
# Install Swag CLI (if not already installed)
go install github.com/swaggo/swag/cmd/swag@latest

# Generate Swagger documentation
swag init -g cmd/api/main.go

# Build for current platform
go build -o bin/api cmd/api/main.go

# Cross-compile for Linux
GOOS=linux GOARCH=amd64 go build -o bin/api-linux-amd64 cmd/api/main.go
```

After making changes to the API, remember to regenerate the Swagger documentation:
```bash
swag init -g cmd/api/main.go
```

## License

MIT
