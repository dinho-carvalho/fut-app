# Estágio de build
FROM golang:1.24-alpine AS builder

# Instalar ferramentas necessárias
RUN apk add --no-cache postgresql-client curl gcc musl-dev

WORKDIR /app

# Copia apenas os arquivos necessários primeiro
COPY go.mod go.sum ./
RUN go mod download

# Instalar ferramentas de desenvolvimento
RUN go install mvdan.cc/gofumpt@latest
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.56.2

# Copia o resto do código
COPY . .

# Compila o binário com otimizações
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# Estágio final - usando a imagem golang em vez de alpine
FROM golang:1.24-alpine

# Instalar ferramentas necessárias
RUN apk --no-cache add ca-certificates postgresql-client wget bc bash

WORKDIR /app

# Copia o binário compilado do estágio de build
COPY --from=builder /app/main .
COPY --from=builder /go/bin/gofumpt /usr/local/bin/gofumpt
COPY --from=builder /go/bin/golangci-lint /usr/local/bin/golangci-lint

# Copiar o código fonte para poder executar testes
COPY . .

# Expõe a porta da aplicação
EXPOSE 8080

# Healthcheck
HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 \
    CMD wget --quiet --tries=1 --spider http://localhost:8080/health || exit 1

CMD ["./main"]