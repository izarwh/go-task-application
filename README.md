# Go Task Application

A task management REST API built with Go, featuring PostgreSQL as the primary database, Redis for caching, and Chi as the web framework.

## Tech Stack
- **Go**: Version 1.25.5
- **Web Framework**: [Chi](https://go-chi.io/)
- **Database**: PostgreSQL (via [GORM](https://gorm.io/))
- **Cache**: Redis (via `go-redis/v9`)
- **Configuration Management**: [Viper](https://github.com/spf13/viper)

## Features
- **Task Management**: Create, Read, Update, and Soft-delete tasks.
- **Caching**: Get/Delete operations are natively cached and invalidated through Redis to ensure snappy request responses and data consistency.
- **RESTful API**: Served over Chi with standard JSON responses.
- **Health Check**: Easy deployment verification using the `/health` endpoint.

## Prerequisites
Before running the project, make sure you have the following installed on your machine:
- [Go 1.25+](https://go.dev/doc/install)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Redis](https://redis.io/download)

## Setup & Installation

1. **Clone the repository** (or navigate to your project directory):
   ```bash
   cd go-task-application
   ```

2. **Install dependencies**:
   ```bash
   go mod tidy
   ```

3. **Configure the Environment**:
   Copy the provided `.env.example` to `.env` and fill in your connection details for Postgres and Redis.
   ```bash
   cp .env.example .env
   ```
   *Make sure your PostgreSQL and Redis services are running matching the credentials inside the `.env` file.*

## Running the Application

To start the API server locally:

```bash
go run ./cmd/api/main.go
```

The server should start on the port configured in your `.env` file (default is usually `8080`).

You can verify the API is running by hitting the health check endpoint:
```bash
curl http://localhost:8080/health
```

## Building for Production

Compile a binary for production deployment:
```bash
go build -o app ./cmd/api
```
Run the compiled binary:
```bash
./app
```
