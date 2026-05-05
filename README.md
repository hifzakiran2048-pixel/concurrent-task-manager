# Concurrent Task Manager

A command-line task manager written in Go. The project stores tasks in PostgreSQL and uses goroutines, channels, and a worker pool to mark tasks as completed asynchronously.

## Features

- Add new tasks from the CLI
- List all saved tasks
- Mark tasks as done through background workers
- Delete tasks by ID
- Persist tasks in a PostgreSQL database
- Layered project structure with database, repository, service, model, and worker packages

## Tech Stack

- Go
- PostgreSQL
- `database/sql`
- `github.com/lib/pq`
- Goroutines, channels, and `sync.WaitGroup`

## Project Structure

```text
.
|-- database/
|   `-- connection.go        # PostgreSQL connection setup
|-- model/
|   `-- task.go              # Task data model
|-- repository/
|   `-- task_repository.go   # Database CRUD operations
|-- service/
|   `-- task_service.go      # Business logic layer
|-- worker/
|   `-- worker.go            # Worker pool task processing
|-- main.go                  # CLI menu and application entry point
|-- go.mod
`-- README.md
```

## Database Setup

Create a PostgreSQL database named `taskdb`:

```sql
CREATE DATABASE taskdb;
```

Create the `tasks` table:

```sql
CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    is_done BOOLEAN NOT NULL DEFAULT FALSE
);
```

The current database connection is configured in `database/connection.go`:

```go
connStr := "user=postgres password=Admin dbname=taskdb sslmode=disable"
```

Update the username, password, database name, or SSL mode if your local PostgreSQL settings are different.

## Run the Project

Install dependencies:

```bash
go mod tidy
```

Run the application:

```bash
go run main.go
```

## CLI Menu

After starting the app, you can choose from:

```text
1. Add Task
2. List Tasks
3. Mark Done (Worker)
4. Delete Task
5. Exit
```

## How It Works

1. `main.go` connects to PostgreSQL.
2. The repository layer handles database operations.
3. The service layer exposes task actions to the CLI and workers.
4. The app starts two worker goroutines.
5. When you select `Mark Done`, the task ID is sent into a channel.
6. A worker receives the task, marks it complete in the database, and sends the result back.

## Concepts Practiced

- Clean layered architecture
- PostgreSQL CRUD operations in Go
- Goroutines
- Channels
- Worker pool pattern
- WaitGroup-based shutdown
