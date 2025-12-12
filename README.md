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

   O projeto utiliza Docker e Docker Compose para gerenciar os serviços de banco de dados, mensageria e aplicação. Para iniciar todos os serviços, execute o comando abaixo:

   ```bash
   docker-compose up -d
   ```

   Este comando irá iniciar os seguintes contêineres:

   - `mysql`: Banco de dados MySQL com as migrations executadas automaticamente.
   - `rabbitmq`: Servidor de mensageria RabbitMQ.
   - `app`: A aplicação Go, construída com um Dockerfile multistage (usando Golang para build e Alpine para runtime).

3. **Verifique os logs:**

   Para ver os logs da aplicação:

   ```bash
   docker-compose logs app
   ```

### Migrations

As migrations do banco de dados são executadas automaticamente durante a inicialização do contêiner MySQL. O arquivo `sql/migrations/001_create_orders_table.sql` cria a tabela `orders` com os campos `id`, `price`, `tax` e `final_price`. Se precisar adicionar novas migrations, coloque os arquivos SQL na pasta `sql/migrations` e reinicie o contêiner MySQL.

## Desenvolvimento Local

Para rodar a aplicação localmente sem Docker:

1. Certifique-se de ter Go instalado e os serviços MySQL e RabbitMQ rodando (pode usar Docker para eles).

2. Crie um arquivo `.env` na raiz do projeto com as seguintes variáveis de ambiente:

   ```
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
   Ajuste os valores conforme sua configuração local.

3. Execute as migrations manualmente no MySQL.

4. Rode o comando:

   ```bash
   cd cmd/ordersystem
   go run main.go wire_gen.go
   ```

Isso iniciará os servidores REST (porta 8000), gRPC (porta 50051) e GraphQL (porta 8080).

## Como Usar

Após iniciar os serviços, a aplicação estará rodando com três APIs disponíveis:

- **API REST:** Porta 8000
- **API gRPC:** Porta 50051
- **API GraphQL:** Porta 8080 (playground em http://localhost:8080)

### API REST

- `POST /order`: Cria um novo pedido.

### API GraphQL

O playground GraphQL permite testar mutações e queries diretamente no navegador.

### API gRPC

Use ferramentas como Evans para interagir com o servidor gRPC.

## Exemplos de Chamadas

### REST

Use ferramentas como Postman, curl ou a extensão REST Client do VS Code.

**Criar um novo pedido:**

```http
POST http://localhost:8000/order
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

Acesse o playground em http://localhost:8080 e execute a seguinte mutação:

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
  "finalPrice": 105.99,
  "id": "123",
  "price": 100.99,
  "tax": 5
}
```

## Estrutura do Projeto

```
.
├── Dockerfile
├── GEMINI.md
├── README.md
├── api
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
│   │   │   ├── generated.go
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
│       └── create_order.go
├── pkg
│   └── events
│       ├── event_dispatcher.go
│       ├── event_dispatcher_test.go
│       └── interface.go
├── sql
│   └── migrations
│       └── 001_create_orders_table.sql
└── tools.go
```
