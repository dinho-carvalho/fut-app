.PHONY: run db-up db-down migrate fmt test clean coverage-html

# Roda a aplicação
run:
	go run ./cmd/main.go

# Sobe ...
up:
	docker-compose up -d

# Derruba....
down:
	docker-compose down

# Roda as migrações do banco
migrate:
	go run main.go migrate

# Formata o código usando gofumpt
fmt:
	gofumpt -w .

# Roda os testes
test:
	go test ./...

# Generates an HTML coverage report
coverage-html:
	@echo "Generating HTML coverage report..."
	@go test -coverprofile=coverage.out -covermode=atomic ./internal/handlers/... ./internal/repositories/... ./internal/services/...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "HTML coverage report generated: coverage.html"

# Limpa o cache e dependências antigas
clean:
	rm -rf bin/* && go clean -modcache
