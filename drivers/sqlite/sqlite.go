package sqlite

import (
	"database/sql"

	_ "github.com/tursodatabase/go-libsql"

	"github.com/lucas11776-golang/orm"
	sqlDriver "github.com/lucas11776-golang/orm/drivers/sql"
	"github.com/lucas11776-golang/orm/drivers/sqlite/migrations"
)

type SQLite struct {
	*sqlDriver.SQL
}

type DB struct {
	db *sql.DB
}

// Comment
func (ctx *DB) DB() *sql.DB {
	return ctx.db
}

// Comment
func (ctx *DB) TablePrimaryKey(table string) (key string, err error) {
	return "id", nil
}

// Comment
func (ctx *SQLite) Migration() orm.Migration {
	return &migrations.Migration{DB: ctx.SQL.DB.DB()}
}

// Comment
func Connect(source string) orm.Database {
	db, err := sql.Open("libsql", source)

	if err != nil {
		panic(err)
	}

	return &SQLite{
		SQL: &sqlDriver.SQL{
			Builder: &sqlDriver.SQLBuilder{},
			DB: &DB{
				db: db,
			},
		},
	}
}
