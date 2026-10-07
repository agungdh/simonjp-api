# simonjp-api

Golang BE sederhana: `go-chi` + `sqlite` + `bun` + `goose`.

## Quick start

```sh
cp .env.example .env
go mod tidy
go run ./cmd/api
```

Server: `http://localhost:8080`

Swagger UI: `http://localhost:8080/swagger/index.html`

Regenerate docs setelah ubah anotasi:

```sh
make swag
```

## Endpoints

- `GET /health`
- `GET /api/items`
- `POST /api/items` — body `{"name":"...","description":"..."}`
- `GET /api/items/{id}`
- `PUT /api/items/{id}`
- `DELETE /api/items/{id}`

## Migrasi

Auto-migrate saat boot via `goose.Up`. Manual:

```sh
make migrate-up
make migrate-down
```

DB default: `data/app.db` (di-gitignore). Ubah via `DB_PATH` / `PORT` di `.env`.
