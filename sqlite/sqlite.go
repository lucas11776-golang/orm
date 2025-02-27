package sqlite

import (
	"database/sql"
	"orm/sqlite/migrations"

	_ "github.com/mattn/go-sqlite3"

	"orm"
)

type SQLite struct {
	DB orm.SQL
}

// Comment
func (ctx *SQLite) Query(statement *orm.Statement) (orm.Results, error) {
	return nil, nil
}

// Comment
func (ctx *SQLite) Count(statement *orm.Statement) (int64, error) {
	return 0, nil
}

// Comment
func (ctx *SQLite) Insert(statement *orm.Statement) (orm.Result, error) {
	return nil, nil
}

// Comment
func (ctx *SQLite) Database() interface{} {
	return ctx.DB
}

// Comment
func (ctx *SQLite) Migration() orm.Migration {
	return &migrations.Migration{
		DB: ctx.DB,
	}
}

// Comment
func Connect(source string) (orm.Database, error) {
	db, err := sql.Open("sqlite3", source)

	if err != nil {
		return nil, err
	}

	return &SQLite{
		DB: db,
	}, nil
}
