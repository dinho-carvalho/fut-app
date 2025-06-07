# Fut App - Estatísticas de Futebol Amador

## 📝 Sobre o Projeto

O Fut App é uma aplicação para gerenciar estatísticas de jogadores de futebol amador. Com ele, você pode:

- Registrar jogadores
- Registrar partidas
- Registrar resultados
- Acompanhar estatísticas de jogadores
- Avaliar desempenho dos jogadores

## 🚀 Tecnologias

- Go 1.21+
- PostgreSQL
- Docker
- Docker Compose
- GORM (ORM)
- Gorilla Mux (Router)

## 📋 Pré-requisitos

- Go 1.21 ou superior
- Docker e Docker Compose (opcional)
- PostgreSQL (se não usar Docker)
- Make (para comandos de conveniência)

## 🔧 Configuração Local

### Com Docker

1. Clone o repositório:
```bash
git clone https://github.com/seu-usuario/fut-app.git
cd fut-app
```

2. Inicie os containers:
```bash
docker-compose up -d
```

3. Execute a aplicação:
```bash
make run
```

### Sem Docker (PostgreSQL Local)

1. Clone o repositório:
```bash
git clone https://github.com/seu-usuario/fut-app.git
cd fut-app
```

2. Configure o PostgreSQL local:
```bash
createdb futebol_stats
```

3. Configure as variáveis de ambiente:
```bash
export DB_HOST=localhost
export DB_USER=seu_usuario
export DB_PASSWORD=sua_senha
export DB_NAME=futebol_stats
export DB_PORT=5432
```

4. Execute a aplicação:
```bash
make run
```

## 🛠️ Comandos Úteis

- `make run`: Executa a aplicação
- `make test`: Executa os testes
- `make test-coverage`: Executa os testes com cobertura
- `make lint`: Executa o linter
- `make build`: Compila o projeto
- `make docker-build`: Constrói a imagem Docker
- `make docker-run`: Executa o container Docker

## 📊 Cobertura de Testes

O projeto tem como meta manter uma cobertura de testes de 98%. Para verificar a cobertura:

```bash
make test-coverage
```

### Arquivos Ignorados na Cobertura

Alguns arquivos são ignorados na cobertura de testes por serem arquivos de configuração ou não necessitarem de testes:

- `cmd/main.go`
- `internal/database/config.go`
- Arquivos de migração

## 🌳 Estrutura do Projeto

```
.
├── cmd/
│   └── main.go
├── internal/
│   ├── database/
│   │   ├── config.go
│   │   ├── connection.go
│   │   └── models/
│   ├── handlers/
│   ├── repositories/
│   ├── routes/
│   └── services/
├── scripts/
├── .gitignore
├── .golangci.yml
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## 📝 Convenções de Código

- Utilizamos o `golangci-lint` para garantir a qualidade do código
- Seguimos as convenções do Go para nomes de pacotes e funções
- Documentamos todas as funções públicas
- Mantemos os testes atualizados com as mudanças no código

## 🤝 Contribuindo

1. Faça o fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## 📞 Contato

Seu Nome - [@seu_twitter](https://twitter.com/seu_twitter) - seu_email@email.com

Link do Projeto: [https://github.com/seu-usuario/fut-app](https://github.com/seu-usuario/fut-app)
