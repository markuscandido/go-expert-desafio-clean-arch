# Desafio Clean Architecture - Go Expert

Sistema de gerenciamento de pedidos seguindo os princípios da Clean Architecture, expondo três APIs: REST, gRPC e GraphQL.

## 🚀 Passo a Passo para Execução

### 1. Clone o repositório

```bash
git clone https://github.com/markuscandido/go-expert-desafio-clean-arch.git
cd go-expert-desafio-clean-arch
```

### 2. Inicie os serviços com Docker

```bash
docker-compose up -d
```

Este comando inicia:
- **MySQL**: Banco de dados com migrations automáticas
- **RabbitMQ**: Servidor de mensageria
- **App**: Aplicação Go

### 3. Verifique se os serviços estão rodando

```bash
docker-compose logs app
```

### 4. Acesse as APIs

| Serviço | Porta | Endpoint |
|---------|-------|----------|
| REST | 8000 | `http://localhost:8000/order` |
| gRPC | 50051 | `localhost:50051` |
| GraphQL | 8080 | `http://localhost:8080` (Playground) |

---

## 📡 APIs Disponíveis

### REST API (Porta 8000)

#### Criar Order
```bash
curl -X POST http://localhost:8000/order \
  -H "Content-Type: application/json" \
  -d '{"id": "order-001", "price": 100.0, "tax": 10.0}'
```

#### Listar Orders
```bash
curl http://localhost:8000/order
```

---

### GraphQL (Porta 8080)

Acesse o playground em: http://localhost:8080

#### Criar Order
```graphql
mutation {
  createOrder(input: { id: "order-001", Price: 100.0, Tax: 10.0 }) {
    id
    Price
    Tax
    FinalPrice
  }
}
```

#### Listar Orders
```graphql
query {
  listOrders {
    id
    Price
    Tax
    FinalPrice
  }
}
```

---

### gRPC (Porta 50051)

Instale o Evans para testar:
```bash
go install github.com/ktr0731/evans@latest
```

Conecte ao servidor:
```bash
evans -r repl -p 50051
```

> **Importante:** Antes de chamar os métodos, você precisa selecionar o package e o serviço.

#### Configuração inicial no Evans REPL
```
127.0.0.1:50051> package pb
127.0.0.1:50051> service OrderService
```

#### Listar Orders
```
pb.OrderService@127.0.0.1:50051> call ListOrders
{}
```

#### Criar Order
```
pb.OrderService@127.0.0.1:50051> call CreateOrder
id (TYPE_STRING) => order-001
price (TYPE_FLOAT) => 100.0
tax (TYPE_FLOAT) => 10.0
```

#### Comandos úteis do Evans
```
show package    # Lista pacotes disponíveis
show service    # Lista serviços do pacote selecionado
show message    # Lista mensagens/tipos disponíveis
```

---

## 🛠️ Desenvolvimento Local

### Pré-requisitos

- Go 1.24+
- Docker e Docker Compose
- protoc (Protocol Buffers compiler)

### Executar sem Docker

1. Inicie MySQL e RabbitMQ via Docker:
```bash
docker-compose up -d mysql rabbitmq
```

2. Crie o arquivo `.env`:
```env
DB_DRIVER=mysql
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=root
DB_NAME=orders
WEB_SERVER_PORT=8000
GRPC_SERVER_PORT=50051
GRAPHQL_SERVER_PORT=8080
RABBITMQ_HOST=localhost
RABBITMQ_PORT=5672
RABBITMQ_USER=guest
RABBITMQ_PASS=guest
```

3. Execute a aplicação:
```bash
go run cmd/ordersystem/main.go cmd/ordersystem/wire_gen.go
```

---

## 📁 Estrutura do Projeto

```
├── cmd/ordersystem/          # Ponto de entrada da aplicação
│   ├── main.go
│   ├── wire.go               # Configuração de DI
│   └── wire_gen.go           # Código gerado pelo Wire
├── internal/
│   ├── entity/               # Entidades de domínio
│   ├── usecase/              # Casos de uso
│   │   ├── create_order.go
│   │   └── list_orders.go
│   └── infra/
│       ├── database/         # Repositórios
│       ├── graph/            # GraphQL
│       ├── grpc/             # gRPC
│       └── web/              # REST API
├── api/
│   └── api.http              # Requests HTTP para testes
├── sql/migrations/           # Scripts de migração
└── docker-compose.yaml
```

---

## 🔧 Comandos de Desenvolvimento

### Regenerar código gRPC
```bash
protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto
```

### Regenerar código GraphQL
```bash
go run github.com/99designs/gqlgen generate
```

### Regenerar Wire (DI)
```bash
go generate ./...
```

### Executar testes
```bash
go test ./...
```

---

## 📋 Funcionalidades

- ✅ **CreateOrder**: Criar pedido via REST, gRPC e GraphQL
- ✅ **ListOrders**: Listar pedidos via REST, gRPC e GraphQL
- ✅ **Event Dispatcher**: Eventos publicados no RabbitMQ ao criar pedido

---

## 🏗️ Arquitetura

O projeto segue **Clean Architecture**:

- **Entity**: Regras de negócio (`Order`)
- **UseCase**: Lógica de aplicação (`CreateOrderUseCase`, `ListOrdersUseCase`)
- **Infra**: Implementações externas (DB, HTTP, gRPC, GraphQL)
- **Dependency Injection**: Google Wire

---

## 📦 Tecnologias

- **Go** 1.24+
- **MySQL** 5.7
- **RabbitMQ** 3
- **gRPC** com Protocol Buffers
- **GraphQL** com gqlgen
- **Docker** e Docker Compose
- **Google Wire** para DI
