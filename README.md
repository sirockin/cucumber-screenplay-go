## cucumber-screenplay-go

## Overview

Demonstration of the Screenplay Pattern using BDD with Cucumber/Gherkin in Go.

This monorepo contains three main components:
- **back-end**: Go service with domain logic and HTTP API
- **front-end**: React frontend (formerly web/)
- **acceptance**: BDD acceptance tests using Cucumber/Gherkin

The same BDD scenarios can be run against different deployment models:
- Direct domain access
- HTTP API with in-process server
- HTTP API with separate server executable
- HTTP API with Docker container
- Full UI testing with frontend and API in containers

Based on the official [Cucumber Screenplay Example](https://github.com/cucumber-school/screenplay-example/tree/code) to `go`, using the official [godog](https://github.com/cucumber/godog/) library.

## About Screenplay

The [Screenplay Pattern](https://serenity-js.org/handbook/design/screenplay-pattern/) is an evolution of the Page Object Pattern for UI test automation. It uses Actors with Abilities to perform Actions which can be grouped into Tasks. Actors can also ask Questions to verify outcomes.

While it has its origins in UI test automation, the pattern is applicable to acceptance testing in general, including API testing as demonstrated here.

It is useful where there are multiple user roles (Actors) with different capabilities (Abilities) performing different operations (Tasks) to achieve different goals (Questions).

It carries with it an extra layer of complexity and is probably best suited to projects in organizations where BDD is already well understood and widely adopted.

## Run Tests

### Using Makefile (Recommended)

```sh
# Show all available commands
make help

# Run fast tests (domain + in-process HTTP)
make test

# Run individual test types
make test-domain              # Domain unit tests (fastest)
make test-http-inprocess      # In-process HTTP integration tests
make test-http-executable     # Real server executable tests
make test-http-docker         # Docker container tests

# Run test suites
make test-fast                # Fast tests only
make test-integration         # All integration tests
make test-all                 # Full test suite including Docker
```

### Direct Go Commands

```sh
# Run all tests (from acceptance directory)
cd acceptance && go test -v .

# Run specific test types
cd acceptance && go test -v -run TestApplication .
cd acceptance && go test -v -run TestHTTPInProcess .
cd acceptance && go test -v -run TestHttpExecutable .
cd acceptance && go test -v -run TestHttpDocker .
cd acceptance && go test -v -run TestUI .
```


## Test Details

The code replicates that of the original javascript project and completes the use of Actor objects to implement each step. Like the original code it:
- Uses Actors with Abilities, and Actions which can be grouped to represent Tasks

Unlike the javascript project, it also uses Questions and associated helper methods, allowing all scenario steps to be delegated to Actor methods

There are some differences in structure:
- `godog` does not seem to support [cucumber expressions](https://github.com/cucumber/cucumber-expressions#readme) so:
   - regular expressions are used to map parameters as per godog examples
   - actors are created and accessed by an `Actor(name string)` method on the `suite` object
- `go` does not support arrow functions so the implementation of actions, tasks etc uses standard functions

- to promote separation of concerns:
   - the domain implementation code is placed in the `back-end/internal/domain` package following Go conventions
   - public domain interfaces are exposed via `back-end/pkg/domain/` for use by acceptance tests
   - the HTTP server implementation is in the `back-end/internal/http` package
   - test drivers in `acceptance/driver` provide different ways to access the domain (direct, HTTP client, UI automation)
   - We inject the application into the test suite via the go test functions so we no longer have an exported InitializeScenarios function. This means the tests can no longer be run from `godog run` but instead should be run from `go test`
   - acceptance test code has been placed in the `acceptance` folder with its own Go module and split into several files

## Architecture

The project follows clean architecture principles with a monorepo structure:

```
back-end/           # Go service (independent module)
├── cmd/server/     # Runnable HTTP server
├── internal/       # Internal implementation packages
│   ├── domain/     # Core business logic
│   └── http/       # HTTP server implementation
└── pkg/            # Public packages for external use
    ├── domain/     # Domain entities and services
    └── http/       # HTTP server interface

front-end/          # React frontend (formerly web/)
├── src/            # React source code
├── public/         # Static assets
└── Dockerfile      # Frontend container

acceptance/         # BDD tests (independent module)
├── driver/         # Test drivers for different deployment modes
│   ├── application/# Direct domain access driver
│   ├── http/       # HTTP client driver
│   └── ui/         # UI automation driver
├── screenplay/     # Screenplay pattern implementation
└── features/       # Gherkin feature files
```

## Test Levels

- **Application Tests** (`make test-domain`): Direct testing of business logic (fastest ~2-3ms)
- **HTTP In-Process** (`make test-http-inprocess`): HTTP API testing with in-process server (~4-5ms)
- **Server Executable** (`make test-http-executable`): Full integration with separate server process (~1-2s)
- **Docker Container** (`make test-http-docker`): Production-like containerized testing (~30-60s)
- **UI Tests** (`make test-ui`): Full stack testing with frontend and API containers using browser automation (~60-120s)

All tests run identical BDD scenarios ensuring contract compliance across all deployment models.

### Development Workflow

```sh
# Fast feedback during development
make test-fast

# Before committing changes
make test-integration

# Full validation (CI/CD)
make test-all
```

## Build and Run Server

```sh
# Build server binary
make build

# Build and run server
make server

# Or run directly from back-end directory
cd back-end && go run ./cmd/server

# Run frontend development server
cd front-end && npm start
```

## Module Structure

Each component is now its own Go module for better dependency management:

- `back-end/go.mod` - Backend service module
- `acceptance/go.mod` - Acceptance tests module that imports backend as dependency
- `front-end/package.json` - Frontend dependencies

The acceptance tests use the backend's public API via `back-end/pkg/` packages.
