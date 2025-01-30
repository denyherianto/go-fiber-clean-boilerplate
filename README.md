# Go Fiber Clean Architecture Boilerplate

[Fiber](https://gofiber.io/) is an Express.js inspired web framework build on top of Fasthttp, the fastest HTTP engine for Go. Designed to ease things up for **fast** development with **zero memory allocation** and **performance** in mind.

## Features

- **Clean Architecture** for building scalable projects
- **JWT** for authentication
- **PostgreSQL** for database
- **Redis** for cache
- **Swagger** for auto-generating API Docs
- **Golang Migrate** for database migrations
- **GolangCI-Lint** for Go linter issues
- **GoCritic** for Go the best practice issues
- **Gosec** for Go security issues


## Prerequisite
1. PostgresSQL
   
   Create new database:

   ```
   sudo -i -u postgres psql
   CREATE DATABASE silog_office;
   GRANT ALL PRIVILEGES ON DATABASE silog_office TO postgres;
   \q
   ``` 

2. Redis

   https://redis.io/docs/latest/operate/oss_and_stack/install/install-redis/install-redis-on-windows/


## ⚡️ Quick start

1. Copy `.env.example` to `.env` and fill it with your environment values.
2. Run `go get`
3. Install following useful Go tools to your system:

   - [golang-migrate/migrate](https://github.com/golang-migrate/migrate#cli-usage) for apply migrations
   - [github.com/swaggo/swag](https://github.com/swaggo/swag) for auto-generating Swagger API docs
   - [github.com/securego/gosec](https://github.com/securego/gosec) for checking Go security issues
   - [github.com/go-critic/go-critic](https://github.com/go-critic/go-critic) for checking Go the best practice issues
   - [github.com/golangci/golangci-lint](https://github.com/golangci/golangci-lint) for checking Go linter issues

3. Run project by this command:

```bash
go run main.go
```

4. Go to API Docs page (Swagger): [127.0.0.1:5001/swagger/index.html](http://127.0.0.1:5001/swagger/index.html)

![Screenshot](https://user-images.githubusercontent.com/11155743/112715187-07dab100-8ef0-11eb-97ea-68d34f2178f6.png)



## 🗄 Template structure

### ./app

**Folder with business logic and specific functionality**.

- `./app/controllers` folder for functional controllers (used in routes)
- `./app/constants` folder for describe `const` of your project
- `./app/models` folder for describe business models and methods of your project
- `./app/queries` folder for describe queries for models of your project
- `./app/middlewares` folder for add middleware (Fiber built-in and yours)
- `./app/routes` folder for describe routes of your project
- `./app/utils` folder with utility functions (server starter, error checker, etc)

### ./configs

**Folder with Configurations only**.

- `./configs/cache` folder with in-memory cache setup functions (by default, Redis)
- `./configs/database` folder with database setup functions (by default, PostgreSQL)
- `./configs/migrations` folder with migration files (used with [golang-migrate/migrate](https://github.com/golang-migrate/migrate) tool)

### ./docs

**Folder with API Documentation**. This directory contains config files for auto-generated API Docs by Swagger.

## ⚙️ Configuration

```ini
# .env

# Stage status to start server:
#   - "dev", for start server without graceful shutdown
#   - "prod", for start server with graceful shutdown
STAGE_STATUS="dev"

# Server settings:
SERVER_HOST="0.0.0.0"
SERVER_PORT=5001
SERVER_READ_TIMEOUT=60

# JWT settings:
JWT_SECRET_KEY="secret"
JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT=15
JWT_REFRESH_KEY="refresh"
JWT_REFRESH_KEY_EXPIRE_HOURS_COUNT=720

# Database settings:
DB_HOST="127.0.0.1"
DB_PORT=5432
DB_USER="postgres"
DB_PASSWORD="password"
DB_DBNAME="postgres"
DB_SSL_MODE="disable"
DB_MAX_CONNECTIONS=100
DB_MAX_IDLE_CONNECTIONS=10
DB_MAX_LIFETIME_CONNECTIONS=2

# Redis settings:
REDIS_HOST="127.0.0.1"
REDIS_PORT=6379
REDIS_PASSWORD=""
REDIS_DB_NUMBER=0
```

## ⚠️ License

Apache 2.0 &copy; [Vic Shóstak](https://shostak.dev/) & [True web artisans](https://1wa.co/).
