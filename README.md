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
Se preferir rodar o PostgreSQL localmente, crie um banco de dados:
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

### **5️⃣ Rodar Migrações**
```sh
go run main.go migrate
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

## 🛠️ Estrutura do Projeto
```
/project-root
 ├── cmd/              # Entrada principal da aplicação
 ├── config/           # Configurações e variáveis de ambiente
 ├── models/           # Modelos do banco de dados
 ├── repositories/     # Camada de acesso aos dados
 ├── services/         # Lógica de negócio
 ├── handlers/         # Controllers e rotas
 ├── migrations/       # Scripts de migração do banco
 ├── .env.example      # Exemplo do arquivo .env
 ├── Dockerfile        # Configuração Docker
 ├── docker-compose.yml# Configuração Docker Compose
 ├── Makefile          # Comandos úteis
 ├── README.md         # Documentação do projeto
```

---

Sempre que algo for alterado no setup do projeto, este README deve ser atualizado! 🚀

