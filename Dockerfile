# Estágio de build
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copia apenas os arquivos necessários primeiro
COPY go.mod go.sum ./
RUN go mod download

# Copia o resto do código
COPY . .

# Compila o binário com otimizações
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# Estágio final
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copia apenas o binário compilado
COPY --from=builder /app/main .

# Usuário não-root para segurança
RUN adduser -D appuser
USER appuser

# Define a porta da aplicação
EXPOSE 8080

# Healthcheck (opcional, mas recomendado)
HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 \
    CMD wget --quiet --tries=1 --spider http://localhost:8080/health || exit 1

CMD ["./main"]