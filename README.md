# Bisnode Person Search API

A Go service that provides an HTTP API for searching person information using Bisnode's API.

## Prerequisites

- Go 1.22 or later
- Bisnode API credentials (username and password)

## Setup

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd bisnode-person-search
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

### Search for a Person

```http
POST /api/v1/persons/search
Content-Type: application/json

{
  "mobileNumber": "12345678"
}
```

### Health Check

```http
GET /health
```

## Configuration

The application is configured using `config.json`:

```json
{
  "bisnode": {
    "base_url": "https://apps.bisnode.no/api",
    "client_id": "your_username_here",
    "client_secret": "your_password_here"
  }
}
```

## Environment Variables

For production, you can also use environment variables:

- `BISNODE_BASE_URL`
- `BISNODE_CLIENT_ID`
- `BISNODE_CLIENT_SECRET`

## Building for Production

```bash
go build -o bin/api cmd/api/main.go
```

## License

MIT
