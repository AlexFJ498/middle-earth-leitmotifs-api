# Middle‑earth Leitmotifs API

A public API focused on the leitmotifs in Howard Shore’s scores for Middle‑earth, as documented by Doug Adams.

## Table of Contents
- About this project
- Project overview
- Tech stack
- Architecture
- API endpoints
- Setup
- Configuration
- Running locally
- Docker
- Testing
- License
- Acknowledgments

## About this project

This project began as a personal initiative to learn and use modern web development tools through a subject I care about. After reading Doug Adams’ book, it became clear that a concise, well‑structured way to browse the leitmotifs would be both useful and enjoyable. I hope the fans of LOTR music find it helpful as well. Issues and pull requests are welcome.

It is free, non‑commercial, and released under the MIT License. The goal is not to provide exhaustive analysis, but rather a practical, readable interface that offers context without attempting to replace the source material. Suggestions, issues, and pull requests are welcome. This is an independent, non‑commercial fan project, not affiliated with New Line Cinema, Warner Bros., The Saul Zaentz Company, or The Tolkien Estate. Please support the creators by purchasing official editions and licensed soundtracks.

Music compositions and track titles are the property of Howard Shore and the respective rights holders. “Middle‑earth”, film titles, character names, place names, and other related names and terms are the property of The Tolkien Estate and other rights holders.

No sheet music or film audio is hosted; embedded players link only to official Spotify releases. Images on this site were generated using AI and serve only as illustrative placeholders; they are not official artwork.

The thematic names and structure presented here are derived from Doug Adams’ book “The Music of The Lord of the Rings Films”. All original naming remains the property of the author. The brief descriptions were written specifically for this site and are intentionally limited.

## Project overview

This repository contains a Go service that exposes a REST API to manage and query leitmotif‑related data such as movies, groups, categories, tracks, and themes. It uses a clean separation of concerns: HTTP layer (Gin), application use‑cases (commands/queries), domain models, and an SQL persistence layer built on top of database/sql with go‑sqlbuilder.

## Tech stack

- Go 1.24
- Gin (HTTP)
- database/sql + lib/pq (PostgreSQL)
- go‑sqlbuilder (SQL generation)
- JWT (golang‑jwt) and bcrypt for auth
- godotenv + envconfig for configuration

## Architecture

The service follows a simple CQRS‑style separation:

- Domain: core entities and errors under `internal/*.go` (e.g., `theme.go`, `track.go`, `movie.go`).
- Application:
	- Commands for create/update/delete under `internal/creating`, `internal/updating`, `internal/deleting`.
	- Queries for get/list under `internal/getting`, `internal/listing`.
	- In‑memory buses in `internal/platform/bus/inmemory`.
- Infrastructure:
	- HTTP server and handlers in `internal/platform/server` (Gin), with JWT and admin middlewares.
	- SQL repositories in `internal/platform/storage/sqldb` using `go-sqlbuilder` and explicit SQL error mapping (`flavor.go`).
- Composition: the entrypoint `cmd/api/main.go` calls `cmd/api/bootstrap/bootstrap.go`, which wires configuration, DB connection, buses, repositories, and services, then starts the HTTP server.

## API endpoints

Public
- GET `/health`
- POST `/login`
- GET `/movies`, GET `/movies/:id`
- GET `/groups`, GET `/groups/:id`
- GET `/categories`, GET `/categories/:id`
- GET `/tracks`, GET `/tracks/:id`
- GET `/themes`, GET `/themes/:id`
- GET `/themes/group/:group_id`

Protected (JWT + admin)
- Users: POST `/users`, GET `/users`
- Movies: POST `/movies`, PUT `/movies/:id`, DELETE `/movies/:id`
- Groups: POST `/groups`, PUT `/groups/:id`, DELETE `/groups/:id`
- Categories: POST `/categories`, PUT `/categories/:id`, DELETE `/categories/:id`
- Tracks: POST `/tracks`, PUT `/tracks/:id`, DELETE `/tracks/:id`
- Themes: POST `/themes`, PUT `/themes/:id`, DELETE `/themes/:id`

For quick HTTP examples, see the `.rest-client/` folder.

## Setup

Requirements
- Go toolchain
- PostgreSQL 16 (or compatible)

Install dependencies
```powershell
go mod download
```

Database
- Apply the SQL migrations in `db/migrations/` to your database (use your preferred migration tool or run the scripts in order).

## Configuration

Configuration is provided via environment variables (loaded with `envconfig` using the `MELA_` prefix). If `MELA_ENV` is not set, `.env.local` is loaded via `godotenv`.

Common variables
- `MELA_HOST` (e.g., `0.0.0.0`)
- `MELA_PORT` (e.g., `8080`)
- `MELA_SHUTDOWNTIMEOUT` (e.g., `5s`)
- `MELA_DBUSER`, `MELA_DBPASSWORD`, `MELA_DBHOST`, `MELA_DBPORT`, `MELA_DBNAME`, `MELA_DBTIMEOUT`
- `DATABASE_URL` (optional; if present it is used instead of individual DB vars)
- `MELA_JWTKEY`, `MELA_JWTEXPIRES`
- `MELA_FRONTENDURL` (for CORS)

## Running locally

Run the API
```powershell
go run ./cmd/api/main.go
```

The server binds to `MELA_HOST:MELA_PORT`. Graceful shutdown is handled via OS signals and the configured shutdown timeout.
By default (with `MELA_HOST=0.0.0.0` and `MELA_PORT=8080`), the API is available at http://localhost:8080.

## Docker

A multi‑stage Dockerfile is included.

Build and run with Compose (requires appropriate `.env` with `MELA_*` values):
```powershell
docker compose up --build
```

This will start the API container and a PostgreSQL 16 container with volumes. The API exposes port 8080 by default.

## Testing

Run all tests
```powershell
go test ./...
```

The repository includes unit tests for services and SQL repositories. SQL tests use `go-sqlmock` where appropriate.

## License

MIT License. See `LICENSE`.

## Acknowledgments

- You may purchase the Doug Adams’ book here: [The Music of The Lord of the Rings Films](https://www.amazon.com/Music-Lord-Rings-Films-Comprehensive/dp/0739071572?ref_=ast_author_dp).
- This project was made from zero following this [course](https://github.com/CodelyTV/go-hexagonal_http_api-course).