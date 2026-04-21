# Blog REST API

A Blog REST API built with [Go Fiber](https://gofiber.io/), [GORM](https://gorm.io/), and PostgreSQL. API docs are auto-generated at runtime using [Huma](https://huma.rocks/) — no code generation step required.

## Stack

- **Framework:** Go Fiber v2
- **ORM:** GORM
- **Database:** PostgreSQL
- **API Docs:** Huma v2 (OpenAPI 3.1, auto-generated from Go types)

## Project Structure

```
goFiber-app/
├── main.go               # App entry, middleware, route registration
├── database/
│   └── database.go       # GORM connection + AutoMigrate
├── models/
│   └── post.go           # Post model
├── handlers/
│   └── post.go           # CRUD handlers + I/O types + route registration
├── docker-compose.yml    # PostgreSQL via Docker
└── .env                  # Environment variables
```

## Getting Started

### Prerequisites

- Go 1.21+
- PostgreSQL (local or Docker)

### Option A — Local PostgreSQL

```bash
createdb blog_db
```

Update `.env` with your local credentials:

```env
DB_HOST=localhost
DB_USER=your_user
DB_PASSWORD=
DB_NAME=blog_db
DB_PORT=5432
APP_PORT=3000
```

### Option B — Docker

```bash
docker compose up -d
```

This starts a PostgreSQL container with the credentials already matching `.env`.

### Run

```bash
go run main.go
```

GORM's `AutoMigrate` creates the `posts` table automatically on first start.

## API Docs

Interactive docs (Stoplight Elements) are served at:

```
http://localhost:3000/docs
```

OpenAPI spec available at:

```
http://localhost:3000/openapi.yaml
http://localhost:3000/openapi.json
```

## Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/posts` | List all posts |
| `GET` | `/posts/:id` | Get a post by ID |
| `POST` | `/posts` | Create a post |
| `PUT` | `/posts/:id` | Update a post |
| `DELETE` | `/posts/:id` | Delete a post |

### Example — Create a post

```bash
curl -X POST http://localhost:3000/posts \
  -H "Content-Type: application/json" \
  -d '{"title": "Hello World", "body": "My first post", "author": "mac"}'
```

### Example — List posts

```bash
curl http://localhost:3000/posts
```
