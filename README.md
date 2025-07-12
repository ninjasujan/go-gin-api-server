# Go Gin API Server

A RESTful API server built with Go and Gin framework with database migrations using golang-migrate.

## 🚀 Features

- RESTful API with CRUD operations
- JSON responses with proper HTTP status codes
- Built with Gin - Fast HTTP web framework
- Database migrations with golang-migrate
- SQLite database support

## 📋 Prerequisites

- Go 1.24.2 or higher
- Git
- golang-migrate CLI tool

## 🛠️ Installation & Setup

### 1. Install golang-migrate

**On macOS:**

```bash
brew install golang-migrate
```

**On Ubuntu/Debian:**

```bash
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz
sudo mv migrate /usr/local/bin/migrate
```

**On Windows:**

```bash
# Using Chocolatey
choco install migrate

# Using Scoop
scoop install migrate
```

**Manual Installation:**

```bash
# Download the latest version
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz
sudo mv migrate /usr/local/bin/migrate
```

### 2. Install Database Drivers

For SQLite (already included with golang-migrate):

```bash
# SQLite driver is built-in, no additional installation needed
```

For other databases:

```bash
# PostgreSQL
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# MySQL
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# SQL Server
go install -tags 'sqlserver' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

### 3. Clone and Setup Project

```bash
git clone https://github.com/yourusername/go-gin-api-server.git
cd go-gin-api-server
go mod tidy
```

### 4. Run Database Migrations

```bash
# Run all pending migrations
migrate -path database/migrations -database "sqlite://test.db" up

# Rollback the last migration
migrate -path database/migrations -database "sqlite://test.db" down 1

# Rollback all migrations
migrate -path database/migrations -database "sqlite://test.db" down

# Check migration status
migrate -path database/migrations -database "sqlite://test.db" version
```

### 5. Run the Server

```bash
go run main.go
```

The server will start on `http://localhost:8080`

## 📁 Project Structure

```
go-gin-api-server/
├── main.go                    # Server entry point
├── routes/
│   └── routes.go             # API route definitions
├── database/
│   ├── migrations/           # Database migration files
│   │   ├── 000001_create_user_table.up.sql
│   │   └── 000001_create_user_table.down.sql
│   └── sqlite.go            # Database connection
├── models/
│   └── book.go              # Data models
├── controller/
│   └── controller.go        # Request handlers
├── seeder/
│   └── seeder.go           # Database seeders
├── go.mod                   # Go module file
├── test.db                  # SQLite database file
└── README.md               # This file
```

## 🔄 Database Migrations

### Migration Commands

```bash
# Basic syntax
migrate -path <migration-files-path> -database <database-url> <command>

# Examples for SQLite
migrate -path database/migrations -database "sqlite://test.db" up
migrate -path database/migrations -database "sqlite://test.db" down
migrate -path database/migrations -database "sqlite://test.db" version

# Examples for PostgreSQL
migrate -path database/migrations -database "postgres://username:password@localhost:5432/dbname?sslmode=disable" up

# Examples for MySQL
migrate -path database/migrations -database "mysql://username:password@tcp(localhost:3306)/dbname" up
```

### Creating New Migrations

```bash
# Create a new migration
migrate create -ext sql -dir database/migrations -seq create_books_table

# This creates two files:
# - 000002_create_books_table.up.sql
# - 000002_create_books_table.down.sql
```

### Migration File Naming Convention

Migrations follow this naming pattern:

- `{version}_{description}.up.sql` - Migration to apply
- `{version}_{description}.down.sql` - Migration to rollback

Example:

```
000001_create_user_table.up.sql
000001_create_user_table.down.sql
000002_create_books_table.up.sql
000002_create_books_table.down.sql
```

### Example Migration Files

**Up Migration (000001_create_user_table.up.sql):**

```sql
CREATE TABLE `users` (
    `id` VARCHAR(36) NOT NULL PRIMARY KEY,
    `name` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255) NOT NULL UNIQUE,
    `password` VARCHAR(255) NOT NULL
);
```

**Down Migration (000001_create_user_table.down.sql):**

```sql
DROP TABLE IF EXISTS `users`;
```

### Useful Migration Commands

```bash
# Force version (useful for fixing dirty state)
migrate -path database/migrations -database "sqlite://test.db" force VERSION

# Drop database and recreate
migrate -path database/migrations -database "sqlite://test.db" drop

# Check for dirty state
migrate -path database/migrations -database "sqlite://test.db" version

# Run specific number of migrations
migrate -path database/migrations -database "sqlite://test.db" up 2

# Rollback specific number of migrations
migrate -path database/migrations -database "sqlite://test.db" down 1
```

## 🛠️ Technologies Used

- **Go** - Programming language
- **Gin** - Web framework
- **GORM** - ORM library
- **SQLite** - Database
- **golang-migrate** - Database migration tool
- **JSON** - Data format

## 📝 License

This project is open source and available under the [MIT License](LICENSE).

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

---

**Happy coding! 🎉**
