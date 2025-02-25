package orm

const DefaultDatabaseName = "default"

type DB map[string]Database

type Models []interface{}

type Migration interface {
	Migrate(models Models) error
	Truncate(models Models) error
}

type Database interface {
	Query() Query
	Database() interface{}
	Migration() Migration
}

var Databases = DB{}

// Comment
func (ctx *DB) Database(name string) interface{} {
	db, ok := Databases[name]

	if !ok {
		return db
	}

	return db
}

// Comment
func (ctx *DB) DefaultDatabase() interface{} {
	return ctx.Database(DefaultDatabaseName)
}

// Comment
func (ctx *DB) Add(name string, database Database) {
	Databases[name] = database
}

// Comment
func (ctx *DB) Remove(name string) {
	delete(Databases, name)
}
