package sqlite

import (
	"database/sql"

	_ "github.com/tursodatabase/go-libsql"

	"github.com/lucas11776-golang/orm"
	sqlDriver "github.com/lucas11776-golang/orm/drivers/sql"
	"github.com/lucas11776-golang/orm/drivers/sqlite/migrations"
)

type SQLiteDatabase struct {
	db *sql.DB
}

// Comment
func (ctx *SQLiteDatabase) DB() *sql.DB {
	return ctx.db
}

// Comment
func (ctx *SQLiteDatabase) TablePrimaryKey(table string) (key string, err error) {
	return "id", nil
}

// Comment
func Connect(source string) orm.Database {
	db, err := sql.Open("libsql", source)

	if err != nil {
		panic(err)
	}

	return sqlDriver.NewSQLDriver(&sqlDriver.DriverOptions{
		QueryBuilder: &sqlDriver.DefaultQueryBuilder{},
		Migration:    &migrations.Migration{DB: db},
		Database:     &SQLiteDatabase{db: db},
	})
}
