.DEFAULT_GOAL := help
.PHONY: help run tidy build migrate-up migrate-down db-reset swag

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-12s %s\n", $$1, $$2}'

run: ## Run the API server
	go run ./cmd/api

tidy: ## Tidy go modules
	go mod tidy

build: ## Build binary to bin/api
	go build -o bin/api ./cmd/api

migrate-up: ## Apply all migrations to data/app.db
	go run github.com/pressly/goose/v3/cmd/goose@latest -dir migrations sqlite3 data/app.db up

migrate-down: ## Roll back the last migration
	go run github.com/pressly/goose/v3/cmd/goose@latest -dir migrations sqlite3 data/app.db down

db-reset: ## Delete data/app.db and re-apply all migrations
	rm -f data/app.db data/app.db-wal data/app.db-shm
	$(MAKE) migrate-up

swag: ## Regenerate swagger docs
	swag init -g cmd/api/main.go -o docs --parseDependency --parseInternal
