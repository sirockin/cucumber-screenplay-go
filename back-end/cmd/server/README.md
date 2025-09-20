# HTTP API Server

A runnable HTTP server that implements the OpenAPI specification by wrapping the domain logic.

## Usage

### Build and Run
```bash
# Build the server
go build -o server ./cmd/server

# Run on default port (8080)
./server

# Run on custom port
./server -port=3000
```

### Direct Run
```bash
# Run without building binary
go run ./cmd/server

# Run on custom port
go run ./cmd/server -port=3000
```

## API Endpoints

The server implements all endpoints from the OpenAPI specification:

- `POST /accounts` - Create a new account
- `GET /accounts/{name}` - Get account details
- `POST /accounts/{name}/activate` - Activate an account
- `POST /accounts/{name}/authenticate` - Authenticate an account
- `GET /accounts/{name}/authentication-status` - Check authentication status
- `GET /accounts/{name}/projects` - Get user projects
- `POST /accounts/{name}/projects` - Create a project
- `DELETE /clear` - Clear all data (for testing)

## Example Usage

```bash
# Start the server
go run ./cmd/server

# Create an account
curl -X POST http://localhost:8080/accounts \
  -H "Content-Type: application/json" \
  -d '{"name": "alice"}'

# Activate the account
curl -X POST http://localhost:8080/accounts/alice/activate

# Check authentication status
curl http://localhost:8080/accounts/alice/authentication-status

# Create a project
curl -X POST http://localhost:8080/accounts/alice/projects

# Get projects
curl http://localhost:8080/accounts/alice/projects
```

## Architecture

The server uses the domain directly for clean architecture:

```
HTTP Request → HTTP Server → Domain Logic
```

The HTTP server (`internal/http`) wraps the domain (`internal/domain`) directly, ensuring the same business logic is used across all access patterns (direct domain access, HTTP API, etc.).