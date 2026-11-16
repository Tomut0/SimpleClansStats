package sqlite

import (
	"database/sql"
	"strings"

	"github.com/Tomut0/SimpleClansStats/internal/infrastructure/persistence/sqlite/generated"
	_ "modernc.org/sqlite"
)

type SQLite struct {
	Db      *sql.DB
	Queries *generated.Queries
}

func New(dbPath string) *SQLite {
	if dbPath == "" {
		panic("dbPath is empty")
	}

	path, options, _ := strings.Cut(dbPath, "?")
	dsn := path + "?_foreign_keys=on&_busy_timeout=4000"
	if options != "" {
		dsn += "&" + options
	}
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		panic(err)
	}

	queries := generated.New(db)

	return &SQLite{
		Db:      db,
		Queries: queries,
	}
}

func (s *SQLite) Close() error {
	return s.Db.Close()
}
