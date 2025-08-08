# Listagem de Deputados

Este projeto implementa um caso de uso para listar deputados (nome, partido, número de votos) ordenados por maior número de votos.

A aplicação é desenvolvida em **Go 1.22**, utilizando os seguintes serviços:

- REST API (GET `/deputado`)
- gRPC (`ListDeputados`)
- GraphQL (`listDeputados`)
- Banco de dados PostgreSQL (com migração automática via Docker)
- Containerização com Docker e Docker Compose

---

## 🧪 Tecnologias utilizadas

- Go 1.22
- PostgreSQL 15
- gRPC + Protobuf
- gqlgen (GraphQL)
- Echo (REST)
- Docker / Docker Compose
- Migrations via `init.sql`

---

## 📦 Como executar o projeto

### 1. Clone o repositório

```bash
  git clone https://github.com/agnaldopidev/deputados-app
  cd deputados-app
```

### 2. Suba os containers (PostgreSQL com migração automática)

```bash
  docker compose up --build
```
Isso irá:
- Criar o banco `deputadosdb`
- Criar e popular a tabela `deputados`
- Subir o PostgreSQL na porta `5432`
- Subir o container para rest `8080`
- Subir o container para grpc `50051`
- Subir o container para graphql `8081`

## Acessos

- REST: http://localhost:8080
- GraphQL: http://localhost:8081
- PostgreSQL: localhost:5432 (user: agnaldo, password: teste123)
- gRPC: localhost:50051

## Teste REST (via api.http ou curl)

```http
GET http://localhost:8080/deputados
```

## Variáveis de ambiente

Todos os serviços usam:

- DB_HOST=db
- DB_PORT=5432
- DB_USER=postres
- DB_PASSWORD=teste123
- DB_NAME=deputadosdb

---

## 🌐 Endpoints disponíveis

### REST

- **GET** `/deputado`  
  Retorna todos os deputados ordenados por votos (descendente)

```http
GET http://localhost:8080/deputado
```

### gRPC

- Serviço: `deputado.DeputadoService`
- Método: `ListDeputados`

Exemplo via grpcurl:

```bash
  grpcurl -plaintext localhost:50051 deputado.DeputadoService/ListDeputados
```

### GraphQL

- Endpoint: `POST /graphql`
- Playground: `http://localhost:8081/graphql`

Exemplo de query:

```graphql
query {
  listDeputados {
    id
    nome
    partido
    votos
  }
}
```

---

## 📂 Estrutura do projeto

```
.
├── api.http                      # Requisições REST
├── docker-compose.yaml           # Serviço PostgreSQL
├── Dockerfile                    # Container da aplicação Go (opcional)
├── go.mod / go.sum
├── cmd/
│    ├── rest/      # Serviço REST (porta 8080)
│    ├── grpc/      # Serviço gRPC (porta 50051)
│    └── graphql/   # Serviço GraphQL (porta 9090)
├── internal/
│   ├── handler/                  # REST Handlers
│   ├── grpc/                     # gRPC Server
│   │      └── proto/
│   │           └── deputado.proto # Protobuf gRPC
│   ├── graphql/                  # Resolvers + schema
│   ├── repository/               # DB access (PostgreSQL)
│   └── domain/                   # Regras de negócio
├── migrations/
│   └── init.sql                  # Criação e seed do banco
└── README.md
```

---

## ⚙️ Configurações

Banco de dados (em `docker-compose.yaml`):

- DB: `deputadosdb`
- Usuário: `postgres`
- Senha: `teste123`
- Porta: `5432`

---

## ✅ Status do Projeto

- [x] REST funcionando
- [x] gRPC funcionando
- [x] GraphQL funcionando
- [x] Banco com seed automático
- [x] Docker e Compose configurados
