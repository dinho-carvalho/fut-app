.PHONY: all build run test clean docker-build docker-run docker-stop lint test-coverage

# Variáveis
APP_NAME=fut-app
DOCKER_IMAGE=fut-app
COVERAGE_FILE=coverage.out
MIN_COVERAGE=98

# Comandos principais
all: clean build

build:
	@echo "🔨 Building..."
	@go build -o bin/$(APP_NAME) cmd/main.go

run:
	@echo "🚀 Running..."
	@go run cmd/main.go

test:
	@echo "🧪 Running tests..."
	@go test -v ./...

clean:
	@echo "🧹 Cleaning..."
	@rm -rf bin/
	@rm -f $(COVERAGE_FILE)
	@go clean -testcache

# Docker
docker-build:
	@echo "🐳 Building Docker image..."
	@docker build -t $(DOCKER_IMAGE) .

docker-run:
	@echo "🐳 Running Docker container..."
	@docker run -p 8080:8080 $(DOCKER_IMAGE)

docker-stop:
	@echo "🛑 Stopping Docker container..."
	@docker stop $(DOCKER_IMAGE) || true
	@docker rm $(DOCKER_IMAGE) || true

# Linting e Qualidade de Código
lint:
	@echo "🔍 Running linter..."
	@golangci-lint run

# Testes com Cobertura
test-coverage:
	@echo "📊 Running tests with coverage..."
	@go test -coverprofile=$(COVERAGE_FILE) ./...
	@go tool cover -func=$(COVERAGE_FILE) | grep total | awk '{print $$3}' | sed 's/%//' | awk '{if ($$1 < $(MIN_COVERAGE)) {print "❌ Coverage " $$1 "% is below minimum $(MIN_COVERAGE)%"; exit 1} else {print "✅ Coverage " $$1 "% meets minimum $(MIN_COVERAGE)%"}}'

coverage-html: test-coverage
	@echo "📈 Generating coverage report..."
	@go tool cover -html=$(COVERAGE_FILE)

# Ajuda
help:
	@echo "🔧 Available commands:"
	@echo "  make build          - Build the application"
	@echo "  make run           - Run the application"
	@echo "  make test          - Run tests"
	@echo "  make clean         - Clean build artifacts"
	@echo "  make docker-build  - Build Docker image"
	@echo "  make docker-run    - Run Docker container"
	@echo "  make docker-stop   - Stop Docker container"
	@echo "  make lint          - Run linter"
	@echo "  make test-coverage - Run tests with coverage"
	@echo "  make coverage-html - Generate coverage report"
