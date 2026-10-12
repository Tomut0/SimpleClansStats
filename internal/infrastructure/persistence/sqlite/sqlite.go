package sqlite

import (
	"context"
	"database/sql"

	"github.com/Tomut0/SimpleClansStats/internal/infrastructure/persistence/sqlite/generated"
	_ "modernc.org/sqlite"
)

type SQLite struct {
	Db      *sql.DB
	Ctx     context.Context
	Queries *generated.Queries
}

func New(dbPath string) *SQLite {
	if dbPath == "" {
		panic("dbPath is empty")
	}

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		panic(err)
	}

	queries := generated.New(db)

	return &SQLite{
		Db:      db,
		Ctx:     context.Background(),
		Queries: queries,
	}
}

func (s *SQLite) Close() error {
	return s.Db.Close()
}
