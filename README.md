# Desafio Clean Architecture

Este repositório contém a implementação de um sistema de gerenciamento de pedidos como parte do desafio do curso Full Cycle Go Expert.

## Arquitetura

O projeto segue os princípios da Clean Architecture, com uma clara separação de responsabilidades entre as camadas de domínio, aplicação e infraestrutura.

## Tecnologias Utilizadas

-   **Linguagem:** Go
-   **Banco de Dados:** MySQL
-   **Mensageria:** RabbitMQ
-   **APIs:**
    -   REST
    -   gRPC
    -   GraphQL
-   **Injeção de Dependência:** Google Wire

## Como Executar

Para executar o projeto, siga os passos abaixo:

1.  **Clone o repositório:**

    ```bash
    git clone https://github.com/devfullcycle/go-expert-desafio-clean-arch.git
    cd go-expert-desafio-clean-arch
    ```

2.  **Inicie os serviços:**

    O projeto utiliza Docker e Docker Compose para gerenciar os serviços de banco de dados e mensageria. Para iniciar todos os serviços, execute o comando abaixo:

    ```bash
    docker-compose up -d
    ```

    Este comando irá iniciar os seguintes contêineres:

    -   `mysql`: Banco de dados MySQL
    -   `rabbitmq`: Servidor de mensageria RabbitMQ
    -   `app`: A aplicação Go

3.  **Acesse as APIs:**

    -   **API REST:** A API REST estará disponível em `http://localhost:8080`.
        -   `POST /order`: Cria um novo pedido.
        -   `GET /order`: Lista todos os pedidos.
    -   **API gRPC:** O servidor gRPC estará disponível na porta `50051`.
    -   **API GraphQL:** O playground do GraphQL estará disponível em `http://localhost:8082`.

## Testando a Aplicação

Você pode utilizar o arquivo `api.http` para testar a API REST com a extensão REST Client do VS Code.

### Exemplo de Requisição

**Criar um novo pedido:**

```http
POST http://localhost:8080/order
Content-Type: application/json

{
  "id": "1",
  "price": 100.0,
  "tax": 10.0
}
```

**Listar todos os pedidos:**

```http
GET http://localhost:8080/order
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
