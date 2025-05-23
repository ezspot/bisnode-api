# Bisnode API Client

A Go client for interacting with the Bisnode API.

## Features

- Search for persons by mobile number
- Search for organizations by organization number
- Search for motor vehicles by license number or VIN
- Additional features available on request for companies.
- Swagger documentation

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

## API Documentation

The API is documented using Swagger/OpenAPI. You can access the interactive documentation at:

```
http://localhost:8080/swagger/index.html
```

## API Endpoints

All endpoints are prefixed with `/api/v1`.

### Authentication

All endpoints require Basic Authentication. Include your credentials in the `Authorization` header:

```
Authorization: Basic <base64-encoded-username:password>
```

### Search for a Person by Mobile Number

#### GET Request
```http
GET http://localhost:8080/api/v1/directory/persons/search?mobileNumber=12345678
```

#### POST Request
```http
POST http://localhost:8080/api/v1/directory/persons/search
Content-Type: application/json

{
  "mobileNumber": "12345678"
}
```

### Search for an Organization by Number

#### GET Request
```http
GET http://localhost:8080/api/v1/directory/organizations/search?organizationNumber=123456789
```

#### POST Request
```http
POST http://localhost:8080/api/v1/directory/organizations/search
Content-Type: application/json

{
  "organizationNumber": "123456789"
}
```

### Search for a Motor Vehicle

#### GET Request (License Number)
```http
GET http://localhost:8080/api/v1/motor-vehicles/search?licenseNumber=AB12345
```

#### GET Request (VIN)
```http
GET http://localhost:8080/api/v1/motor-vehicles/search?vin=WBAKG7C5XBE123456
```

#### POST Request (JSON Body)
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
GET http://localhost:8080/health
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

## Usage Examples

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
