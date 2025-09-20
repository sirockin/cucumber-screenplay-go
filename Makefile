.PHONY: test test-domain test-http-inprocess test-http-executable test-http-docker test-ui test-fast test-integration test-all clean build server help

# Default target
help: ## Show this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Individual test targets
test-domain: ## Run application unit tests (fastest)
	cd acceptance && go test -v -run TestApplication .

test-http-inprocess: ## Run in-process HTTP integration tests
	cd acceptance && go test -v -run TestHTTPInProcess .

test-http-executable: ## Run real server executable tests
	cd acceptance && go test -v -run TestHttpExecutable .

test-http-docker: ## Run Docker container tests (slowest)
	cd acceptance && go test -v -run TestHttpDocker .

test-ui: ## Run UI tests with frontend and API containers (requires Docker)
	cd acceptance && go test -v -run TestUI .

# Test suites
test-fast: ## Run fast tests (application + in-process HTTP)
	cd acceptance && go test -v -run "TestApplication|TestHTTPInProcess" .

test-integration: ## Run all integration tests (excluding Docker and UI)
	cd acceptance && go test -v -run "TestHTTPInProcess|TestHttpExecutable" .

test-all: ## Run all tests including Docker and UI (full suite)
	cd acceptance && go test -v .

test: test-fast ## Default test target (fast tests only)

# Test with short mode (unit tests only)
test-short: ## Run tests in short mode (skips slow integration tests)
	cd acceptance && go test -short -v .

# Build targets
build: ## Build the server binary
	cd back-end && go build -o bin/server ./cmd/server

server: build ## Build and run the server
	./back-end/bin/server

# Clean up
clean: ## Clean build artifacts
	rm -rf back-end/bin/

# Development helpers
fmt: ## Format Go code
	cd back-end && go fmt ./...
	cd acceptance && go fmt ./...

vet: ## Run go vet
	cd back-end && go vet ./...
	cd acceptance && go vet ./...

sec: ## Run security checks with gosec
	cd back-end && gosec ./...
	cd acceptance && gosec ./...

lint: fmt vet sec ## Run formatting and vetting

# Coverage
coverage: ## Run tests with coverage
	cd acceptance && go test -coverprofile=coverage.out .
	cd acceptance && go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: acceptance/coverage.html"