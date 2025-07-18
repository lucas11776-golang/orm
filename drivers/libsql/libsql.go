package libsql

import (
	"database/sql"

	_ "github.com/tursodatabase/go-libsql"

	"github.com/lucas11776-golang/orm"
	driver "github.com/lucas11776-golang/orm/drivers/sql"
	"github.com/lucas11776-golang/orm/drivers/sqlite/migrations"
)

type LibSQL struct {
	db *sql.DB
}

// Comment
func (ctx *LibSQL) DB() *sql.DB {
	return ctx.db
}

// Comment
func (ctx *LibSQL) TablePrimaryKey(table string) (key string, err error) {
	return "id", nil
}

// Comment
func Connect(source string) orm.Database {
	db, err := sql.Open("libsql", source)

	if err != nil {
		panic(err)
	}

	return driver.NewSQLDriver(&driver.DriverOptions{
		QueryBuilder: &driver.DefaultQueryBuilder{},
		Migration:    &migrations.Migration{DB: db},
		Database:     &LibSQL{db: db},
	})
}
