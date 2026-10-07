package db

import (
	"database/sql"
	"os"
	"path/filepath"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	_ "modernc.org/sqlite"
)

func New(dbPath string) (*bun.DB, error) {
	if dir := filepath.Dir(dbPath); dir != "." && dir != "" {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return nil, err
		}
	}

	// modernc.org/sqlite driver name is "sqlite".
	// FK enforces relations, WAL allows readers during writes,
	// synchronous=NORMAL is the safe+fast pairing for WAL mode,
	// busy_timeout makes a concurrent writer wait instead of SQLITE_BUSY.
	dsn := "file:" + dbPath + "?cache=shared&_pragma=foreign_keys(1)&_pragma=journal_mode(WAL)&_pragma=synchronous(NORMAL)&_pragma=busy_timeout(5000)"

	sqldb, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, err
	}

	if err := sqldb.Ping(); err != nil {
		return nil, err
	}

	return bun.NewDB(sqldb, sqlitedialect.New()), nil
}
