package sqlite

import (
	"database/sql"
	"orm/sqlite/migrations"

	_ "github.com/mattn/go-sqlite3"

	"orm"
)

type Database *sql.DB

type SQLite struct {
	db Database
}

// Comment
func (ctx *SQLite) Query() orm.Query {
	return &Query{}
}

// Comment
func (ctx *SQLite) Database() interface{} {
	return nil
}

// Comment
func (ctx *SQLite) Migration() orm.Migration {
	return &migrations.Migration{DB: ctx.db}
}

// Comment
func Connect(source string) (orm.Database, error) {
	db, err := sql.Open("sqlite3", source)

	if err != nil {
		return nil, err
	}

	return &SQLite{db: db}, nil
}
