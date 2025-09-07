# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Development Commands

### Core Commands
- `make run` - Run the application (starts server on port 8080)
- `make test` - Run all tests with `go test ./...`
- `make fmt` - Format code using gofumpt (`gofumpt -w .`)
- `make up` - Start Docker services with `docker-compose up -d`
- `make down` - Stop Docker services with `docker-compose down`
- `make clean` - Clean cache and dependencies (`rm -rf bin/* && go clean -modcache`)

### Linting and Code Quality
- `golangci-lint run` - Run linter (configured in `.golangci.yml`)
- Enabled linters: govet, errcheck, unused, typecheck
- Test files exclude errcheck linting

### Testing
- Run single test: `go test ./path/to/package -run TestName`
- Coverage: Coverage reports are generated in `coverage.out` and `gocov.html`

## Project Architecture

This is a Go web application for managing amateur football player statistics using Clean Architecture principles.

### Directory Structure
```
cmd/                    # Application entry points
├── main.go            # Main application entry with server setup
├── routes.go          # HTTP route definitions
└── app_dependencies.go # Dependency injection container

internal/
├── database/          # Database layer
│   ├── models/       # GORM models (Player, Position, Match, Rating)
│   ├── repositories/ # Data access layer
│   └── gateway/      # Gateway implementations
├── domain/           # Business entities and domain logic
├── usecase/         # Business use cases
├── handlers/        # HTTP handlers and DTOs
│   ├── dto/         # Data Transfer Objects
│   ├── middleware/  # HTTP middleware (validation, error handling)
│   └── httprespond/ # Response utilities
├── services/        # Application services
└── errors/          # Error definitions

pkg/
└── logger/          # Logging utilities
```

### Architecture Patterns
- **Clean Architecture**: Clear separation between layers (handlers → usecase → gateway → repository)
- **Dependency Injection**: Dependencies injected via `InjectDependencies()` function in `cmd/app_dependencies.go`
- **Repository Pattern**: Data access abstracted through repository interfaces
- **Gateway Pattern**: External dependencies abstracted through gateway interfaces

### Key Components
- **Database**: GORM with PostgreSQL, auto-migration on startup
- **HTTP Framework**: Gorilla Mux router
- **Validation**: JSON validation middleware using go-playground/validator
- **Logging**: Structured logging with slog
- **Environment**: Environment-specific config loading with godotenv

### Data Models
- `Player`: Main entity with name, positions (many-to-many), and JSONB stats
- `Position`: Football positions with unique names
- `Match`: Game records
- `Rating`: Player ratings (45-99 scale for amateur football)

### Current Implementation Status
- Player registration endpoint: `POST /players`
- Health check endpoint: `GET /health`
- Other CRUD operations are commented out in routes but partially implemented

### Database Configuration
- Uses PostgreSQL in production/docker
- SQLite for testing
- Environment variables: `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`
- Connection managed through `internal/database/config.go`

### Testing Approach
- Unit tests for repositories, use cases, DTOs, and middleware
- Test files follow `*_test.go` pattern
- Uses testify library for assertions

### Code Style
- Uses gofumpt for formatting (stricter than gofmt)
- Portuguese comments and messages in some areas
- Structured logging with contextual information