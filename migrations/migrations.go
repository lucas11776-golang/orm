package migrations

import (
	"fmt"

	"github.com/lucas11776-golang/orm"
)

type Migration interface {
	Up()
	Down()
}

// Comment
func database(name string) orm.Database {
	db := orm.DB.Database(name)

	if db == nil {
		panic(fmt.Errorf("database %s does not exits", name))
	}

	return db
}

// Comment
func Create(connection string, table string, builder func(table *Table)) {
	db := database(connection)

	tb := &Table{}

	builder(tb)

	db.Migration().Migrate(&orm.TableScheme{
		Name:    table,
		Columns: tb.columns,
	})
}

// Comment
func Drop(connection string, table string) {
	database(connection).Migration().Drop(table)
}

type Migrator struct {
	migrations []Migration
}

// Comment
func Migrations(migration ...Migration) *Migrator {
	return &Migrator{migrations: migration}
}

// Comment
func (ctx *Migrator) Up() {
	for _, migration := range ctx.migrations {
		migration.Up()
	}
}

// Comment
func (ctx *Migrator) Down() {
	for _, migration := range ctx.migrations {
		migration.Down()
	}
}
