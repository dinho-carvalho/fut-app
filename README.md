# 📌 Projeto Futebol Stats

Este projeto tem como objetivo gerenciar estatísticas de jogadores de futebol amador, permitindo a avaliação de desempenho com notas de 45 a 99.

---

## 🚀 Como Rodar o Projeto do Zero

### **1️⃣ Pré-requisitos**
Antes de começar, você precisará ter instalado:
- **Golang** (versão 1.20+)
- **Docker** e **Docker Compose**
- **PostgreSQL** (caso prefira rodar localmente sem Docker)
- **Make** (opcional, para rodar comandos mais facilmente)
- **gofumpt** (para manter o padrão de formatação do código)

### **2️⃣ Configurar Projeto**
Se estiver utilizando Docker, basta rodar:
```sh
make up
```
Ou, manualmente:
```sh
docker-compose up -d
```
Se preferir rodar o PostgreSQL localmente:
```sh
docker run --name futebol-db -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=admin -e POSTGRES_DB=futebol_stats -p 5432:5432 -d postgres
```

E configure a conexão no `.env`.

### **3️⃣ Configurar Variáveis de Ambiente**
Crie um arquivo **`.env`** na raiz do projeto e adicione:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=futebol_stats
```

### **4️⃣ Instalar Dependências**
```sh
go mod tidy
```

### **6️⃣ Rodar a Aplicação**
```sh
go ./cmd/run main.go
```
Ou, com Makefile:
```sh
make run
```

### **7️⃣ Testar API**
Acesse `http://localhost:8080` para verificar se a API está rodando.

---

## 📌 Comandos Úteis

### **🚀 Subir o app com docker**
```sh
make up
```
### **🛑 Para instância do docker**
```sh
make down
```
### **🧹 Limpar Dependências e Cache**
```sh
make clean
```
### **🔄 Rodar Tests**
```sh
make test
```
### **📝 Formatar Código com gofumpt**
```sh
gofumpt -w .
```

---

---

## CI/CD Pipeline

This project uses GitHub Actions for its CI/CD pipeline. The workflow is defined in `.github/workflows/go.yml` and includes the following stages:

1.  **Setup Go**: Initializes the Go environment.
2.  **Checkout Code**: Checks out the repository's code.
3.  **Download Dependencies**: Fetches the necessary Go modules.
4.  **Code Formatting**: Uses `gofumpt` to ensure consistent code style. The build will fail if any files need reformatting.
5.  **Linting**: A dedicated GitHub Actions workflow (`.github/workflows/golangci-lint.yml`) employs `golangci-lint` with the configurations defined in `.golangci.yml` to perform static analysis and identify potential issues.
6.  **Unit Testing & Coverage**: Runs all unit tests using `go test`. It also calculates code coverage and aims for a target of 98%. (Note: Test coverage enforcement will be active once tests are implemented).
7.  **Build**: Compiles the application using `go build`.

This pipeline helps maintain code quality and ensures that the application builds correctly with every push and pull request to the `main` branch.
