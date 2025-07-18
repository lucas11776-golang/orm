package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/lucas11776-golang/orm"
	driver "github.com/lucas11776-golang/orm/drivers/sql"
	"github.com/lucas11776-golang/orm/drivers/sqlite/migrations"
	utils "github.com/lucas11776-golang/orm/utils/sql"
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
	return utils.TableInfoPrimaryKey(ctx.db, table)
}

// Comment
func Connect(source string) orm.Database {
	db, err := sql.Open("sqlite3", source)

	if err != nil {
		panic(err)
	}

	return driver.NewSQLDriver(&driver.DriverOptions{
		QueryBuilder: &driver.DefaultQueryBuilder{},
		Migration:    &migrations.Migration{DB: db},
		Database:     &SQLiteDatabase{db: db},
	})
}
