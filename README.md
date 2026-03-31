<img width="1077" height="714" alt="image" src="https://github.com/user-attachments/assets/25398407-0ae8-4466-8cb0-a2baebc8426b" /># Go Task Application

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

## Evidence
- **Success Input Data**
  <img width="1077" height="714" alt="image" src="https://github.com/user-attachments/assets/745fa947-0822-4154-9e31-83ea82aa1baf" />

- **Success Get Tasks**
  <img width="1078" height="921" alt="image" src="https://github.com/user-attachments/assets/6373bc44-2951-4276-b528-038a2db2a269" />

- **Success Get Task By ID**
  <img width="1082" height="835" alt="image" src="https://github.com/user-attachments/assets/7e7cac2b-af07-413a-a168-d40cea219905" />

- **Success Delete**
  <img width="1082" height="686" alt="image" src="https://github.com/user-attachments/assets/f6d3b505-1821-477c-8804-347a3504cf76" />

- **Success Update**
  <img width="1074" height="755" alt="image" src="https://github.com/user-attachments/assets/83f8b94c-b673-4083-8470-93891626c3b3" />
