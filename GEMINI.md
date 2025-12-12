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

2.  Run the application:

    ```bash
    go run cmd/ordersystem/main.go
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

Database migrations are located in the `sql/migrations` directory. They are automatically applied when the application starts.
