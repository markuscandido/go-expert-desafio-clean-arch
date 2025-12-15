# Project Overview

This project is a Go-based order management system that follows Clean Architecture principles. It exposes three different APIs for interacting with the system: REST, gRPC, and GraphQL. The system uses MySQL as its database, RabbitMQ for event-driven communication, and Google Wire for dependency injection.

The project is well-structured, with a clear separation of concerns:
- `internal/entity`: Contains the core domain entities, such as the `Order` entity.
- `internal/usecase`: Contains the application's business logic, such as creating and listing orders.
- `internal/infra`: Contains the infrastructure-level implementations, such as database repositories, web handlers, gRPC services, and GraphQL resolvers.
- `cmd/ordersystem`: Contains the main application entry point, where all the components are wired together and the servers are started.

# Building and Running

## Using Docker (Recommended)

The easiest way to run the project is by using Docker Compose:

```bash
docker-compose up -d
```

This will start the following services:
- `mysql`: The MySQL database.
- `rabbitmq`: The RabbitMQ message broker.
- `app`: The Go application.

You can then access the different APIs on the following ports:
- **REST API**: `http://localhost:8000`
- **gRPC API**: `localhost:50051`
- **GraphQL API**: `http://localhost:8080`

## Locally

To run the project locally, you need to have Go, Docker, and Docker Compose installed.

1. **Start the database and message broker:**
   ```bash
   docker-compose up -d mysql rabbitmq
   ```

2. **Create a `.env` file** with the following content:
   ```env
   DB_DRIVER=mysql
   DB_HOST=localhost
   DB_PORT=3306
   DB_USER=root
   DB_PASSWORD=root
   DB_NAME=orders
   WEB_SERVER_PORT=:8000
   GRPC_SERVER_PORT=50051
   GRAPHQL_SERVER_PORT=8080
   RABBITMQ_HOST=localhost
   RABBITMQ_PORT=5672
   RABBITMQ_USER=guest
   RABBITMQ_PASS=guest
   ```

3. **Run the application:**
   ```bash
   go run cmd/ordersystem/main.go
   ```

# Development Conventions

## Testing

To run the tests, use the following command:

```bash
go test ./...
```

## Code Generation

The project uses code generation for gRPC, GraphQL, and dependency injection.

- **gRPC:** To regenerate the gRPC code, run the following command:
  ```bash
  protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto
  ```

- **GraphQL:** To regenerate the GraphQL code, run the following command:
  ```bash
  go run github.com/99designs/gqlgen generate
  ```

- **Dependency Injection (Wire):** To regenerate the dependency injection code, run the following command:
  ```bash
  go generate ./...
  ```
