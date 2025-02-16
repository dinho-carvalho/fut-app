.PHONY: run db-up db-down migrate fmt test clean

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

# Limpa o cache e dependências antigas
clean:
	rm -rf bin/* && go clean -modcache
