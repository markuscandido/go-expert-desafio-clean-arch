# Project Overview

This project is a Go application that manages orders, following the principles of Clean Architecture. It exposes three different APIs: REST, gRPC, and GraphQL. The application is containerized using Docker and uses MySQL for the database and RabbitMQ for messaging.

## Building and Running

### With Docker

To build and run the application using Docker, run the following command:

```bash
docker-compose up -d
```

This will start the application, the MySQL database, and the RabbitMQ message broker.

### Locally

To run the application locally, you need to have Go, MySQL, and RabbitMQ installed.

1.  Install dependencies:

    ```bash
    go mod tidy
    ```

2.  Create a `.env` file in the root of the project with the following variables:

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

3.  Run the application:

    ```bash
    go run cmd/ordersystem/main.go cmd/ordersystem/wire_gen.go
    ```

## APIs

The application exposes three different APIs: REST, gRPC, and GraphQL.

### GraphQL

The project uses [gqlgen](https://gqlgen.com/) to implement the GraphQL API. The GraphQL playground is available at [http://localhost:8080](http://localhost:8080) when running the application locally.

To create a new order, you can use the following mutation:

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

### gRPC

The project uses [gRPC](https://grpc.io/) to implement the gRPC API. The gRPC server is running on port `50051`.

To interact with the gRPC server, you can use a tool like [Evans](https://github.com/ktr0731/evans). To install Evans, run the following command:

```bash
go install github.com/ktr0731/evans@latest
```

Once installed, you can connect to the server with the following command:

```bash
evans -r repl localhost:50051
```

In the Evans REPL, you can call the `CreateOrder` service:

```
evans > call CreateOrder
id (TYPE_STRING) => 123
price (TYPE_FLOAT) => 100.0
tax (TYPE_FLOAT) => 10.0
```

## Development Conventions

### Testing

To run the tests, use the following command:

```bash
go test ./...
```

### Dependency Injection

The project uses Google Wire for dependency injection. To generate the wire code, run the following command:

```bash
go generate ./...
```

### Database Migrations

Database migrations are located in the `sql/migrations` directory. They are automatically applied when the application starts if you are using Docker. If you are running the application locally, you need to run the migrations manually.
