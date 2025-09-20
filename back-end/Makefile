.PHONY: build server clean fmt vet help

# Default target
help: ## Show this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Build targets
build: ## Build the server binary
	go build -o bin/server ./cmd/server

server: build ## Build and run the server
	./bin/server

# Clean up
clean: ## Clean build artifacts
	rm -rf bin/

# Development helpers
fmt: ## Format Go code
	go fmt ./...

vet: ## Run go vet
	go vet ./...

lint: fmt vet ## Run formatting and vetting