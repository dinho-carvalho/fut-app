# 📊 Fut-App - Sistema de Gestão de Jogadores de Futebol Amador

Uma aplicação web em Go para gerenciamento de estatísticas de jogadores de futebol amador, desenvolvida seguindo os princípios de Clean Architecture.

## 🏗️ Arquitetura

O projeto utiliza **Clean Architecture** com as seguintes camadas:
- **Domain**: Entidades e regras de negócio
- **Use Cases**: Casos de uso da aplicação
- **Handlers**: Camada de apresentação (HTTP)
- **Database**: Acesso a dados com GORM

### 📂 Estrutura do Projeto
```
cmd/                    # Ponto de entrada da aplicação
├── main.go            # Setup do servidor
├── routes.go          # Definição das rotas HTTP
└── app_dependencies.go # Container de injeção de dependência

internal/
├── database/          # Camada de dados
│   ├── models/       # Modelos GORM (Player, Position, Match, Rating)
│   └── repositories/ # Repositórios de acesso a dados
├── domain/           # Entidades de negócio
├── usecase/         # Casos de uso
├── handlers/        # Handlers HTTP e DTOs
└── services/        # Serviços da aplicação
```

## 🚀 Como Rodar

### Pré-requisitos
- **Go** 1.20+
- **Docker** e **Docker Compose**
- **Make** (opcional)

### 1. Clonar e configurar
```bash
git clone <repository-url>
cd fut-app
```

### 2. Subir banco de dados
```bash
make up
# ou
docker-compose up -d
```

### 3. Configurar ambiente
Crie um arquivo `.env`:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=futebol_stats
```

### 4. Instalar dependências
```bash
go mod tidy
```

### 5. Rodar aplicação
```bash
make run
# ou
go run cmd/main.go
```

A API estará disponível em `http://localhost:8080`

## 🛠️ Comandos Disponíveis

| Comando | Descrição |
|---------|-----------|
| `make run` | Executa a aplicação |
| `make test` | Roda todos os testes |
| `make fmt` | Formata código com gofumpt |
| `make up` | Inicia serviços Docker |
| `make down` | Para serviços Docker |
| `make clean` | Limpa cache e dependências |
| `golangci-lint run` | Executa linter |

## 📊 Modelo de Dados

- **Player**: Jogador com nome, posições e estatísticas
- **Position**: Posições de futebol
- **Match**: Registros de partidas
- **Rating**: Avaliações de jogadores (escala 45-99)

## 🔗 Endpoints Disponíveis

- `POST /players` - Cadastro de jogador
- `GET /health` - Verificação de saúde da API

## 🧪 Testes

```bash
# Todos os testes
make test

# Teste específico
go test ./internal/usecase -run TestRegisterPlayer

# Com coverage
go test ./... -coverprofile=coverage.out
```

## 📋 Qualidade de Código

- **Linting**: golangci-lint com govet, errcheck, unused, typecheck
- **Formatação**: gofumpt (mais rigoroso que gofmt)
- **Validação**: Middleware de validação JSON
- **Logging**: Logs estruturados com slog

## 🐳 Docker

O projeto inclui configuração Docker Compose com PostgreSQL. Para desenvolvimento local, também suporta SQLite para testes.

---

![Go CI/CD](https://github.com/dinho-carvalho/fut-app/workflows/Go%20CI%20CD/badge.svg)
[![codecov](https://codecov.io/gh/dinho-carvalho/fut-app/branch/main/graph/badge.svg)](https://codecov.io/gh/dinho-carvalho/fut-app)
