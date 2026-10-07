package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"

	"github.com/agungdh/simonjp-api/internal/config"
	"github.com/agungdh/simonjp-api/internal/db"
	"github.com/agungdh/simonjp-api/internal/routes"

	_ "github.com/agungdh/simonjp-api/docs"
)

// @title           simonjp-api
// @version         1.0
// @description     Simple CRUD API: chi + sqlite + bun + goose.
// @host            localhost:8080
// @BasePath        /
// @schemes         http
func main() {
	cfg := config.Load()

	bunDB, err := db.New(cfg.DBPath)
	if err != nil {
		log.Fatalf("open db: %v", err)
	}
	defer bunDB.Close()

	if err := goose.SetDialect("sqlite3"); err != nil {
		log.Fatalf("goose dialect: %v", err)
	}
	if err := goose.Up(bunDB.DB, "migrations"); err != nil {
		log.Fatalf("migrate: %v", err)
	}

	r := routes.New(bunDB)

	addr := ":" + cfg.Port
	fmt.Printf("listening on http://localhost%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
