package orm

import (
	"log"
	"reflect"
)

const DefaultDatabaseName = "default"

type db map[string]Database

type Models []interface{}

type Migration interface {
	Migrate(models Models) error
	Truncate(models Models) error
}

type QueryResults []map[string]interface{}

type Database interface {
	Query(statement *Statement) QueryResults
	Database() interface{}
	Migration() Migration
}

var DB = db{}

// Comment
func (ctx *db) Database(name string) Database {
	db, ok := DB[name]

	if !ok {
		return db
	}

	return db
}

// Comment
func Model[T any](model T) QueryBuilder[T] {
	if reflect.ValueOf(model).Type().Kind() != reflect.Struct {
		log.Fatalf("Model is not type of struct: %v", model)
	}

	return &QueryStatement[T]{
		Statement: &Statement{
			Model: model,
		},
	}
}

// Comment
func (ctx *db) Add(name string, database Database) {
	DB[name] = database
}

// Comment
func (ctx *db) Remove(name string) {
	delete(DB, name)
}
