.PHONY: run tidy build migrate-up migrate-down swag

run:
	go run ./cmd/api

tidy:
	go mod tidy

build:
	go build -o bin/api ./cmd/api

migrate-up:
	go run github.com/pressly/goose/v3/cmd/goose@latest -dir migrations sqlite3 data/app.db up

migrate-down:
	go run github.com/pressly/goose/v3/cmd/goose@latest -dir migrations sqlite3 data/app.db down

swag:
	swag init -g cmd/api/main.go -o docs --parseDependency --parseInternal
