package orm

import (
	"database/sql"
	"log"
	str "orm/utils/strings"
	"reflect"
	"strings"
)

const DefaultDatabaseName = "default"

type SQL *sql.DB
type MongoDB interface{}

type db map[string]Database

type Models []interface{}

type Migration interface {
	Migrate(models Models) error
	Truncate(models Models) error
}

type Result map[string]interface{}

type Results []Result

type Database interface {
	Query(statement *Statement) (Results, error)
	Count(statement *Statement) (int64, error)
	Insert(statement *Statement) (Result, error)
	Update(values Values) error
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
func TableName(model interface{}) string {
	vType := reflect.ValueOf(model)

	if vType.Kind() != reflect.Struct {
		log.Fatalf("Model is not a struct: %v", model)
	}

	return str.Plural(strings.ToLower(vType.Type().Name()))
}

type options struct {
	connection string
	table      string
}

// Comment
func getOptions(model interface{}) *options {
	opt := &options{
		connection: DefaultDatabaseName,
		table:      TableName(model),
	}

	vType := reflect.ValueOf(model).Type()

	for i := 0; i < vType.NumField(); i++ {
		tag := vType.Field(i).Tag

		if tag.Get("connection") != "" {
			opt.connection = tag.Get("connection")
		}

		if tag.Get("table") != "" {
			opt.table = tag.Get("table")
		}
	}

	return opt
}

// Comment
func Model[T any](model T) QueryBuilder[T] {
	if reflect.ValueOf(model).Type().Kind() != reflect.Struct {
		log.Fatalf("Model is not type of struct: %v", model)
	}

	options := getOptions(model)
	database, ok := DB[options.connection]

	if !ok {
		log.Fatalf("Connection %s does not exist", options.connection)
	}

	return &QueryStatement[T]{
		Model:     model,
		Database:  database,
		Statement: &Statement{},
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
