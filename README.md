# Messaging Service

This project implements an automatic messaging service that retrieves unsent messages from a database and sends them to recipients at regular intervals. The system is built with Golang and includes features like database integration, Redis caching, and Swagger documentation.

---

## Features

- Automatically sends unsent messages every 2 minutes.
- Messages sent are marked as "sent" in the database.
- Sent message IDs and timestamps are cached in Redis.
- Provides REST API endpoints to:
  - Start and stop the message-sending scheduler.
  - Retrieve a list of sent messages.
- Fully documented with Swagger.

---

## Technologies Used

- **Programming Language**: Go (Golang)
- **Database**: PostgreSQL
- **Caching**: Redis
- **API Documentation**: Swagger
- **Task Scheduling**: Goroutines

---

## Running with Docker

### Prerequisites

- Docker installed on your machine

### Steps

1. Build the Docker image:
   ```bash
   docker build -t messaging-service .
   ```

2. Create a Docker network (if not already created):
   ```bash
   docker network create messaging-network
   ```

3. Run PostgreSQL container:
   ```bash
   docker run --name postgres-db --network messaging-network -e POSTGRES_USER=your_user -e POSTGRES_PASSWORD=your_password -e POSTGRES_DB=messaging_db -p 5432:5432 -d postgres
   ```

4. Run Redis container:
   ```bash
   docker run --name redis-cache --network messaging-network -p 6379:6379 -d redis
   ```

5. Run the messaging service container:
   ```bash
   docker run --name messaging-service --network messaging-network -p 8080:8080 -d messaging-service
   ```

6. Access the application:
   - The API should be accessible at `http://localhost:8080`.
   - Swagger documentation should be available at `http://localhost:8080/swagger/index.html`.

---

## Directory Structure

```
project/
├── cmd/
├── internal/
│   ├── api/            # HTTP handlers and API routes
│   ├── database/       # Database connection and queries
│   ├── messaging/      # Messaging logic and scheduling
│   ├── httpclient/     # HTTP client for external message sending
│   ├── redis/          # Redis caching logic
│   ├── models/         # Data models and schemas
│   ├── scheduler/      # Scheduler for sending messages at intervals
│   ├── tests/          # Unit and integration tests
├── config/             # Configuration files
├── docs/               # Swagger-generated API documentation
```

---

## Installation and Setup

### Prerequisites

- Go version 2.23 or later
- PostgreSQL
- Redis

### Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/messaging-service.git
   cd messaging-service
   ```

2. Set up the database:
   - Create a PostgreSQL database.
   - Run the following SQL commands to create the `messages` table:
     ```sql
     CREATE TABLE messages (
         id SERIAL PRIMARY KEY,
         content TEXT NOT NULL,
         recipient VARCHAR(20) NOT NULL,
         status VARCHAR(10) DEFAULT 'unsent',
         sent_at TIMESTAMP NULL
     );
     ```
   - Insert sample data if needed.

3. Create a `config.yml` file in the root directory and add the following configuration:
    ```yaml
    database:
      host: localhost  # Change this if your database is hosted on a different server
      port: 5432
      user: your_user
      password: your_password
      name: messaging_db
    redis:
      addr: localhost:6379  # Change this if Redis is hosted on a different server
    ```

4. Install dependencies:
   ```bash
   go mod tidy
   ```

5. Install `swag` if you haven't already:
   ```bash
   go install github.com/swaggo/swag/cmd/swag@latest
   ```

6. Generate Swagger documentation:
   ```bash
   swag init -g cmd/main.go
   ```

7. Run the application:
   ```bash
   go run cmd/main.go
   ```

---

## API Endpoints

### Start Scheduler
- **Endpoint**: `GET /start-messages`
- **Description**: Starts the automatic message-sending process.

### Stop Scheduler
- **Endpoint**: `GET /stop-messages`
- **Description**: Stops the automatic message-sending process.

### Get Sent Messages
- **Endpoint**: `GET /sent-messages`
- **Description**: Retrieves a list of all sent messages.

---

## Testing

1. Use `curl` or tools like Postman to test the API endpoints.
2. Verify messages are being sent and cached in Redis.
3. Check the database to ensure message statuses are updated.

---

Happy coding!
