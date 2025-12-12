# Desafio Clean Architecture

Este repositório contém a implementação de um sistema de gerenciamento de pedidos como parte do desafio do curso Full Cycle Go Expert.

## Arquitetura

O projeto segue os princípios da Clean Architecture, proposta por Robert C. Martin (Uncle Bob), que visa criar sistemas independentes de frameworks, testáveis e fáceis de manter. A arquitetura é dividida em camadas concêntricas, onde as camadas internas não dependem das externas:

- **Entidades (Entity)**: Representam as regras de negócio centrais. No projeto, a entidade `Order` encapsula a lógica de cálculo do preço final.
- **Casos de Uso (Use Case)**: Contêm a lógica de aplicação, orquestrando as entidades e interfaces. O `CreateOrderUseCase` gerencia a criação de pedidos.
- **Infraestrutura (Infra)**: Implementa as interfaces definidas nas camadas internas. Inclui repositórios de banco de dados, handlers web, GraphQL, gRPC, etc.
- **Adaptadores Externos**: Camada mais externa, como frameworks (Gin para web, gqlgen para GraphQL), bancos de dados (MySQL), mensageria (RabbitMQ).

Essa separação permite que o código seja testável, com dependências injetadas via Google Wire, e facilita mudanças em tecnologias externas sem afetar o núcleo do negócio.

## Tecnologias Utilizadas

- **Linguagem:** Go
- **Banco de Dados:** MySQL
- **Mensageria:** RabbitMQ
- **APIs:**
  - REST (usando Gin)
  - gRPC
  - GraphQL (usando gqlgen)
- **Injeção de Dependência:** Google Wire
- **Containerização:** Docker e Docker Compose

## Como Executar

Para executar o projeto, siga os passos abaixo:

1. **Clone o repositório:**

   ```bash
   git clone https://github.com/devfullcycle/go-expert-desafio-clean-arch.git
   cd go-expert-desafio-clean-arch
   ```

2. **Inicie os serviços:**

   O projeto utiliza Docker e Docker Compose para gerenciar os serviços de banco de dados e mensageria. Para iniciar todos os serviços, execute o comando abaixo:

   ```bash
   docker-compose up -d
   ```

   Este comando irá iniciar os seguintes contêineres:

   - `mysql`: Banco de dados MySQL
   - `rabbitmq`: Servidor de mensageria RabbitMQ
   - `app`: A aplicação Go

3. **Verifique os logs:**

   Para ver os logs da aplicação:

   ```bash
   docker-compose logs app
   ```

## Como Usar

Após iniciar os serviços, a aplicação estará rodando com três APIs disponíveis:

- **API REST:** Porta 8080
- **API gRPC:** Porta 50051
- **API GraphQL:** Porta 8082 (playground em http://localhost:8082)

### API REST

- `POST /order`: Cria um novo pedido.
- `GET /order`: Lista todos os pedidos (não implementado no código atual, mas mencionado no README original).

### API GraphQL

O playground GraphQL permite testar mutações e queries diretamente no navegador.

### API gRPC

Use ferramentas como Evans para interagir com o servidor gRPC.

## Exemplos de Chamadas

### REST

Use ferramentas como Postman, curl ou a extensão REST Client do VS Code.

**Criar um novo pedido:**

```http
POST http://localhost:8080/order
Content-Type: application/json

{
  "id": "123",
  "price": 100.0,
  "tax": 10.0
}
```

**Resposta esperada:**

```json
{
  "id": "123",
  "price": 100.0,
  "tax": 10.0,
  "final_price": 110.0
}
```

### GraphQL

Acesse o playground em http://localhost:8082 e execute a seguinte mutação:

```graphql
mutation {
  createOrder(input: {
    id: "e",
    Price: 2,
    Tax: 1
  }) {
    id
    Price
    Tax
    FinalPrice
  }
}
```

**Resposta esperada:**

```json
{
  "data": {
    "createOrder": {
      "id": "e",
      "Price": 2,
      "Tax": 1,
      "FinalPrice": 3
    }
  }
}
```

### gRPC

Para testar via gRPC, instale o Evans:

```bash
go install github.com/ktr0731/evans@latest
```

Em seguida, conecte ao servidor:

```bash
evans -r repl localhost:50051
```

No REPL do Evans, chame o serviço:

```
evans > call CreateOrder
id (TYPE_STRING) => 123
price (TYPE_FLOAT) => 100.0
tax (TYPE_FLOAT) => 10.0
```

**Resposta esperada:**

```
{
  "id": "123",
  "price": 100,
  "tax": 10,
  "final_price": 110
}
```

## Estrutura do Projeto

```
.
├── api
│   ├── api.http
│   └── create_order.http
├── cmd
│   └── ordersystem
│       ├── main.go
│       ├── wire.go
│       └── wire_gen.go
├── configs
│   └── config.go
├── docker-compose.yaml
├── go.mod
├── go.sum
├── gqlgen.yml
├── internal
│   ├── entity
│   │   ├── interface.go
│   │   ├── order.go
│   │   └── order_test.go
│   ├── event
│   │   ├── handler
│   │   │   └── order_created_handler.go
│   │   └── order_created.go
│   ├── infra
│   │   ├── database
│   │   │   ├── order_repository.go
│   │   │   └── order_repository_test.go
│   │   ├── graph
│   │   │   ├── generated
│   │   │   │   └── generated.go
│   │   │   ├── model
│   │   │   │   └── models_gen.go
│   │   │   ├── resolver.go
│   │   │   ├── schema.graphqls
│   │   │   └── schema.resolvers.go
│   │   ├── grpc
│   │   │   ├── pb
│   │   │   │   ├── order.pb.go
│   │   │   │   └── order_grpc.pb.go
│   │   │   ├── protofiles
│   │   │   │   └── order.proto
│   │   │   └── service
│   │   │       └── order_service.go
│   │   └── web
│   │       ├── order_handler.go
│   │       └── webserver
│   │           ├── starter.go
│   │           └── webserver.go
│   └── usecase
│       ├── create_order.go
│       └── list_orders.go
├── pkg
│   └── events
│       ├── event_dispatcher.go
│       ├── event_dispatcher_test.go
│       └── interface.go
├── README.md
├── sql
│   └── migrations
│       └── create_orders_table.sql
└── tools.go
```
